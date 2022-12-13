package repository

import (
	"context"

	"gorm.io/gorm"
)

type DAO[T any] interface {
	FromEntity(e T) DAO[T]
	ToEntity() T
}
type IRepository[T any] interface {
	Find(id string) (*T, error)
	FindById(id string) (*T, error)
	FreeFind() *gorm.DB
	First() (*T, error)
	Insert(entry *T) error
}
type Repository[M DAO[E], E any] struct {
	db *gorm.DB
}

func NewRepository[M DAO[E], E any](ctx context.Context) *Repository[M, E] {
	db := ctx.Value("db").(gorm.DB)
	return &Repository[M, E]{
		db: &db,
	}
}
func (r *Repository[M, E]) Insert(ctx context.Context, entity *E) error {
	var start M
	model := start.FromEntity(*entity).(M)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	*entity = model.ToEntity()
	return nil
}
func (r *Repository[M, E]) FindByID(ctx context.Context, id uint) (*E, error) {
	// retrieve a record by id from a database
	var model M
	err := r.db.WithContext(ctx).First(&model, id).Error
	{
		return nil, err
	}
	// handle error

	// map data into Entity
	rv := model.ToEntity()
	return &rv, nil
}

func (r *Repository[M, E]) Find(ctx context.Context, specification Specification, params ...any) ([]E, error) {
	// retreive reords by some criteria
	var models []M
	if err := r.db.WithContext(ctx).Where(specification.Query()).Find(&models).Error; err != nil {
		return []E{}, err
	}
	// mapp all records into Entities
	result := make([]E, 0, len(models))
	for _, row := range models {
		result = append(result, row.ToEntity())
	}

	return result, nil
}
