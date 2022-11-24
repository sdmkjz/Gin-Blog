package model

// 分类
type Category struct {
	ID   uint   `gorm:"primarykey" json:"id""`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
