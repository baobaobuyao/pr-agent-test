package service

import "errors"

// User 定义用户结构
type User struct {
	ID   int
	Name string
}

// GetUserByID 模拟获取用户逻辑
func GetUserByID(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}
	// 模拟返回
	return &User{ID: id, Name: "TestUser"}, nil
}
