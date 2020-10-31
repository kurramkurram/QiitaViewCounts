package qiita

import (
	"../data"
	"../exporter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var token = ""

func GetQiitaViews() {
	url := "https://qiita.com/api/v2/authenticated_user/items?page=1&per_page=20"
	resp, err := doHttpRequest(url)
	defer resp.Body.Close()

	if err != nil {
		log.Println("can not get response from " + url)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can not read body")
		return
	}

	var pageInfos []data.PageInfo
	if err = json.Unmarshal(body, &pageInfos); err != nil {
		log.Println("can not unmarshal json user info")
		return
	}

	index := 0
	for _, page := range pageInfos {
		url = "https://qiita.com/api/v2/items/" + page.Id
		resp, err := doHttpRequest(url)
		defer resp.Body.Close()

		if err != nil {
			log.Println("can not get response from " + url)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("can not read body")
			return
		}

		if err := json.Unmarshal(body, &page); err != nil {
			log.Println("can not unmarshal json page info")
			return
		}
		pageInfos[index].Page_views_count = page.Page_views_count
		index += 1
		fmt.Println(page.Title, page.Page_views_count, page.Likes_count)
	}
	exporter.ToCsv(pageInfos)
}

func doHttpRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("can not create request")
		return nil, err
	}

	if token == "" {
		buf, err := ioutil.ReadFile("token.txt")
		if err != nil {
			log.Println("can not file open")
		}
		token = string(buf)
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer " + token)

	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, err
}
