package exporter

import (
	"../data"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ToCsv(userInfos []data.UserInfo) {
	file, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("can not create csv file")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"title", "like_count", "page_views_count"})
	for _, user := range userInfos {
		likesCount := strconv.Itoa(user.Likes_count)
		pageViewsCount := strconv.Itoa(user.Page_views_count)
		writer.Write([]string{user.Title, likesCount, pageViewsCount})
	}
	writer.Flush()
}
