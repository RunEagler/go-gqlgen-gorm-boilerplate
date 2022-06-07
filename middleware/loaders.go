package middleware

import (
	"context"
	"go-gqlgen-gorm-boilerplate/graph/dataloader"
	"go-gqlgen-gorm-boilerplate/graph/generated"
	"net/http"

	"github.com/jinzhu/gorm"
)

const loadersKey = "dataloadres"

type Loaders struct {
	TodosByUserID *generated.ComplexityRoot
}

func DataLoaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			TodosByUserID: dataloader.Todoloader(db),
			//SkillByID:       SkillLoader(db.back),
			//SkillsByFieldID:            dataloader.SkillsByFieldIDLoader(db),
			//InterestsByFieldID:         dataloader.InterestsLoader(db),
			//ArticlesBySkillID:          dataloader.ArticleLoader(db),
			//CodesBySkillID:             dataloader.CodeLoader(db),
			//IssuesBySkillID:            dataloader.IssueLoader(db),
			//ChildSkillsByParentSKillID: dataloader.SkillChildLoader(db),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
