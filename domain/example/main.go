package example

type ArticleRepository interface {
	FetchByID(id int) (*Article, error)
}

type Article struct {
	ID    int
	Title string
}
