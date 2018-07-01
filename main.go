package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/hihebark/godirsearch/lib"
)

//Const
const (
	version string = "0.6.0"
	LOGO    string = " ▄▄▄▄\n █ ▄ █\n █▄▄▄█\n"
)

var (
	tor, http                                               *bool
	thread                                                  *int
	host, proxy, cookie, wordlist, proxyfile, userAgent, ex *string
)

func init() {

	ex = flag.String("ex", "", "Extension separate by comma like php,txt .")
	tor = flag.Bool("tor", false, "Use the test with Tor for anonymity.")
	host = flag.String("host", "", "Host/Target to search for subdirectory example: http://example.com/ .")
	proxy = flag.String("proxy", "", "Use a proxy to brutforce.")
	thread = flag.Int("thread", 4, "Number of thread (not set).")
	cookie = flag.String("cookie", "", "Cookie if needed.")
	wordlist = flag.String("wordlist", "data/wordlist.txt", "Wordlist to use for the search.")
	proxyfile = flag.String("proxyfile", "", "Use a proxy file (not set).")
	http = flag.Bool("http", false, "http server to consult the result.")
	userAgent = flag.String("useragent", "", "UserAgent if not set it will generate one randomly.")

}

func main() {

	fmt.Printf("%s GoDirSearch \033[92m~%s\n\n\033[0m", lib.SayMe(lib.LIGHTRED, LOGO), version)
	flag.Parse()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf(lib.SayMe(lib.LIGHTRED, "Ctrl+c detected quiting now ..."))
		os.Exit(1)
	}()

	if *host != "" {

		re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
		status := lib.CheckConnectivity(*host)
		if re.MatchString(*host) && (status >= 200 && status < 300) {
			if !strings.HasSuffix(*host, "/") {
				*host += "/"
			}
			if *userAgent == "" {
				*userAgent = strings.Split(lib.GetRandLine("data/user-agents.txt"), "\n")[0]
			}
			if *http {
				go func() {
					lib.StartListning()
				}()
			}
			lib.Run(fmt.Sprintf("Connection to %s %s\n",
				lib.SayMe(lib.LIGHTRED, *host),
				lib.SayMe(lib.GREEN, "OK")))
			refolder := regexp.MustCompile(`^(?:https?:\/\/+)`)
			resultFile := refolder.Split(*host, 2)[1]
			os.MkdirAll("data/results/"+resultFile, 0755)
			o := lib.Options{
				Host:       *host,
				Proxyfile:  *proxyfile,
				Wordlist:   *wordlist,
				UserAgent:  *userAgent,
				Cookie:     *cookie,
				Ex:         strings.Split(*ex, ","),
				Proxy:      *proxy,
				Tor:        *tor,
				ResultFile: resultFile,
				IsUp:       *http,
			}
			lib.StartWork(o)
			//lib.StartSearch(req)
		} else {
			lib.Good(fmt.Sprintf("Connection to %s %s\n",
				lib.SayMe(lib.LIGHTRED, *host),
				lib.SayMe(lib.RED, "Not reachable")))
		}
	} else {
		lib.StartListning()
	}
}
