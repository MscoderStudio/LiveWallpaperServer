package upupoo

import (
	"LiveWallpaperServer/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetWallpapers() {
	//请求服务器获取json
	res, err := http.Get("http://wallpaper.upupoo.com/async/asyncSearch--3-2-2-2.htm?callback=&_=1537975783716")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	//多了两个括号需要移除掉
	tmpJson := string(body[1 : len(body)-1])
	// 反序列化
	var temp SearchResult
	err = json.Unmarshal([]byte(tmpJson), &temp)
}

func GetTags() (result []model.Tag, err error) {
	//请求服务器获取json
	res, err := http.Get("http://wallpaper.upupoo.com/async/getTags.htm?callback=")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	//多了两个括号需要移除掉
	tmpJson := string(body[1 : len(body)-1])
	// 反序列化
	var temp UTags
	err = json.Unmarshal([]byte(tmpJson), &temp)

	// UTags转换成标准Tags
	for _, v := range temp.Data {
		result = append(result, model.Tag{ID: v.TagId, Name: v.TagName})
	}
	return
}
