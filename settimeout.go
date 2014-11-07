package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

var version = "0.2.2"
var favicon []byte
var index []byte
var robots []byte

//plain-text
var invalidFormat = []byte("invalid_format")
var done = []byte("done")

//int
var one = []byte("1")

//json
var jsonTrue = []byte("true")

//js
var jsTrue = []byte("_settimeoutio=true")

//css
var cssShow = []byte(".settimeoutio {display: block;}")

//callback
var emptyCallback = []string{""}

func main() {
	port := flag.String("addr", ":80", "HTTP address to listen on (empty to disable)")
	tcpPort := flag.String("tcpaddr", ":5103", "Socket address to listen on (empty to disable)")
	flag.Parse()

	var err error
	favicon, err = Asset("assets/favicon.ico")
	if err != nil {
		log.Fatal("Failed to read favicon.ico: " + err.Error())
	}
	index, err = Asset("assets/index.html")
	if err != nil {
		log.Fatal("Failed to read index.html: " + err.Error())
	}
	robots, err = Asset("assets/robots.txt")
	if err != nil {
		log.Fatal("Failed to read robots.txt: " + err.Error())
	}

	go startHTTPServer(*port)
	go startTCPServer(*tcpPort)

	//wait for a SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func startTCPServer(addr string) {
	if addr == "" {
		return
	}
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to start TCP server: " + err.Error())
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			// handle error
			continue
		}
		go socketHandler(conn)
	}
}

func startHTTPServer(addr string) {
	if addr == "" {
		return
	}
	err := http.ListenAndServe(addr, http.HandlerFunc(httpHandler))
	if err != nil {
		log.Fatal("Failed to start HTTP server: " + err.Error())
	}
}

func socketHandler(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	buf := bufio.NewReader(conn)
	str, err := buf.ReadString('\n')
	if err != nil {
		return
	}
	d, err := parseDurationString(str)
	if err != nil {
		conn.Write(invalidFormat)
		return
	}
	<-time.After(d)
	conn.Write(done)
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	reqHeader := w.Header()
	reqHeader.Set("Access-Control-Allow-Origin", "*")
	reqHeader.Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	reqHeader.Set("Server", "SetTimeout/"+version)
	if req.Method != "GET" && req.Method != "POST" {
		//assume its the browser just checking headers
		w.WriteHeader(http.StatusOK)
		return
	}
	str := strings.TrimLeft(req.URL.Path, "/")

	switch str {
	case "": //index page
		w.Write(index)
	case "favicon.ico":
		w.Write(favicon)
	case "robots.txt":
		w.Write(robots)
	default:
		d, err := parseDurationString(str)
		query := req.URL.Query()
		exists := func(key string) bool { _, ok := query[key]; return ok }
		var resp []byte
		var contentType string
		switch {
		case exists("js"):
			resp = jsTrue
			contentType = "application/javascript"
		case exists("callback"):
			callback, _ := query["callback"]
			if len(callback) == 0 {
				callback = emptyCallback
			}
			resp = []byte(callback[0] + "(true)")
			contentType = "application/javascript"
		case exists("int"):
			resp = one
			contentType = "text/plain"
		case exists("json"):
			resp = jsonTrue
			contentType = "application/json"
		case exists("css"):
			resp = cssShow
			contentType = "text/css"
		default:
			resp = done
			contentType = "text/plain"
		}
		reqHeader.Set("Content-Length", strconv.Itoa(len(resp)))
		reqHeader.Set("Content-Type", contentType + "; charset=utf-8")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else if d.Nanoseconds() > 0 {
			//if they're going to be waiting more than 5 seconds, immediately
			//send back headers to prevent timeout
			if d.Seconds() > 3 {
				//send back headers but not a response yet
				w.WriteHeader(http.StatusOK)
				w.(http.Flusher).Flush()
			}

			<-time.After(d)
		}
		w.Write(resp)
	}
}

func parseDurationString(timeStr string) (time.Duration, error) {
	timeStr = strings.Trim(timeStr, " \r\n\t")
	var d time.Duration
	ms, err := strconv.ParseUint(timeStr, 10, 64)
	if err != nil {
		d, err = time.ParseDuration(timeStr)
		if err != nil {
			return d, err
		}
	} else {
		d = time.Duration(ms) * time.Millisecond
	}
	return d, nil
}
