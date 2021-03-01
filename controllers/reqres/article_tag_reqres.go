package reqres

import "github.com/jaeyo/personal-archive/models"

type ArticleTagCountsResponse struct {
	OK            bool                  `json:"ok"`
	TagCounts     []*models.TagCountDTO `json:"tagCounts"`
	UntaggedCount int64                 `json:"untaggedCount"`
	AllCount      int64                 `json:"allCount"`
}

type UpdateTagRequest struct {
	Tag string `json:"tag"`
}
