package models

// import "gorm.io/gorm"

// type Paginator struct {
// }

// func (p *Paginator) Count(db *gorm.DB) int64 {
// 	var total int64
// 	db.Model(&Topic{}).Count(&total)

// 	return total
// }

// func (p *Paginator) Take(db *gorm.DB, limit int, offset int) interface{} {
// 	var topics []Topic

// 	db.Offset(offset).Limit(limit).Find(&topics)

// 	return topics
// }
