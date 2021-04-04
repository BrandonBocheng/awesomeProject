package common

type Response struct {
	Code     uint                `json:"code"`
	Msg      string              `json:"msg"`
	NewsList []map[string]string `json:"newsList"`
}

type NewsData struct {
	Timestamp uint
	Source    string
	Title     string
	Body      string
	Types     []string
}
