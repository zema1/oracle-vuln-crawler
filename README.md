# 一个 Oracle 历史漏洞爬取工具

通过制定关键字，可以自动检索 WebLogic, Database, Management Center, Testing Suite 等历史漏洞并统计。

## 准备
```
go build .
```

## 运行

+ 检索所有历史 WebLogic 漏洞，输出到屏幕

```
./main --filter WebLogic
```

+ 检索所有历史 WebLogic 漏洞, 输出到 weblogic.md

```
./main --filter WebLogic --output weblogic.md
```


+ 检索所有历史 WebLogic 漏洞，输出到 weblogic.md，并检查对应 CVE github 是否有 repo, token 是 Github Token, 越多速度越快，同一账户的多个 token 会被视为一个。

```
./main --filter WebLogic --output weblogic.md --tokens token1,token2,token3
```
+ 抽取 Github 搜索结果不为 No 的

```
awk -F '|' '$9!=" No "  {print $0}' weblogic.md > weblogic2.md
```

## Weblogic Demo

|  CVE-ID | Product | Component | Protocol | NeedAuth | AffectedVersion | Alert/Patch | GithubInfo |
|  ----  | ----  | ----  | ----  | ----  | ----  | ---- | ---- |
| CVE-2020-2546 | Oracle WebLogic Server | Application Container - JavaEE | T3 | true | 10.3.6.0.0,12.1.3.0.0 | [CPU - January 2020 ](https://www.oracle.com/security-alerts/cpujan2020.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2020-2546) |
| CVE-2020-2551 | Oracle WebLogic Server | WLS Core Components | IIOP | true | 10.3.6.0.0,12.1.3.0.0,12.2.1.3.0,12.2.1.4.0 | [CPU - January 2020 ](https://www.oracle.com/security-alerts/cpujan2020.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2020-2551) |
| CVE-2015-9251 | Oracle WebLogic Server | Sample apps (jQuery) | HTTP | true | 12.1.3.0,12.2.1.3 | [CPU - January 2019 ](https://www.oracle.com/security-alerts/cpujan2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2015-9251) |
| CVE-2019-2888 | Oracle WebLogic Server | EJB Container | HTTP | true | 10.3.6.0.0,12.1.3.0.0,12.2.1.3.0,12.2.1.4.0 | [CPU - October 2019 ](https://www.oracle.com/security-alerts/cpuoct2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-2888) |
| CVE-2015-9251 | Oracle WebLogic Server | Web Services (jQuery) | HTTP | true | 12.1.3.0.0,12.2.1.3.0 | [CPU - October 2019 ](https://www.oracle.com/security-alerts/cpuoct2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2015-9251) |
| CVE-2019-11358 | Oracle WebLogic Server | Sample apps (jQuery) | HTTP | true | 12.1.3.0.0,12.2.1.3.0 | [CPU - October 2019 ](https://www.oracle.com/security-alerts/cpuoct2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-11358) |
| CVE-2019-11358 | Oracle WebLogic Server | Console (jQuery) | HTTP | true | 10.3.6.0.0,12.1.3.0.0,12.2.1.3.0 | [CPU - October 2019 ](https://www.oracle.com/security-alerts/cpuoct2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-11358) |
| CVE-2019-2890 | Oracle WebLogic Server | Web Services | T3 | false | 10.3.6.0.0,12.1.3.0.0,12.2.1.3.0 | [CPU - October 2019 ](https://www.oracle.com/security-alerts/cpuoct2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-2890) |
| CVE-2019-2729 | Oracle WebLogic Server | Web Services | HTTP | true | 10.3.6.0.0,12.1.3.0.0,12.2.1.3.0 | [Alert for CVE-2019-2729](https://www.oracle.com/security-alerts/alert-cve-2019-2729.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-2729) |
| CVE-2019-2725 | Oracle WebLogic Server | Web Services | HTTP | true | 10.3.6.0,12.1.3.0 | [Alert for CVE-2019-2725](https://www.oracle.com/security-alerts/alert-cve-2019-2725.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-2725) |
| CVE-2019-2615 | Oracle WebLogic Server | WLS Core Components | HTTP | false | 10.3.6.0.0,12.1.3.0.0,12.2.1.3.0 | [CPU - April 2019 ](https://www.oracle.com/security-alerts/cpuapr2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-2615) |
| CVE-2019-2618 | Oracle WebLogic Server | WLS Core Components | HTTP | false | 10.3.6.0.0,12.1.3.0.0,12.2.1.3.0 | [CPU - April 2019 ](https://www.oracle.com/security-alerts/cpuapr2019.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2019-2618) |
| CVE-2015-7501 | Oracle WebLogic Server | None | HTTP | true | 10.3.6.0,12.1.3.0,12.2.1.0 | [CPU - October 2016 ](https://www.oracle.com/security-alerts/cpuoct2016.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2015-7501) |
| CVE-2018-3248 | Oracle WebLogic Server | WLS - Web Services | HTTP | true | 10.3.6.0 | [CPU - October 2018 ](https://www.oracle.com/security-alerts/cpuoct2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-3248) |
| CVE-2018-3252 | Oracle WebLogic Server | WLS Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.3 | [CPU - October 2018 ](https://www.oracle.com/security-alerts/cpuoct2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-3252) |
| CVE-2018-3245 | Oracle WebLogic Server | WLS Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.3 | [CPU - October 2018 ](https://www.oracle.com/security-alerts/cpuoct2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-3245) |
| CVE-2018-3191 | Oracle WebLogic Server | WLS Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.3 | [CPU - October 2018 ](https://www.oracle.com/security-alerts/cpuoct2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-3191) |
| CVE-2017-7525 | Oracle WebLogic Server | Sample apps (jackson-databind) | HTTP | true | 10.3.6.0,12.1.3.0,12.2.1.2,12.2.1.3 | [CPU - April 2018 ](https://www.oracle.com/security-alerts/cpuapr2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-7525) |
| CVE-2015-7501 | Oracle WebLogic Portal | - (Apache Commons Collections) | HTTP | false | 10.3.6.0.0 | [CPU - April 2018 ](https://www.oracle.com/security-alerts/cpuapr2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2015-7501) |
| CVE-2018-2628 | Oracle WebLogic Server | WLS Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.2,12.2.1.3 | [CPU - April 2018 ](https://www.oracle.com/security-alerts/cpuapr2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-2628) |
| CVE-2017-5645 | Oracle WebLogic Server | WL Diagnostics Framework (Apache Log4j) | HTTP | true | 10.3.6.0,12.1.3.0,12.2.1.2,12.2.1.3 | [CPU - April 2018 ](https://www.oracle.com/security-alerts/cpuapr2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-5645) |
| CVE-2018-2893 | Oracle WebLogic Server | WLS Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.2,12.2.1.3 | [CPU - July 2018 ](https://www.oracle.com/security-alerts/cpujul2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-2893) |
| CVE-2018-2894 | Oracle WebLogic Server | WLS - Web Services | HTTP | true | 12.1.3.0,12.2.1.2,12.2.1.3 | [CPU - July 2018 ](https://www.oracle.com/security-alerts/cpujul2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-2894) |
| CVE-2018-7489 | Oracle WebLogic Server | Console (jackson-databind) | HTTP | true | 12.2.1.2,12.2.1.3 | [CPU - July 2018 ](https://www.oracle.com/security-alerts/cpujul2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2018-7489) |
| CVE-2017-5645 | Oracle WebLogic Server | Sample apps (Apache Log4j) | TCP/UDP | true | 10.3.6.0.0,12.1.3.0.0,12.2.1.2.0,12.2.1.3.0 | [CPU - January 2018 ](https://www.oracle.com/security-alerts/cpujan2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-5645) |
| CVE-2017-10352 | Oracle WebLogic Server | WLS - Web Services | HTTP | true | 12.2.1.3.0 | [CPU - January 2018 ](https://www.oracle.com/security-alerts/cpujan2018.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-10352) |
| CVE-2017-10148 | Oracle WebLogic Server | Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.1,12.2.1.2 | [CPU - July 2017 ](https://www.oracle.com/security-alerts/cpujul2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-10148) |
| CVE-2017-10147 | Oracle WebLogic Server | Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.1,12.2.1.2 | [CPU - July 2017 ](https://www.oracle.com/security-alerts/cpujul2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-10147) |
| CVE-2017-5638 | Oracle WebLogic Server | Sample apps (Struts 2) | HTTP | true | 10.3.6.0,12.1.3.0,12.2.1.1,12.2.1.2 | [CPU - July 2017 ](https://www.oracle.com/security-alerts/cpujul2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-5638) |
| CVE-2017-10271 | Oracle WebLogic Server | WLS Security | T3 | true | 10.3.6.0.0,12.1.3.0.0,12.2.1.1.0,12.2.1.2.0 | [CPU - October 2017 ](https://www.oracle.com/security-alerts/cpuoct2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-10271) |
| CVE-2017-10352 | Oracle WebLogic Server | WLS-WebServices | HTTP | true | 10.3.6.0.0,12.1.3.0.0,12.2.1.1.0,12.2.1.2.0,12.2.1.3.0 | [CPU - October 2017 ](https://www.oracle.com/security-alerts/cpuoct2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-10352) |
| CVE-2017-3506 | Oracle WebLogic Server | Web Services | HTTP | true | 10.3.6.0,12.1.3.0,12.2.1.0,12.2.1.1,12.2.1.2 | [CPU - April 2017 ](https://www.oracle.com/security-alerts/cpuapr2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-3506) |
| CVE-2017-5638 | Oracle WebLogic Server | Samples (Struts 2) | HTTP | true | 10.3.6.0,12.1.3.0,12.2.1.0,12.2.1.1,12.2.1.2 | [CPU - April 2017 ](https://www.oracle.com/security-alerts/cpuapr2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-5638) |
| CVE-2017-3248 | Oracle WebLogic Server | Core Components | T3 | true | 10.3.6.0,12.1.3.0,12.2.1.0,12.2.1.1 | [CPU - January 2017 ](https://www.oracle.com/security-alerts/cpujan2017.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2017-3248) |
| CVE-2016-0638 | Oracle WebLogic Server | Java Messaging Service | JMS | true | 10.3.6,12.1.2,12.1.3,12.2.1 | [CPU - April 2016 ](https://www.oracle.com/security-alerts/cpuapr2016v3.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2016-0638) |
| CVE-2016-3510 | Oracle WebLogic Server | WLS Core Components | HTTP | true | 10.3.6.0,12.1.3.0,12.2.1.0 | [CPU - July 2016 ](https://www.oracle.com/security-alerts/cpujul2016.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2016-3510) |
| CVE-2013-2186 | Oracle WebLogic Portal | HTTP | Core Services | true | 10.3.6 | [CPU - January 2016 ](https://www.oracle.com/security-alerts/cpujan2016.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2013-2186) |
| CVE-2015-4852 | Oracle WebLogic Server | T3 | WLS Security | true | 10.3.6.0,12.1.2.0,12.1.3.0,12.2.1.0 | [Alert for CVE-2015-4852](https://www.oracle.com/security-alerts/alert-cve-2015-4852.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2015-4852) |
| CVE-2013-2186 | Oracle WebLogic Server | HTTP | WLS Config, WLS Console | true | 10.3.6.0,12.1.1.0,12.1.2.0,12.1.3.0 | [CPU - January 2015 ](https://www.oracle.com/security-alerts/cpujan2015.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2013-2186) |
| CVE-2014-0114 | Oracle WebLogic Portal | HTTP | Third Party Tools | true | 10.0.1.0,10.2.1.0,10.3.6.0 | [CPU - January 2015 ](https://www.oracle.com/security-alerts/cpujan2015.html) | [Yes](https://api.github.com/search/repositories?q=CVE-2014-0114) |
