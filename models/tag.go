package models

import "github.com/jaeyo/personal-archive/common"

type Tag struct {
	ID         int64  `gorm:"column:id;primarykey" json:"id"`
	Name       string `gorm:"column:name;type:varchar(60);not null" json:"name"`
	IsFavorite bool   `gorm:"column:is_favorite;type:integer;not null;default:0" json:"isFavorite"`
}

func (t *Tag) TableName() string {
	return "tag"
}

type Tags []*Tag

func (t Tags) ExtractTagNames() []string {
	names := []string{}
	for _, tag := range t {
		names = append(names, tag.Name)
	}
	return names
}

func (t Tags) DivideByContained(tagNames common.Strings) (Tags, Tags) {
	contained := Tags{}
	notContained := Tags{}
	for _, tag := range t {
		if tagNames.Contain(tag.Name) {
			contained = append(contained, tag)
		} else {
			notContained = append(notContained, tag)
		}
	}
	return contained, notContained
}
