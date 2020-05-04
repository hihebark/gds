package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/hihebark/gds/core"
)

const (
	name    string = "g.d.s"
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
	tor = flag.Bool("tor", false, "Use the test with Tor for anonymity.")
	host = flag.String("host", "", "Host/Target to search for subdirectory example: http://example.com/ .")
	http = flag.Bool("http", false, "http server to consult the result.")
	proxy = flag.String("proxy", "", "Use a proxy to brutforce.")
	cookie = flag.String("cookie", "", "Cookie if needed.")
	output = flag.String("output", "", "Output result as a JSON")
	wordlist = flag.String("wordlist", "", "Wordlist to use for the search.")
	proxyfile = flag.String("proxyfile", "", "Use a proxy file (not set).")
	extension = flag.String("ex", "", "Extension separate by comma like php,txt .")
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
			fmt.Printf("Cheking connectivity: OK\n")
			if !strings.HasSuffix(*host, "/") {
				*host += "/"
			}
			if *userAgent == "" {
				*userAgent = "Mozilla/5.0 (X11; Linux i586; rv:31.0) Gecko/20100101 Firefox/31"
			}
			options := core.Options{
				URL:        *host,
				Proxy:      *proxy,
				Serve:      *http,
				Output:     *output,
				Cookie:     *cookie,
				Wordlist:   *wordlist,
				UserAgent:  *userAgent,
				Proxyfile:  *proxyfile,
				Extensions: strings.Split(*extension, ","),
			}
			core.NewRequest(options).Run()
			fmt.Println("\nDone...")
		} else {
			fmt.Printf("Cheking connectivity: NOK\n")
			fmt.Printf("Please check that you're connected or set a proxy!.\n")
		}
	}
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
