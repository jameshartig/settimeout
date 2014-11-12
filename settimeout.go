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

var version = "0.3.0"
var favicon []byte
var index []byte
var robots []byte

var runningProcs = 0
var totalTimeoutRequests = 0
var totalTCPRequests = 0
var totalIndexRequests = 0

var http10StatusOK = []byte("HTTP/1.0 200 OK\r\n")
var http11StatusOK = []byte("HTTP/1.1 200 OK\r\n")

//plain-text
var invalidFormat = []byte("invalid_format")
var done = []byte("done")
var invalidCommand = []byte("Unknown command. Type 'info'\n")

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

var serverName = "settimeout/" + version

func main() {
	addr := flag.String("addr", ":80", "HTTP address to listen on (empty to disable)")
	tcpAddr := flag.String("tcpaddr", ":5103", "Socket address to listen on (empty to disable)")
	statsAddr := flag.String("statsaddr", "127.0.0.1:5104", "Socket address to listen on for stats (empty to disable)")
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

	go startHTTPServer(*addr)
	go startTCPServer(*tcpAddr)
	go startStatsTCPServer(*statsAddr)

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

func startStatsTCPServer(addr string) {
	if addr == "" {
		return
	}
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to start stats TCP server: " + err.Error())
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			// handle error
			continue
		}
		go func() {
			defer conn.Close()
			for {
				conn.SetReadDeadline(time.Now().Add(3 * time.Second))
				buf := bufio.NewReader(conn)
				str, err := buf.ReadString('\n')
				if err != nil {
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						//if its a i/o timeout then just continue the loop again
						continue
					}
					return
				}
				str = strings.Trim(str, " \r\n\t")
				switch str {
				case "close", "quit", "exit":
					return
				case "info":
					infos := make([]string, 4)
					infos[0] = "Running Requests: " + strconv.Itoa(runningProcs)
					infos[1] = "Total HTTP Index Requests: " + strconv.Itoa(totalIndexRequests)
					infos[2] = "Total TCP Requests: " + strconv.Itoa(totalTCPRequests)
					infos[3] = "Total HTTP Timeout Requests: " + strconv.Itoa(totalTimeoutRequests)
					conn.Write([]byte(strings.Join(infos, "\n")))
				default:
					conn.Write(invalidCommand)
				}
			}
		}()
	}
}

func startHTTPServer(addr string) {
	if addr == "" {
		return
	}
	s := &http.Server{
		Addr:        addr,
		Handler:     http.HandlerFunc(httpHandler),
		ReadTimeout: time.Duration(5) * time.Second,
	}
	err := s.ListenAndServe()
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
	totalTCPRequests++
	d, err := parseDurationString(str)
	if err != nil {
		conn.Write(invalidFormat)
		return
	}
	waitForTimeoutClose(conn, d, done)
}

func waitForTimeoutClose(conn net.Conn, timeout time.Duration, resp []byte) {
	if timeout.Nanoseconds() <= 0 {
		conn.Write(resp)
		return
	}

	runningProcs++
	closedCh := make(chan int)
	go func() {

		defer func() {
			runningProcs--
		}()

		singleChar := make([]byte, 1)
		for {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			_, err := conn.Read(singleChar)
			//if we got a read error and it wasn't a timeout then close the connection
			if netErr, ok := err.(net.Error); err != nil && (!ok || !netErr.Timeout()) {
				select {
				case closedCh <- 1:
				default:
				}
			}
			return
		}
	}()
	select {
	case <-closedCh:
		return
	case <-time.After(timeout):
		conn.Write(done)
	}
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	reqHeader := w.Header()
	reqHeader.Set("Access-Control-Allow-Origin", "*")
	reqHeader.Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	reqHeader.Set("Server", serverName)
	if req.Method != "GET" && req.Method != "POST" {
		//assume its the browser just checking headers
		w.WriteHeader(http.StatusOK)
		return
	}
	str := strings.TrimLeft(req.URL.Path, "/")

	switch str {
	case "": //index page
		reqHeader.Set("Cache-Control", "max-age=3600") //1 hour
		w.Write(index)
		totalIndexRequests++
	case "favicon.ico":
		reqHeader.Set("Cache-Control", "max-age=31536000") //1 year
		w.Write(favicon)
	case "robots.txt":
		reqHeader.Set("Cache-Control", "max-age=31536000") //1 year
		w.Write(robots)
	default:
		totalTimeoutRequests++
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

		reqHeader.Set("Content-Type", contentType+"; charset=utf-8")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		hj, ok := w.(http.Hijacker)
		if !ok {
			log.Println("Connection doesn't support hijacking for some reason")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//now hijack the connection so we can see if they disconnected
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			log.Println("Failed to hijack http connection: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		reqHeader.Set("Content-Length", strconv.Itoa(len(resp)))
		reqHeader.Set("Connection", "close")

		//send back headers
		statusLine := http10StatusOK
		if req.ProtoAtLeast(1, 1) {
			statusLine = http11StatusOK
		}
		bufrw.Write(statusLine)
		reqHeader.Write(bufrw)

		//if they're going to be waiting more than 3 seconds immediately
		//send back headers to prevent timeout
		if d.Seconds() > 3 {
			bufrw.Flush()
		}
		waitForTimeoutClose(conn, d, resp)
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
