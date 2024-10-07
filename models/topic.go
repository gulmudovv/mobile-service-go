package models

import "gorm.io/gorm"

type Topic struct {
	Id    uint   `gorm:"primarykey" json:"id" form:"-"`
	Title string `gorm:"unique;column:title" json:"title" form:"title"`
	Image string `gorm:"column:image" json:"image"`
}

func (t *Topic) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Topic{}).Count(&total)

	return total
}

func (t *Topic) Take(db *gorm.DB, limit int, offset int) interface{} {
	var topics []Topic

	db.Offset(offset).Limit(limit).Order("id").Find(&topics)

	return topics
}
