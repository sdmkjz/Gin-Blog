package model

// 个人信息
type Profile struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(20)" json:"name"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	QQchat    string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat    string `gorm:"type:varchar(100)" json:"wechat"`
	Weibo     string `gorm:"type:varchar(200)" json:"weibo"`
	Bili      string `gorm:"type:varchar(200)" json:"bili"`
	Email     string `gorm:"type:varchar(200)" json:"email"`
	Img       string `gorm:"type:varchar(200)" json:"img"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"`
}
