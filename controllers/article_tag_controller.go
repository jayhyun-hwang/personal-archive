package controllers

import (
	"github.com/jaeyo/personal-archive/common/http"
	"github.com/jaeyo/personal-archive/controllers/reqres"
	"github.com/jaeyo/personal-archive/repositories"
	"github.com/labstack/echo/v4"
)

type ArticleTagController struct {
	tagRepository     repositories.TagRepository
	articleRepository repositories.ArticleRepository
}

func NewArticleTagController() *ArticleTagController {
	return &ArticleTagController{
		tagRepository:     repositories.GetTagRepository(),
		articleRepository: repositories.GetArticleRepository(),
	}
}

func (c *ArticleTagController) Route(e *echo.Echo) {
	e.GET("/apis/article-tags", http.Provide(c.FindArticleTagCounts))
	e.PUT("/apis/article-tags/tag/:tag", http.Provide(c.UpdateTag))
}

func (c *ArticleTagController) FindArticleTagCounts(ctx http.ContextExtended) error {
	// TODO IMME
	articleTagCounts, err := c.articleTagRepository.FindCounts()
	if err != nil {
		return ctx.InternalServerError(err, "failed to find article tag counts")
	}

	untaggedCount, err := c.articleRepository.GetUntaggedCount()
	if err != nil {
		return ctx.InternalServerError(err, "failed to get untagged count")
	}

	allCount, err := c.articleRepository.GetAllCount()
	if err != nil {
		return ctx.InternalServerError(err, "failed to get all count")
	}

	return ctx.Success(reqres.ArticleTagCountsResponse{
		OK:               true,
		ArticleTagCounts: articleTagCounts,
		UntaggedCount:    untaggedCount,
		AllCount:         allCount,
	})
}

func (c *ArticleTagController) UpdateTag(ctx http.ContextExtended) error {
	tag := ctx.ParamStr("tag")

	var req reqres.UpdateTagRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequestf("invalid request body: %s", err.Error())
	}

	if err := c.tagRepository.UpdateTag(tag, req.Tag); err != nil {
		return ctx.InternalServerError(err, "failed to update tag")
	}

	return ctx.Success(http.SuccessResponse{OK: true})
}
