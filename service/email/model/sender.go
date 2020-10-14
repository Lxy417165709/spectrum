package model

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

const (
	RegisterEmailSubject       = "注册邮件"
	ChangePasswordEmailSubject = "修改密码邮件"
	RegisterKeyPrefix          = "register"
	ChangePasswordKeyPrefix    = "changePassword"
	UnExpectContentType        = ""
)

type TypeOfEmailContext string

const (
	HtmlType  TypeOfEmailContext = "html"
	PlainType                    = "plain"
)

type EmailContext struct {
	EmailAddrOfSender    string
	EmailAddrOfReceivers []string
	Subject              string
	Body                 string
	Type                 TypeOfEmailContext
}

func (c *EmailContext) GetContentTypeString() string {
	var contentType string
	switch c.Type {
	case HtmlType:
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", c.Type)
	case PlainType:
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", c.Type)
	default:
		logs.Warn("Invalid content type", contentType)
		return UnExpectContentType
	}
	return contentType
}
