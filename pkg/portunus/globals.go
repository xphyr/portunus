package portunus

import (
	"bytes"
	"log"
	"os"
)

var (
	buf        bytes.Buffer
	logger     = log.New(&buf, "logger ", log.Lshortfile)
	myHostname = os.Getenv("HOSTNAME")
)
