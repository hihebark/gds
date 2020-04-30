package core

// Options struct
type Options struct {
	URL        string
	Proxyfile  string
	Wordlist   string
	UserAgent  string
	Cookie     string
	Extensions []string
	Proxy      string
	Tor        bool
	Output     string
	Serve      bool
}

// NewOptions create new Options
func NewOptions(url, proxyFile, worldlist, useragent, cookie, proxy, output string,
	ext []string, tor, serve bool) *Options {
	return &Options{
		url,
		proxyFile,
		worldlist,
		useragent,
		cookie,
		ext,
		proxy,
		tor,
		output,
		serve,
	}
}
