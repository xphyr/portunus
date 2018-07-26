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
	"flag"
	"net"
	"net/http"

	logrus "github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	var status string

	target := r.URL.Query().Get("target")
	port := r.URL.Query().Get("port")
	if target == "" {
		http.Error(w, "'target' parameter must be specified", 400)
		return
	}
	if port == "" {
		http.Error(w, "'port' parameter must be specified", 400)
		return
	}
	destination := target + ":" + port
	log.Printf("Testing connection to %v", destination)
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		log.Println("Connection error:", err)
		status = "Unreachable"
	} else {
		status = "Online"
		conn.Close()
	}
	log.Println(status)
}
func init() {
	log.Formatter = new(logrus.TextFormatter)
	flag.Parse()
}
func main() {
	log.Info("Starting Portunus")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head>
            <title>Portunus Port Tester</title>
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
            </head>
            <body>
            <h1>Portunus Port Tester</h1>
            <form action="/query">
            <label>Target:</label> <input type="text" name="target" placeholder="X.X.X.X" value="1.2.3.4"><br>
			<label>Port:</label> <input type="text" name="port" placeholder="80" value="80"><br>
			<input type="submit" value="Submit">
            </form>
            </html>`))
	})
	http.HandleFunc("/query", queryHandler)
	listenPort := ":8080"
	log.Info("Listening on port: ", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
