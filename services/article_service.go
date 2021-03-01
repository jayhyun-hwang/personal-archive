package services

import (
	"fmt"
	"github.com/jaeyo/personal-archive/common"
	"github.com/jaeyo/personal-archive/models"
	"github.com/jaeyo/personal-archive/repositories"
	"github.com/jaeyo/personal-archive/services/generators"
	"github.com/pkg/errors"
	"sync"
)

type ArticleService interface {
	Initialize()
	CreateByURL(url string, tags common.Strings) (*models.Article, error)
	Search(keyword string, offset, limit int) ([]*models.Article, int64, error)
	UpdateTitle(id int64, newTitle string) error
	UpdateTags(id int64, tags common.Strings) error
	UpdateContent(id int64, content string) error
	DeleteByIDs(ids []int64) error
}

type articleService struct {
	articleGenerator        generators.ArticleGenerator
	articleRepository       repositories.ArticleRepository
	articleSearchRepository repositories.ArticleSearchRepository
	tagRepository           repositories.TagRepository
}

var GetArticleService = func() func() ArticleService {
	var once sync.Once
	var instance ArticleService
	return func() ArticleService {
		once.Do(func() {
			instance = &articleService{
				articleGenerator:        generators.GetArticleGenerator(),
				articleRepository:       repositories.GetArticleRepository(),
				articleSearchRepository: repositories.GetArticleSearchRepository(),
				tagRepository:           repositories.GetTagRepository(),
			}
		})
		return instance
	}
}()

func (s *articleService) Initialize() {
	if err := s.articleSearchRepository.Initialize(); err != nil {
		panic(err)
	}
}

func (s *articleService) CreateByURL(url string, tagNames common.Strings) (*models.Article, error) {
	article, err := s.articleGenerator.NewArticle(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate new article")
	}

	tags, err := s.tagRepository.FindByNames(tagNames)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find tags")
	}

	toBeCreatedTagNames := tagNames.FilterNotContained(tags.ExtractTagNames())
	for _, tagName := range toBeCreatedTagNames {
		tags = append(tags, &models.Tag{
			Name:       tagName,
			IsFavorite: false,
		})
	}

	article.Tags = tags

	if err = s.articleRepository.Save(article); err != nil {
		return nil, errors.Wrap(err, "failed to save article")
	}

	return article, nil
}

func (s *articleService) Search(keyword string, offset, limit int) ([]*models.Article, int64, error) {
	ids, err := s.articleSearchRepository.Search(keyword)
	if err != nil {
		return nil, -1, errors.Wrap(err, "failed to search")
	}

	articles, cnt, err := s.articleRepository.FindByIDsWithPage(ids, offset, limit)
	if err != nil {
		return nil, -1, errors.Wrap(err, "failed to find article by ids")
	}

	return articles, cnt, nil
}

func (s *articleService) UpdateTitle(id int64, newTitle string) error {
	exist, err := s.articleRepository.ExistByTitle(newTitle)
	if err != nil {
		return errors.Wrap(err, "failed to check exist by title")
	} else if exist {
		return fmt.Errorf("title %s already exists", newTitle)
	}

	article, err := s.articleRepository.GetByID(id)
	if err != nil {
		return errors.Wrap(err, "failed to get article")
	}

	article.Title = newTitle

	if err := s.articleRepository.Save(article); err != nil {
		return errors.Wrap(err, "failed to save article")
	}
	return nil
}

func (s *articleService) UpdateTags(id int64, tags common.Strings) error {
	article, err := s.articleRepository.GetByID(id)
	if err != nil {
		return errors.Wrap(err, "failed to get article")
	}

	updatedTags := models.Tags{}

	toBePreserved := article.Tags.Filter(tags)
	updatedTags = append(updatedTags, toBePreserved...)

	toBeAdded := tags.FilterNotContained(article.Tags.ExtractTagNames())
	for _, tagName := range toBeAdded {
		updatedTags = append(updatedTags, &models.Tag{
			Name:       tagName,
			IsFavorite: false,
		})
	}

	if err := s.articleRepository.SaveTags(article, updatedTags); err != nil {
		return errors.Wrap(err, "failed to save tags")
	}

	return nil
}

func (s *articleService) UpdateContent(id int64, content string) error {
	article, err := s.articleRepository.GetByID(id)
	if err != nil {
		return errors.Wrap(err, "failed to get article)")
	}

	article.Content = content

	if err := s.articleRepository.Save(article); err != nil {
		return errors.Wrap(err, "failed to save article")
	}
	return nil
}

func (s *articleService) DeleteByIDs(ids []int64) error {
	articles, err := s.articleRepository.FindByIDs(ids)
	if err != nil {
		return errors.Wrap(err, "failed to find articles")
	} else if len(ids) != len(articles) {
		return fmt.Errorf("invalid ids: %v", ids)
	}

	// TODO: delete tag association
	//if err := s.articleTagRepository.DeleteByIDs(articles.ExtractTagIDs()); err != nil {
	//	return errors.Wrap(err, "failed to delete article tag by ids")
	//}

	if err := s.articleRepository.DeleteByIDs(ids); err != nil {
		return errors.Wrap(err, "failed to delete article by ids")
	}
	return nil
}
