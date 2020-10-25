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

	var userInfos []data.UserInfo
	if err = json.Unmarshal(body, &userInfos); err != nil {
		log.Println("can not unmarshal json user info")
		return
	}

	index := 0
	for _, user := range userInfos {
		url = "https://qiita.com/api/v2/items/" + user.Id
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

		if err := json.Unmarshal(body, &user); err != nil {
			log.Println("can not unmarshal json page info")
			return
		}
		userInfos[index].Page_views_count = user.Page_views_count
		index += 1
		fmt.Println(user.Title, user.Page_views_count, user.Likes_count)
	}
	exporter.ToCsv(userInfos)
}

func doHttpRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("can not create request")
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer ")

	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, err
}
