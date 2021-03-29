package global

import (
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/settings"
)

// 包全局变量
var (
	ServerSetting   *settings.ServerSettingS
	AppSetting      *settings.AppSettingS
	DatabaseSetting *settings.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *settings.JWTSettingS
	EmailSetting    *settings.EmailSettingS
)
