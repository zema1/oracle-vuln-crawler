package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type AlertPage struct {
	Title   string
	Link    string
	Date    string
	Details []*AlertDetail
}

type AlertDetail struct {
	Source          *AlertPage
	CVEID           string
	Product         string
	Component       string
	Protocol        string
	NeedAuth        bool
	AffectedVersion string
	GithubSearch    string
}

func main() {
	var output string
	var parallel int
	var outFile io.WriteCloser
	var filterString string
	var token string
	var tokens []string
	flag.StringVar(&output, "output", "-", "markdown output file, default is stdout")
	flag.IntVar(&parallel, "parallel", 5, "crawl parallel")
	flag.StringVar(&filterString, "filter", "WebLogic", "filter results with keywords")
	flag.StringVar(&token, "tokens", "", "github tokens, comma separated")
	flag.Parse()
	if output == "-" {
		outFile = os.Stdout
	} else {
		var err error
		outFile, err = os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
		if err != nil {
			panic(err)
		}
	}
	defer outFile.Close()
	tokens = strings.Split(token, ",")
	if len(tokens) == 0 {
		log.Println("github token is empty, github search after crawl will be disabled")
	}
	log.Printf("filter string is %s\n", filterString)
	log.Printf("github tokens is %v", tokens)

	client := buildClient()
	alertPages, err := parseHomePage(client)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("got %d alerts and security updates", len(alertPages))

	sort.Slice(alertPages, func(i, j int) bool {
		return alertPages[i].Date > alertPages[j].Date
	})

	wg := sync.WaitGroup{}
	wg.Add(parallel)
	log.Printf("crawl parallel: %d\n", parallel)

	in := make(chan *AlertPage, 10)
	go func() {
		defer close(in)
		for _, page := range alertPages {
			in <- page
		}
	}()

	for i := 0; i < parallel; i++ {
		go func() {
			defer wg.Done()
			for page := range in {
				fetchDetails(client, page, tokens)
			}
		}()
	}
	wg.Wait()

	for _, page := range alertPages {
		var s []*AlertDetail
		for i, detail := range page.Details {
			if strings.Contains(detail.Product, filterString) {
				s = append(s, page.Details[i])
			}
		}
		page.Details = s
	}

	// github 有速率限制，这里没必要并发，根本快不起来
	if len(tokens) != 0 {
		for _, page := range alertPages {
			checkGithubRepo(page, client, tokens)
		}
	}

	log.Printf("writing to file...")
	n, err := writeToFile(alertPages, outFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote down, %d entries in total", n)
}

func buildClient() *http.Client {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &http.Client{
		Transport: tr,
	}
}

func sendReq(client *http.Client, url string) ([]byte, error) {
	firstReq, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := client.Do(firstReq)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func parseHomePage(client *http.Client) ([]*AlertPage, error) {
	home := "https://www.oracle.com/security-alerts/"
	log.Println("Fetching ", home)
	data, err := sendReq(client, home)
	if err != nil {
		return nil, err
	}

	var alertPages []*AlertPage
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	selection := doc.Find("tbody tr")
	cnt := selection.Size()
	for i := 0; i < cnt; i++ {
		part := selection.Eq(i)
		tds := part.Find("td")
		if tds.Size() != 2 {
			continue
		}
		info := tds.Eq(0)
		date := tds.Eq(1)

		a := info.Find("a")
		title := a.Text()
		href, ok := a.Attr("href")
		if !ok {
			continue
		}

		dateText := date.Text()
		parts := strings.Split(dateText, ",")
		if len(parts) != 2 {
			continue
		}
		t, err := time.Parse("02 January 2006", strings.TrimSpace(parts[1]))
		if err != nil {
			log.Println(err)
			continue
		}
		if strings.Contains(title, "Alert") || strings.Contains(title, "Critical Patch") {
			alertPages = append(alertPages, &AlertPage{
				Title: title,
				Link:  href,
				Date:  t.Format("2006-01-02"),
			})
		}
	}
	return alertPages, nil
}

func fetchDetails(client *http.Client, alertPage *AlertPage, tokens []string) {
	u := "https://www.oracle.com" + alertPage.Link
	log.Printf("fetching %s", u)
	data, err := sendReq(client, u)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	selection := doc.Find("tr")
	for i := selection.Size(); i >= 0; i-- {
		tr := selection.Eq(i)
		CVEID := tr.Find("th").Text()
		var surplus int
		if CVEID == "" {
			CVEID = tr.Find("span").Text()
			if CVEID != "" {
				surplus = 1
			}
		}
		if !strings.HasPrefix(CVEID, "CVE") || strings.HasPrefix(CVEID, "CVE#") {
			continue
		}
		tds := tr.Find("td")
		detail := &AlertDetail{
			Source:          alertPage,
			CVEID:           CVEID,
			Product:         tds.Eq(0 + surplus).Text(),
			Component:       tds.Eq(1 + surplus).Text(),
			Protocol:        tds.Eq(2 + surplus).Text(),
			NeedAuth:        strings.TrimSpace(tds.Eq(3+surplus).Text()) == "Yes",
			AffectedVersion: strings.Join(strings.Fields(tds.Eq(tds.Size()-2).Text()), ""),
		}
		alertPage.Details = append(alertPage.Details, detail)
	}
	if len(alertPage.Details) == 0 {
		panic("got bug for " + u)
	}
	log.Printf("%s ok", u)
}

func checkGithubRepo(alertPage *AlertPage, client *http.Client, tokens []string) {
	for _, detail := range alertPage.Details {
		var resp *http.Response
		var err error
		var urlStr string
		urlStr = fmt.Sprintf("https://api.github.com/search/repositories?q=%s", strings.TrimSpace(detail.CVEID))
		log.Println("checking github repo for ", urlStr, detail.CVEID)
		for {
			req, _ := http.NewRequest(http.MethodGet, urlStr, nil)
			req.Header.Set("Authorization", "token "+randomToken(tokens))
			resp, err = client.Do(req)
			if err != nil {
				log.Println(err)
				randomSleep(time.Second*3, time.Second*5)
				continue
			}
			if resp.StatusCode != 200 {
				log.Println("api limit,sleep a random time")
				randomSleep(time.Second*5, time.Second*10)
			} else {
				break
			}
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		_ = resp.Body.Close()
		m := make(map[string]interface{})
		err = json.Unmarshal(data, &m)
		if err != nil {
			log.Println(err)
			continue
		}
		cnt := int(m["total_count"].(float64))
		if cnt == 0 {
			log.Println("no repo found for ", detail.CVEID)
		} else {
			log.Println("found repos for", detail.CVEID)
			detail.GithubSearch = urlStr
		}
	}
}

func randomToken(tokens []string) string {
	return tokens[rand.Intn(len(tokens))]
}

func randomSleep(from, to time.Duration) {
	sleep := rand.Int63n(int64(to-from)) + int64(from)
	time.Sleep(time.Duration(sleep))
}

func writeToFile(alertPages []*AlertPage, outFile io.Writer) (int, error) {
	var total int
	_, err := outFile.Write([]byte("|  CVE-ID | Product | Component | Protocol | NeedAuth | AffectedVersion | Alert/Patch | GithubInfo |\n"))
	if err != nil {
		return 0, err
	}
	_, err = outFile.Write([]byte("|  ----  | ----  | ----  | ----  | ----  | ----  | ---- | ---- |\n"))
	if err != nil {
		return 0, err
	}
	for _, page := range alertPages {
		for _, detail := range page.Details {
			total++
			title := strings.ReplaceAll(detail.Source.Title, "Critical Patch Update", "CPU")
			link := "https://www.oracle.com" + detail.Source.Link
			github := detail.GithubSearch
			if github == "" {
				github = "No"
			} else {
				github = fmt.Sprintf("[Yes](%s)", github)
			}
			_, err = outFile.Write([]byte(fmt.Sprintf("| %s | %s | %s | %s | %v | %s | [%s](%s) | %s |\n", detail.CVEID, detail.Product,
				detail.Component, detail.Protocol, detail.NeedAuth, detail.AffectedVersion, title, link, github)))
			if err != nil {
				return 0, err
			}
		}
	}
	return total, nil
}
