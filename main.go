package main

import (
	"encoding/json"
	"fmt"
	"github.com/tomasen/fcgi_client"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	dftNetwork = "tcp"
	dftAddress = "127.0.0.1:9000"
	dftScript  = "/status"
	dftKey     = "listen queue"
)

var (
	envDebug = os.Getenv("DEBUG") == "true"

	optNetwork string
	optAddress string
	optScript  string
	optKey     string
)

func exit(err *error) {
	if *err != nil {
		log.Println("exited with error:", (*err).Error())
		os.Exit(1)
	}
}

func argVal(idx int, out *string, dft string) {
	if idx < len(os.Args) {
		*out = os.Args[idx]
	}
	if len(*out) == 0 {
		*out = dft
	}
}

func main() {
	var err error
	defer exit(&err)

	argVal(1, &optKey, dftKey)
	argVal(2, &optNetwork, dftNetwork)
	argVal(3, &optAddress, dftAddress)
	argVal(4, &optScript, dftScript)

	var stats map[string]interface{}
	if stats, err = status(optNetwork, optAddress, optScript); err != nil {
		return
	}

	if stats[optKey] == nil {
		return
	}

	if val, ok := stats[optKey].(float64); ok {
		fmt.Printf("%d", int64(val))
	}
}

func status(network, address string, script string) (stats map[string]interface{}, err error) {
	var client *fcgiclient.FCGIClient
	if client, err = fcgiclient.Dial(network, address); err != nil {
		return
	}
	defer client.Close()

	var res *http.Response
	if res, err = client.Get(map[string]string{
		"SCRIPT_FILENAME": script,
		"SCRIPT_NAME":     script,
		"QUERY_STRING":    "json",
	}); err != nil {
		return
	}
	defer res.Body.Close()

	var buf []byte
	if buf, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	if envDebug {
		fmt.Printf("%s\n", buf)
	}

	if err = json.Unmarshal(buf, &stats); err != nil {
		return
	}

	return
}
