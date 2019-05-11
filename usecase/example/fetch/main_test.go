package fetch

import (
	"os"
	"reflect"
	"testing"

	"github.com/smith-30/exaauth/infra/rdb"

	"github.com/jinzhu/gorm"
	domain_example "github.com/smith-30/exaauth/domain/example"
)

func TestMain(m *testing.M) {
	rdb.InitFakeRDB()
	code := m.Run()
	os.Exit(code)
}

func Test_usecase_Exec(t *testing.T) {
	type fields struct {
		tx          *gorm.DB
		articleRepo domain_example.ArticleRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *domain_example.Article
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &usecase{
				tx:          tt.fields.tx,
				articleRepo: tt.fields.articleRepo,
			}
			got, err := a.Exec()
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
