package dataloader

import (
	"github.com/jinzhu/gorm"
	"go-gqlgen-gorm-boilerplate/graph/generated"
	"go-gqlgen-gorm-boilerplate/graph/model"
	"time"
)

func CodeLoader(db *gorm.DB) *generated.TodosByUserID {
	return generated.NewTodosByUserID(generated.TodosByUserIDConfig{
		Wait:     1 * time.Millisecond,
		MaxBatch: 1000,
		Fetch: func(userID []int) ([]model.Todo, []error) {
			var todos []model.Todo
			todos := make([]model.Todo, 0)
			db.Where("user_id IN (?)", userID).Find(&todos)

			for _, skillCodeModel := range skillCodeModels {
				skillCodes = append(skillCodes, &dto.Code{
					ID:          skillCodeModel.ID,
					Code:        skillCodeModel.Code.String,
					Language:    skillCodeModel.Language,
					Pos:         int(skillCodeModel.Pos),
					Description: util.PtrString(skillCodeModel.Description.String),
					UpdatedAt:   skillCodeModel.UpdatedAt,
				})
			}
			return skillCodes, nil
		},
	})
}
