package portunus

import (
	"bytes"
	"log"
)

var (
	buf bytes.Buffer
	logger = log.New(&buf, "logger ", log.Lshortfile)
)