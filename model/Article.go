package model

import "gorm.io/gorm"

// 文章
type Article struct {
	//`gorm:"foreignKey:ID"`
	Category Category `gorm:"ForeignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title""`
	Desc    string `gorm:"type:varchar(200)" json:"desc""`
	Cid     uint   `gorm:"type:uint;not null;" json:"cid"`
	Content string `gorm:"type:longtext" json:"content""`
	Img     string `gorm:"type:varchar(100)" json:"img""`
}
