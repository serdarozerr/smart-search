package search

type Response struct {
	Snippet string
	Url     string
	Title   string
}

type ResponseList []Response
