package upupoo

//大写开头才能在外部包中访问
type Tag struct {
	CreateTime int    `json:"createTime"` //Jsontag 是映射json对应的字段
	TagId      int    `json:"tagId"`
	IsDelete   int    `json:"isDelete"`
	TagName    string `json:"tagName"`
}

type Tags struct {
	Data []Tag `json:"data"`
	Flag bool  `json:"flag"`
}
