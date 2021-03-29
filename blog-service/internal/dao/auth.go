package dao

import "github.com/go-programming-tour-book/blog-service/internal/model"

// function GetAuth 获取认证信息
func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}


