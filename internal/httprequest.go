package internal

import (
	"fmt"
	"github.com/fireflyhoo/huobi_golang/config"
	"github.com/fireflyhoo/huobi_golang/logging/perflogger"
	"io/ioutil"
	"net/http"
	"strings"
)
import "net/url"

func HttpGet(urlStr string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()
	var transport *http.Transport
	if config.HttpProxy {
		proxy, err := url.Parse("http://" + config.ProxyHost + ":" + config.ProxyPort)
		fmt.Print(err)
		transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	} else {
		transport = &http.Transport{}
	}

	client := &http.Client{Transport: transport}

	rep, _ := http.NewRequest("GET", urlStr, nil)
	resp, err := client.Do(rep)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("GET", urlStr)

	return string(result), err
}

func HttpPost(urlStr string, body string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()
	var transport *http.Transport
	if config.HttpProxy {
		proxy, err := url.Parse("http://" + config.ProxyHost + ":" + config.ProxyPort)
		fmt.Print(err)
		transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	} else {
		transport = &http.Transport{}
	}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("POST", urlStr)
	return string(result), err
}
