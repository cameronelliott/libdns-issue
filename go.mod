module github.com/cameronelliott/libdns-issue

go 1.16

require (
	github.com/caddyserver/certmagic v0.12.0
	github.com/libdns/duckdns v0.1.0
)

replace github.com/libdns/duckdns => ../duckdns
