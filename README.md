<p align="center">
    <a href="https://godoc.org/github.com/hihebark/gds">
        <img src="https://godoc.org/github.com/hihebark/gds?status.svg" alt="GoDoc">
    </a>
    <a href="https://goreportcard.com/report/github.com/hihebark/gds">
        <img src="https://goreportcard.com/badge/github.com/hihebark/gds" alt="GoReportCard">
    </a>
    <a href="https://travis-ci.org/hihebark/gds">
        <img src="https://travis-ci.org/hihebark/gds.svg?branch=master" alt="travis">
    </a>
    <a href="https://github.com/hihebark/gds/blob/master/LICENSE">
        <img src="https://img.shields.io/aur/license/yaourt.svg" alt="license">
    </a>
    <a href="https://codecov.io/gh/hihebark/ds">
        <img src="https://codecov.io/gh/hihebark/gds/branch/master/graph/badge.svg" />
    </a>
</p>

<p align="center">
	<a href="https://hihebark.github.io/gds/">
		<img src="logo.png" width="300">
	</a>
</p>

Go directory search
===========

Gds is a golang application to brute force web site and search for hidden files or directories.

Installation & Build:
---------------------

Installation:

`go get github.com/hihebark/gds`

Build:

`go build`

Or get a pre-released version:

[Gds releases](https://github.com/hihebark/gds/releases)

Usage:
------

`./gds -host http://example.com/ -e txt,php -wordlist ~/path/to/my/dictionary.txt`
```sh
 ▄▄▄▄
 █ ▄ █ GDS
 █▄▄▄█ 1.0

  -cookie string
        Cookie if needed.
  -ex string
        Extension separate by comma like php,txt .
  -host string
        Host/Target to search for subdirectory example: http://example.com/ .
  -http
        http server to consult the result.
  -output string
        Output result as a JSON
  -proxy string
        Use a proxy to brutforce.
  -proxyfile string
        Use a proxy file (not set).
  -tor
        Use the test with Tor for anonymity.
  -useragent string
        UserAgent if not set it will generate one randomly.
  -wordlist string
        Wordlist to use for the search.
```