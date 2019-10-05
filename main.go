package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {

	pc, err := net.ListenPacket("udp", fmt.Sprint(":", conf.SrvPort))
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	fmt.Println("Server Listening on port: ", conf.SrvPort)

	for {
		buf := make([]byte, 1024)
		n, _, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		fmt.Println(string(buf[:n]))
		messages <- chanToRabbit{string(buf[0:n]), getRouteKey(buf[0:n])}
		//messages <- chanToRabbit{string(buf[0:n]), "tcx.metrics"}
	}
}

// stringContains checkes the srcString for any matches in the
// list, which is an array of strings.
func stringContains(a string, list []string) bool {
	for _, b := range list {
		if strings.Contains(a, b) {
			return true
		}
	}
	return false
}
