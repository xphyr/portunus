package portunus

import (
	"fmt"
	"net/http"
)

// WhatsMyPod returns the hostname of the pod the app is running in
func WhatsMyPod(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("I am running on pod %s\n", myHostname)))
}
