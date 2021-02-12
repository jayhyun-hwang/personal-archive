package repositories

import (
	"github.com/jaeyo/personal-archive/internal"
	"github.com/jaeyo/personal-archive/models"
	"github.com/pkg/errors"
	"sync"
)

type TagRepository interface {
	UpdateTag(tag, newTag string) error
	FindByNames(names []string) (models.Tags, error)
}

type tagRepository struct {
	database *internal.DB
}

var GetTagRepository = func() func() TagRepository {
	var instance TagRepository
	var once sync.Once

	return func() TagRepository {
		once.Do(func() {
			instance = &tagRepository{
				database: internal.GetDatabase(),
			}
		})
		return instance
	}
}()

func (r *tagRepository) UpdateTag(tag, newTag string) error {
	query := r.database.
		Model(&models.Tag{}).
		Where("name = ?", tag).
		Update("name", newTag)
	if query.RowsAffected <= 0 {
		return errors.New("no row affected")
	}
	return query.Error
}

func (r *tagRepository) FindByNames(names []string) (models.Tags, error) {
	var tags []*models.Tag
	err := r.database.
		Where("name IN ?", names).
		Find(&tags).Error
	return tags, err
}
