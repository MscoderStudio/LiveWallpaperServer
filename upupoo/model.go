package upupoo

//大写开头才能在外部包中访问
type UTag struct {
	CreateTime int    `json:"createTime"` //Jsontag 是映射json对应的字段
	TagId      int    `json:"tagId"`
	IsDelete   int    `json:"isDelete"`
	TagName    string `json:"tagName"`
}

type UTags struct {
	Data []UTag `json:"data"`
	Flag bool   `json:"flag"`
}
