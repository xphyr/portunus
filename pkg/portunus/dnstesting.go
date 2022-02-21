package portunus

import (
	"fmt"
	"net"
	"net/http"
)

// DNSTestHandler exports a http response writer that tests DNS queries
func DNSTestHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("lookup")
	if target == "" {
		http.Error(w, "'hostname' parameter must be specified for dnslookup", 400)
		return
	}
	logger.Printf("Testing dns lookup of %v", target)
	ips, err := net.LookupIP(target)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("DNS lookup failed with error: %v\n", err)))
		logger.Printf("Testing dns lookup of %v failed with %v", target, err)
		return
	}
	for _, ip := range ips {
		w.Write([]byte(fmt.Sprintf("%v IN A %s\n", target, ip.String())))
	}
}

// GetDNSTesterForm creates the dns test html form
func GetDNSTesterForm(w http.ResponseWriter) {
	w.Write([]byte(`<h2>DNS Tester</h2>
			<form action="/dnstest">
            <label>Lookup:</label> <input type="text" name="lookup" placeholder="example.com" value="example.com"><br>
			<input type="submit" value="Submit">
            </form>`))
}
