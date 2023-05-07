package main

import "encoding/base64"

type Configuration struct {
	Users      []User `json:"users"`      // 用户信息
	ShowClient bool   `json:"showClient"` // 是否显示浏览器窗口
}

type User struct {
	ID       string `json:"user"`     // 登录手机号
	Password string `json:"password"` // 登录密码
	Encoded  bool   `json:"encoded"`  // 密码是否已经加密
	Ignore   bool   `json:"ignore"`   // 是否忽略该账号
	Steps    Steps  `json:"steps"`    // 步数范围
}

type Steps struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type ChangeCommand struct {
	UserID     string
	Password   string
	StepNumber int
}

func (user User) getPassword() string {
	if user.Encoded {
		// 使用base64解码
		password, err := base64.StdEncoding.DecodeString(user.Password)
		if err != nil {
			return ""
		}
		return string(password)
	}
	return user.Password
}
