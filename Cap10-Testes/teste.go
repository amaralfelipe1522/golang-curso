package teste

import (
	"io/ioutil"
	"net/http"
)

func extrairSite(url string) <-chan string {
	c := make(chan string)
	go func(url string) {
		resp, _ := http.Get(url)
		html, _ := ioutil.ReadAll(resp.Body)
		c <- string(html)
	}(url)
	return c
}
