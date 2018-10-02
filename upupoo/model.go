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

type SearchResult struct {
	Flag bool `json:"flag"`
	Data SearchResultData
}

type UWallpaper struct {
	AuthorId          int
	AuthorName        string
	CreateTime        int
	DownCount         int
	DownStr           string
	DownUrl           string
	Heat              int
	Icon              string
	IsOriginal        int
	PaperId           int
	PaperImg          string
	PaperName         string
	PaperTypeId       int
	PaperUrl          string
	PaperViewCountStr int
	ReprintUrl        string
	Size              int
	UpdateTime        int
	VirtualDownCount  int
}

type SearchResultData struct {
	EndPage   int
	Limit     int
	Offset    int
	ParamMap  ParamMap
	Rows      []UWallpaper
	StartPage int
	Total     int
	TotalPage int
}

type ParamMap struct {
	End         int
	PaperName   string
	PaperTypeId int
	Sort        int
	Start       int
	TagId       int
}
