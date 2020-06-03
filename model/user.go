package model

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            int64      `xorm:"pk autoincr" json:"id"`
	Name          string
	Username      string
	Password      string
	CreatedAt time.Time `xorm:"created"` // 这个Field将在Insert时自动赋值为当前时间
	UpdatedAt time.Time `xorm:"updated"` // 这个Field将在Insert或Update时自动赋值为当前时间
	DeletedAt time.Time `xorm:"deleted"` // 如果带DeletedAt这个字段和标签，xorm删除时自动软删除
}

func (u User) IsValid() bool {
	return u.Id > 0
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashed []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false
	}

	return true
}