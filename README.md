<p align="center">
    <a href="https://godoc.org/github.com/hihebark/godirsearch">
        <img src="https://godoc.org/github.com/hihebark/godirsearch?status.svg" alt="GoDoc">
    </a>
    <a href="https://goreportcard.com/report/github.com/hihebark/godirsearch">
        <img src="https://goreportcard.com/badge/github.com/hihebark/godirsearch" alt="GoReportCard">
    </a>
    <a href="https://travis-ci.org/hihebark/godirsearch">
        <img src="https://travis-ci.org/hihebark/godirsearch.svg?branch=master" alt="travis">
    </a>
    <a href="https://github.com/hihebark/godirsearch/blob/master/LICENSE">
        <img src="https://img.shields.io/aur/license/yaourt.svg" alt="license">
    </a>
</p>

<p align="center">
	<img src="https://golang.org/doc/gopher/pkg.png">
</p>

Godirsearch
===========

Godirsearch is a golang application to brute force web site and search for hidden or Folder and Directorie.

I started this project to learn Go-Lang so if you spot an error be kind a report'it i will digg'in to find a solution.

This project is still in a development.

Installation & Build:
------

Installation:

`go get github.com/hihebark/godirsearch`

Build:

`go build`

Usage:
------

```
 ▄▄▄▄
 █ ▄ █
 █▄▄▄█
 GoDirSearch ~0.4.4-Dev

Usage of ./GoDirSearch:
  -cookie string
    	Cookie if needed
  -ex string
    	Extension separate by comma like [php,txt]
  -host string
    	Host/Target to search for subdirectory example: http://example.com/
  -http
    	http server to consult the result
  -proxy string
    	Use a proxy to brutforce
  -proxyfile string
    	Use a proxy file (not set)
  -thread int
    	Number of thread (not set) (default 4)
  -tor
    	Use the test with Tor for anonymity
  -useragent string
    	UserAgent if not set it will generate one randomly
  -worlist string
    	Wordlist to use for the search (default "data/wordlist.txt")


```

> "The only way to learn a new programming language is by writing programs in it." - Dennis Ritchie

:octocat:
