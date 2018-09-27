package upupoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetTags() (result Tags) {
	//请求服务器获取json
	res, err := http.Get("http://wallpaper.upupoo.com/async/getTags.htm?callback=")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	//反序列化
	if err := json.Unmarshal(body, &result); err != nil {
		//反序列化错误
		fmt.Printf("%s", body)
		log.Fatal(err)
	}
	return
}
