<p align="center">
    <a href="https://godoc.org/github.com/hihebark/godirsearch">
        <img src="https://godoc.org/github.com/hihebark/godirsearch?status.svg" alt="GoDoc">
    </a>
    <a href="https://goreportcard.com/badge/github.com/hihebark/godirsearch">
        <img src="https://goreportcard.com/badge/github.com/hihebark/godirsearch" alt="GoReportCard">
    </a>
    <a href="https://travis-ci.org/hihebark/godirsearch.svg?branch=master">
        <img src="https://travis-ci.org/hihebark/godirsearch.svg?branch=master" alt="travis">
    </a>
    https://travis-ci.org/hihebark/godirsearch.svg?branch=master
</p>

![Golang gopher](https://golang.org/doc/gopher/pkg.png)

godirsearch
===========
Nothing to see here!

TODO:
-----

- [x] DONE regex for verifing url
- [x] DONE check the connecivity to the url
- [x] DONE set extension
- [x] DONE set proxy
- [x] DONE add support of Tor (not completly)
- [x] DONE add read from grepproxylist.sh
- [ ] TODO set Cookies
- [ ] TODO generating json file of the result +
- [ ] TODO a log file for debbuging
- [ ] TODO use thread
- [ ] TODO why no an UI Just saying

Usage:
------

```
	GoDirSearch ~0.0.2
Usage of GoDirSearch:
  -cookie string
    	cookie
  -host string
    	Host to brutforce
  -proxy string
    	Use a proxy to brutforce
  -proxyfile string
    	Use a proxy file
  -thread int
    	Number of thread (default 4)
  -tor
    	Brutforce using Tor
  -useragent string
    	userAgent (default "Golang_Spider_Bot/3.0")
  -worlist string
    	wordlist to brutforce (default "test.txt")
```
