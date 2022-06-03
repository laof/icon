package api

import (
	"fmt"
	"icon/tools"
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {

}

func IconHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")

	if strings.HasPrefix(url, "http") {

		var cs = make(chan string)

		go tools.GetLink(url, cs)

		link := <-cs

		if link == "" {
			w.Write([]byte("parse url failed"))
			return
		}

		// data, _ := getRemote(link)

		// img := fmt.Sprintf("<img src='%s'/>", link)

		w.Write([]byte(link))
		return

	}

	w.Write([]byte("parameter error"))

}

func getRemote(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {

		fmt.Println(err.Error())
		// 如果有错误返回错误内容
		return nil, err
	}
	// 使用完成后要关闭，不然会占用内存
	defer res.Body.Close()
	// 读取字节流
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, err
}
