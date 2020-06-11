package model

import (
	"time"
)

type User struct {
	Id            int64       `xorm:"pk autoincr"`
	Name          string	  `xorm:"varchar(50) notnull"`
	Username      string      `xorm:"varchar(50) notnull unique"`
	Password      string      `xorm:"varchar(255) notnull"`
	CreatedAt     time.Time   `xorm:"created"` // 这个Field将在Insert时自动赋值为当前时间
	UpdatedAt     time.Time   `xorm:"updated"` // 这个Field将在Insert或Update时自动赋值为当前时间
	DeletedAt     time.Time   `xorm:"deleted"` // 如果带DeletedAt这个字段和标签，xorm删除时自动软删除
}

func (this *User) ResponseUser() interface{} {
	return map[string]interface{} {
		"id": this.Id,
		"name": this.Name,
		"username": this.Username,
		"created_at": this.CreatedAt,
	}
}
