package exporter

import (
	"../data"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ToCsv(pageInfos []data.PageInfo) {
	file, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("can not create csv file")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"title", "like_count", "page_views_count"})
	for _, page := range pageInfos {
		likesCount := strconv.Itoa(page.Likes_count)
		pageViewsCount := strconv.Itoa(page.Page_views_count)
		writer.Write([]string{page.Title, likesCount, pageViewsCount})
	}
	writer.Flush()
}
