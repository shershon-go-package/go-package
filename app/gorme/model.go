/**
 * @Description 定义model
 **/
package gorme

import (
	"database/sql"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	NickName     string         `gorm:"type:varchar(20);not null;default:'';comment:昵称"`
	Age          uint8          `gorm:"size:4;comment:年龄"`
	Phone        string         `gorm:"type:char(11);index:un_phone;comment:手机号"`
	MemberNumber string         `gorm:"type:varchar(20);index:un_phone;comment:会员编号"`
	Birthday     sql.NullString `gorm:"type:varchar(10);comment:生日"`
	ActivatedAt  sql.NullTime   `gorm:"comment:激活时间"`
}

// 用户地址表
type UserAddress struct {
	gorm.Model
	Uid      uint   `gorm:"comment:用户id"`
	Province string `gorm:"type:varchar(10);not null;default:'';comment:所属省份"`
	City     string `gorm:"type:varchar(20);not null;default:'';comment:所属城市"`
	Area     string `gorm:"type:varchar(20);not null;default:'';comment:所属区、县"`
	Detail   string `gorm:"type:varchar(255);not null;default:'';comment:具体地址"`
}

type UserList struct {
	gorm.Model
	NickName     string         `gorm:"type:varchar(20);not null;default:'';comment:昵称"`
	Age          uint8          `gorm:"size:3;comment:年龄"`
	Phone        string         `gorm:"type:char(11);uniqueIndex:un_phone;comment:手机号"`
	MemberNumber string         `gorm:"type:varchar(20);index:un_number;comment:会员编号"`
	Likes        string         `gorm:"type:varchar(200);comment:爱好"`
	Address      string         `gorm:"type:varchar(200);comment:家庭地址"`
	Birthday     sql.NullString `gorm:"type:varchar(10);comment:生日"`
}
