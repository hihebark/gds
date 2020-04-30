package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/hihebark/godirsearch/core"
)

const (
	name    string = "GoDirsearch"
	version string = "0.7.0"
	banner  string = " ▄▄▄▄\n █ ▄ █ " + name + "\n █▄▄▄█ " + version + "\n"
	regexH  string = `^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`
	refile  string = `^(?:https?:\/\/+)`
)

var (
	tor       *bool
	http      *bool
	host      *string
	proxy     *string
	cookie    *string
	wordlist  *string
	output    *string
	proxyfile *string
	userAgent *string
	extension *string
)

func init() {
	extension = flag.String("ex", "", "Extension separate by comma like php,txt .")
	tor = flag.Bool("tor", false, "Use the test with Tor for anonymity.")
	host = flag.String("host", "", "Host/Target to search for subdirectory example: http://example.com/ .")
	proxy = flag.String("proxy", "", "Use a proxy to brutforce.")
	cookie = flag.String("cookie", "", "Cookie if needed.")
	wordlist = flag.String("wordlist", "", "Wordlist to use for the search.")
	output = flag.String("output", "", "Output result as a JSON")
	proxyfile = flag.String("proxyfile", "", "Use a proxy file (not set).")
	http = flag.Bool("http", false, "http server to consult the result.")
	userAgent = flag.String("useragent", "", "UserAgent if not set it will generate one randomly.")
}

func main() {
	fmt.Printf("%s\n", banner)
	flag.Parse()
	catchExit()
	if *host != "" {
		status, err := core.Healthz(*host)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		re := regexp.MustCompile(regexH)
		if re.MatchString(*host) && (status >= 200 && status < 500) {
			fmt.Printf("Chaeking connectivity: OK\n")
			if !strings.HasSuffix(*host, "/") {
				*host += "/"
			}
			if *userAgent == "" {
				*userAgent = "Mozilla/5.0 (X11; Linux i586; rv:31.0) Gecko/20100101 Firefox/31"
			}
			options := core.Options{
				URL:        *host,
				Proxyfile:  *proxyfile,
				Wordlist:   *wordlist,
				UserAgent:  *userAgent,
				Cookie:     *cookie,
				Extensions: strings.Split(*extension, ","),
				Proxy:      *proxy,
				Output:     *output,
				Serve:      *http,
			}
			core.NewRequest(options).Run()

		}
	}
	fmt.Println("Done...")
}

func catchExit() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("Ctrl+c detected quiting now ...")
		os.Exit(1)
	}()
}
