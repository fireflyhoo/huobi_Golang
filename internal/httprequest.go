package internal

import (
	"github.com/fireflyhoo/huobi_golang/config"
	"github.com/fireflyhoo/huobi_golang/logging/perflogger"
	"io/ioutil"
	"net/http"
	"strings"
)
import "net/url"

func HttpGet(urlstr string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	proxy, err := url.Parse("http://" + config.ProxyHost + ":" + config.ProxyPort)
	//fmt.Print(proxy)
	transport := &http.Transport{Proxy: http.ProxyURL(proxy)}
	client := &http.Client{Transport: transport}

	rep, _ := http.NewRequest("GET", urlstr, nil)
	resp, err := client.Do(rep)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("GET", urlstr)

	return string(result), err
}

func HttpPost(url string, body string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("POST", url)

	return string(result), err
}
