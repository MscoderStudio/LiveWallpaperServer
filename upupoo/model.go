package upupoo

type Tag struct {
	CreateTime string `json:"createTime"`
	TagId      string `json:"tagId"`
	IsDelete   string `json:"isDelete"`
	TagName    string `json:"tagName"`
}

type Tags struct {
	Data []Tag `json:"data"`
	Flag bool  `json:"flag"`
}
