package repositories

import (
	"github.com/jaeyo/personal-archive/internal"
	"github.com/jaeyo/personal-archive/models"
	"github.com/pkg/errors"
	"sync"
)

type ArticleTagRepository interface {
	UpdateTag(tag, newTag string) error
	FindCounts() ([]*models.ArticleTagCountDTO, error)
	Delete(articleTags models.ArticleTags) error
	DeleteByIDs(ids []int64) error
}

type articleTagRepository struct {
	database *internal.DB
}

func newArticleTagRepository(app appIface) ArticleTagRepository {
	return &articleTagRepository{}
}

var GetArticleTagRepository = func() func() ArticleTagRepository {
	var instance ArticleTagRepository
	var once sync.Once

	return func() ArticleTagRepository {
		once.Do(func() {
			instance = &articleTagRepository{
				database: internal.GetDatabase(),
			}
		})
		return instance
	}
}()

func (r *articleTagRepository) UpdateTag(tag, newTag string) error {
	query := r.database.
		Model(&models.ArticleTag{}).
		Where("tag = ?", tag).
		Update("tag", newTag)
	if query.RowsAffected <= 0 {
		return errors.New("no row affected")
	}
	return query.Error
}

func (r *articleTagRepository) FindCounts() ([]*models.ArticleTagCountDTO, error) {
	var counts []*models.ArticleTagCountDTO
	err := r.database.
		Model(&models.ArticleTag{}).
		Select("tag", "count(*) AS cnt").
		Group("tag").
		Order("tag ASC").
		Find(&counts).Error
	return counts, err
}

func (r *articleTagRepository) Delete(articleTags models.ArticleTags) error {
	return r.database.Delete(&articleTags).Error
}

func (r *articleTagRepository) DeleteByIDs(ids []int64) error {
	return r.database.Where("id IN ?", ids).Delete(&models.ArticleTag{}).Error
}
