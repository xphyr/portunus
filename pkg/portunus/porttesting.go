package portunus

import (
	"net"
	"net/http"

)
func PortTestHandler(w http.ResponseWriter, r *http.Request) {
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
	logger.Printf("Testing connection to %v", destination)
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		w.Write([]byte(`<html>
		<body><h1>
		Unreachable</h1>
		</body>
		</html>`))
	} else {
		conn.Close()
		w.Write([]byte(`<html>
		<body><h1>
		Success</h1>
		</body>
		</html>`))
	}
}
 
func GetPortTesterForm(w http.ResponseWriter) {
	w.Write([]byte(`<h2>Port Tester</h2>
            <form action="/porttester">
            <label>Target:</label> <input type="text" name="target" placeholder="X.X.X.X" value="1.2.3.4"><br>
			<label>Port:</label> <input type="text" name="port" placeholder="80" value="80"><br>
			<input type="submit" value="Submit">
			</form>`)) //Port tester.
}