package model

import "time"

// 用户表
type NfUser struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Password  string    `gorm:"column:password" json:"-"`
	RealName  string    `gorm:"column:real_name" json:"real_name"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (p *NfUser) TableName() string {
	return "nf_user"
}
