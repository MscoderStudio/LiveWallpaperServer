package upupoo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getTags() {
	res, err := http.Get("http://wallpaper.upupoo.com/async/getTags.htm?callback=")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
