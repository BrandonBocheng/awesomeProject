package main

import (
	"awesomeProject/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	urlFormat := "http://api.tianapi.com/topnews/index?key=%s"
	url := fmt.Sprintf(urlFormat, common.APP_KEY)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	respJson := &common.Response{}
	err := json.Unmarshal(body, respJson)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// fmt.Println(*respJson)

	newsList := respJson.NewsList

	dataList := []common.NewsData{}

	for _, news := range newsList {
		var curData = common.NewsData{}

		ctimeVal, ok1 := news["ctime"]
		if ok1 {
			timeStamp, _ := time.Parse("2006-01-02 03:04:05 PM", ctimeVal) // String时间转化为unix时间戳
			curData.Timestamp = uint(timeStamp.Unix())
		} else {
			fmt.Println("提取时间戳失败")
		}

		sourceVal, ok2 := news["source"]
		if ok2 {
			curData.Source = sourceVal
		} else {
			fmt.Println("提取新闻源失败")
		}

		titleVal, ok3 := news["title"]
		if ok3 {
			curData.Title = titleVal
		} else {
			fmt.Println("提取新闻标题失败")
		}

		descVal, ok4 := news["description"]
		if ok4 {
			curData.Body = descVal
		} else {
			fmt.Println("提取新闻描述失败")
		}

		dataList = append(dataList, curData)
		fmt.Println(dataList)
	}

	fmt.Println(dataList)
}
