package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/caddyserver/certmagic"
	"github.com/libdns/duckdns"
)

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Lookit my cool website over HTTPS!")
	})

	token := os.Getenv("DUCKDNS_TOKEN")
	if token == "" {
		panic("no token")
	}

	certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA
	certmagic.DefaultACME.DNS01Solver = &certmagic.DNS01Solver{
		DNSProvider: &duckdns.Provider{
			APIToken: token,
		},
	}

	err := certmagic.HTTPS([]string{"bug7.duckdns.org"}, mux)
	checkPanic(err)

}
