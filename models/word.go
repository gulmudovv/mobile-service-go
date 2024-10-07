package models

import "gorm.io/gorm"

type Word struct {
	Id         uint   `gorm:"primarykey" json:"id"`
	Arabic     string `gorm:"column:arabic" json:"arabic" form:"arabic"`
	Russian    string `gorm:"unique;column:russian" json:"russian" form:"russian"`
	Transcript string `gorm:"column:transcript" json:"transcript" form:"transcript"`
	Audio      string `gorm:"column:audio" json:"audio"`
	Image      string `gorm:"column:image" json:"image"`
	TopicId    string `gorm:"column:topic_id" json:"topic_id" form:"topic_id"`
}

func (w *Word) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Word{}).Count(&total)

	return total
}

func (w *Word) Take(db *gorm.DB, limit int, offset int) interface{} {
	var words []Word

	db.Offset(offset).Order("id DESC").Limit(limit).Find(&words)

	return words
}
