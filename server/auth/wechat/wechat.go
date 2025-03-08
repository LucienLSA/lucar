package wechat

import (
	"fmt"

	"github.com/medivhzhan/weapp/v2"
)

// Service implements a wechat auth service
type Service struct {
	AppID     string
	AppSecret string
}

// Resolver resolves authorization code to wechat open id.
func (s *Service) Resolve(code string) (string, error) {
	resp, err := weapp.Login(s.AppID, s.AppSecret, code)
	if err != nil {
		return "", fmt.Errorf("weapp.Login failed, err:%v\n", err)
	}
	err = resp.GetResponseError()
	if err != nil {
		return "", fmt.Errorf("weapp response failed, err:%v\n", err)
	}
	return resp.OpenID, nil
}
