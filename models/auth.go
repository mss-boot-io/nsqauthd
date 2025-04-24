package models

import (
	"context"
	"gorm.io/gorm"
)

type Auth struct {
	Login       string `gorm:"column:login;primaryKey;type:varchar(255);not null"`
	IP          string `gorm:"column:ip;type:varchar(255)"`
	TlsRequired bool   `gorm:"column:tls_required"`
	Topic       string `gorm:"column:topic;type:varchar(255)"`
	Channel     string `gorm:"column:channel;type:varchar(255)"`
	Subscribe   string `gorm:"column:subscribe;type:varchar(255)"`
	Publish     string `gorm:"column:publish;type:varchar(255)"`
}

func (*Auth) TableName() string {
	return "nsq_auth"
}

func (e *Auth) ToAuthPermissions() ([]string, string, []string) {
	permissions := make([]string, 0)
	if e.Subscribe != "" {
		permissions = append(permissions, e.Subscribe)
	}
	if e.Publish != "" {
		permissions = append(permissions, e.Publish)
	}
	return []string{e.Channel}, e.Topic, permissions
}

func (e *Auth) GetByLogin(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Where("login = ?", e.Login).First(e).Error
}
