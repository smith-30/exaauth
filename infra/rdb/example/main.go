package example

import (
	"github.com/jinzhu/gorm"
	domain_example "github.com/smith-30/petit/domain/example"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(d *gorm.DB) domain_example.ArticleRepository {
	a := &articleRepository{
		db: d,
	}
	return a
}

func (a *articleRepository) FetchByID(id int) (*domain_example.Article, error) {
	return nil, nil
}
