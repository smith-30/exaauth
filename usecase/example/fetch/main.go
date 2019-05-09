package fetch

import (
	"github.com/jinzhu/gorm"
	domain_example "github.com/smith-30/petit/domain/example"
	"github.com/smith-30/petit/infra/rdb"
	rdb_example "github.com/smith-30/petit/infra/rdb/example"
)

type Fetch interface {
	Exec() (*domain_example.Article, error)
}

type usecase struct {
	tx          *gorm.DB
	articleRepo domain_example.ArticleRepository
}

func NewFetchUsecase() (Fetch, error) {
	tx := rdb.RDB.Conn().Begin()
	a := &usecase{
		tx:          tx,
		articleRepo: rdb_example.NewArticleRepository(tx),
	}

	return a, tx.Error
}

func (a *usecase) Exec() (*domain_example.Article, error) {
	if _, err := a.articleRepo.FetchByID(1); err != nil {
		a.tx.Rollback()
		return nil, err
	}

	return nil, a.tx.Commit().Error
}
