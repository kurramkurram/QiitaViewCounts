package data

type UserInfo struct {
	Id               string `json:"id"`
	Likes_count      int    `json:"likes_count"`
	Title            string `json:"title"`
	Page_views_count int    `json:"page_views_count"`
}
