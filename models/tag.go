package models

type Tag struct {
	ID         int64  `gorm:"column:id;primarykey" json:"id"`
	Name       string `gorm:"column:name;type:varchar(60);not null" json:"name"`
	IsFavorite bool   `gorm:"column:is_favorite;type:integer;not null;default:0" json:"isFavorite"`
}

func (t *Tag) TableName() string {
	return "tag"
}
