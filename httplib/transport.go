package httplib

import (
	"crypto/tls"
	"net/http"
)

// Create a Transport that supports insecure https

func NewTransport(insecure bool) http.RoundTripper {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	if insecure {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return transport
}
