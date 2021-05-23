package puppetweb

type Article struct {
	Title  string `json:"Title"`
	Digest string `json:"Digest"`
	Cover  string `json:"Cover"`
	Url    string `json:"Url"`
}

type Subscribe struct {
	UserName       string    `json:"UserName"`
	MPArticleCount int       `json:"MPArticleCount"`
	MPArticleList  []Article `json:"MPArticleList"`
	Time           int       `json:"Time"`
	NickName       string    `json:"NickName"`
}
