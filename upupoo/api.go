package upupoo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetTags() (result Tags, err error) {
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

	// // 用字符串声明的方式
	// tmpJson = `({"data":[{"createTime":1483243995000,"tagId":0,"isDelete":1,"tagName":"全部"},{"createTime":1483519853000,"tagId":2,"isDelete":1,"tagName":"游戏风象"},{"createTime":1483519857000,"tagId":3,"isDelete":1,"tagName":"鬼畜"},{"createTime":1483596223000,"tagId":4,"isDelete":1,"tagName":"绝对领域"},{"createTime":1483596225000,"tagId":5,"isDelete":1,"tagName":"女朋友"},{"createTime":1483883866000,"tagId":6,"isDelete":1,"tagName":"萌物"},{"createTime":1483883879000,"tagId":7,"isDelete":1,"tagName":"美の風景"},{"createTime":1483883899000,"tagId":9,"isDelete":1,"tagName":"神曲"},{"createTime":1483883912000,"tagId":11,"isDelete":1,"tagName":"解忧杂货铺"},{"createTime":1484084578000,"tagId":12,"isDelete":1,"tagName":"ACG动漫"},{"createTime":1484084588000,"tagId":13,"isDelete":1,"tagName":"艺术感知"},{"createTime":1484328498000,"tagId":1,"isDelete":1,"tagName":"抖不停"},{"createTime":1484333623000,"tagId":10,"isDelete":1,"tagName":"科幻世界"}],"flag":true})`
	// runes := []rune(tmpJson)
	// // 处理多余的括号
	// tmpJson = string(runes[1 : len(runes)-1])

	// 反序列化
	err = json.Unmarshal([]byte(tmpJson), &result)
	return
}
