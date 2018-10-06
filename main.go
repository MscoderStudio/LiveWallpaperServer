package main

import (
	"LiveWallpaperServer/upupoo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

func getTags(writer http.ResponseWriter, request *http.Request) {
	tags, err := upupoo.GetTags()
	json, err := json.Marshal(tags)
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}

func getWallpapers(writer http.ResponseWriter, request *http.Request) {
	tag := request.URL.Query().Get("tag")
	if tag == "" {
		tag = "0"
	}
	sort := request.URL.Query().Get("sort")
	if sort == "" {
		sort = "0"
	}
	page := request.URL.Query().Get("page")
	if page == "" {
		page = "0"
	}

	result, err := upupoo.GetWallpapers(tag, sort, page)
	json, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}

func getSorts(writer http.ResponseWriter, request *http.Request) {
	sorts, err := upupoo.GetSorts()
	json, err := json.Marshal(sorts)
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}

func loadConfig() {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(currentDir, "config.json")
	fmt.Printf("配置文件：%s\r\n", currentDir)
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func main() {
	loadConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/tags", getTags)
	mux.HandleFunc("/sorts", getSorts)
	mux.HandleFunc("/wallpapers", getWallpapers)
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("本项目属于技术交流，不承担任何版权责任")
	fmt.Println("不使用时关闭本窗口释放多余的性能开销")
	fmt.Println("巨应工作室 qq交流群: 641405255")
	fmt.Println("源码地址：https://github.com/MscoderStudio/LiveWallpaperServer")
	fmt.Println("《巨应动态壁纸》 第三方数据解析器运行中....")
	server.ListenAndServe()
}
