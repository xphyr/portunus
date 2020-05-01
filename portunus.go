/*
Copyright 2018 Mark DeNeve.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xphyr/portunus/pkg/portunus"
)

var (
	buf        bytes.Buffer
	logger     = log.New(&buf, "logger ", log.Lshortfile)
	myHostname = os.Getenv("HOSTNAME")
	k8_node    = getk8s()
)

func init() {
	flag.Parse()
	logger.SetOutput(os.Stderr)
}

func getk8s() string {
	openshift_node := os.Getenv("MY_K8_NODE")
	if openshift_node != "" {
		return openshift_node
	}
	return ""
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`<html>`)) // Begin html block
	w.Write([]byte(`<head> 
		<title>Simple K8s Test Pod</title>
		<style>
		label{
		display:inline-block;
		width:75px;
		}
		form label {
		margin: 10px;
		}
		form input {
		margin: 10px;
		}
		</style>
		</head><body>`)) // Head  Begin Body
	w.Write([]byte(fmt.Sprintf("<h3>My hostname: %s</h3>", myHostname)))
	if k8_node != "" {
		w.Write([]byte(fmt.Sprintf("<hr><h3>My k8s node is: %v</h3><hr>", k8_node)))
	}
	portunus.GetPortTesterForm(w)
	w.Write([]byte(`<hr>`)) //Line Break
	portunus.GetDNSTesterForm(w)
	w.Write([]byte(`</body></html>`)) //END

}

func main() {

	logger.Printf("Starting App")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/porttester", portunus.PortTestHandler)
	http.HandleFunc("/dnstest", portunus.DNSTestHandler)
	http.HandleFunc("/whatsmypod", portunus.WhatsMyPod)
	listenPort := ":8080"
	logger.Printf("Listening on port: %v", listenPort)
	http.ListenAndServe(listenPort, nil)
}
