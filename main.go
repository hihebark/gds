package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/hihebark/godirsearch/core"
)

const version string = "0.1.1"

var (
	tor                                                     *bool
	thread                                                  *int
	host, proxy, cookie, wordlist, proxyfile, userAgent, ex *string
)

func init() {

	ex = flag.String("ex", "txt", "separate with coma like php,txt ...")
	tor = flag.Bool("tor", false, "Brutforce using Tor")
	host = flag.String("host", "", "Host to brutforce")
	proxy = flag.String("proxy", "", "Use a proxy to brutforce")
	thread = flag.Int("thread", 4, "Number of thread")
	cookie = flag.String("cookie", "", "cookie")
	wordlist = flag.String("worlist", "test.txt", "wordlist to brutforce")
	proxyfile = flag.String("proxyfile", "", "Use a proxy file")
	userAgent = flag.String("useragent", "Golang_Spider_Bot/3.0", "userAgent")

}

func main() {

	fmt.Printf("\tGoDirSearch \033[92m~%s\n\033[0m", version)
	flag.Parse()
	if *host == "" {
		core.Que("No host argument found! add -host http://examples.com/")
		os.Exit(0)
	}

	/***************************************************************************
	 * Best regex `^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`
	 * http://www.golangprograms.com/golang-package-examples/regular-expression-to-extract-domain-from-url.html
	 ****************************************************************************/

	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	status := core.CheckConnectivty(*host)
	if re.MatchString(*host) && (status >= 200 && status < 300) {

		if !strings.HasSuffix(*host, "/") {
			*host += "/"
		}
		core.Run("Connection to the target Ok!")
		req := core.NetRequest{
			Host:      *host,
			Proxyfile: *proxyfile,
			Wordlist:  *wordlist,
			UserAgent: *userAgent,
			Cookie:    *cookie,
			Ex:        strings.Split(*ex, ","),
			Proxy:     *proxy,
			Tor:       *tor,
		}
		core.Fuxe(req)
		//core.GetBody(req)

	} else {
		core.Bad(fmt.Sprintf("Host not recheable status: %s", status))
	}

}
