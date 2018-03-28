<p align="center">

![GoDoc](https://godoc.org/github.com/hihebark/godirsearch?status.svg)](https://godoc.org/github.com/hihebark/godirsearch)

![Golang gopher](https://golang.org/doc/gopher/pkg.png)
</p>
godirsearch
===========
Nothing to see here!

TODO:
-----

- [x] DONE regex for verifing url
- [x] DONE check the connecivity to the url
- [x] DONE set extension
- [x] DONE set proxy
- [x] TODO add support of Tor (not completly)
- [ ] TODO add read from grepproxylist.sh
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
