package settimeout

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	version    = "0.5.0"
	serverName = "settimeout/" + version

	//plain-text
	done = []byte("done")

	//int
	one = []byte("1")

	//json
	jsonTrue = []byte("true")

	//js
	jsTrue = []byte("_settimeoutio=true")

	//css
	cssShow = []byte(".settimeoutio {display: block;}")

	//gif
	gifPixel []byte

	//callback
	emptyCallback = []string{""}
)

func init() {
	// from http://probablyprogramming.com/2009/03/15/the-tiniest-gif-ever
	gifPixel, _ = base64.StdEncoding.DecodeString(`R0lGODlhAQABAIABAP///wAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==`)
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, req *http.Request) {
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
	case exists("gif"):
		resp = gifPixel
		contentType = "image/gif"
	default:
		resp = done
		contentType = "text/plain"
	}

	reqHeader.Set("Content-Type", contentType+"; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqHeader.Set("Content-Length", strconv.Itoa(len(resp)))
	reqHeader.Set("Connection", "close")

	w.WriteHeader(http.StatusOK)

	select {
	//case <-req.Context().Done():
	case <-time.After(d):
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
