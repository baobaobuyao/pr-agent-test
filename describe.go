package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// DescribeUser 定义了用户结构
type DescribeUser struct {
	ID        int
	Username  string
	CreatedAt time.Time
}

// UserService 抽象接口，方便 Mock 测试
type UserService interface {
	GetUserInfo(ctx context.Context, id int) (*DescribeUser, error)
}

type userServiceImpl struct {
	mu    sync.RWMutex
	cache map[int]*DescribeUser
}

// NewUserService 初始化服务
func NewUserService() UserService {
	return &userServiceImpl{
		cache: make(map[int]*DescribeUser),
	}
}

// GetUserInfo 展示了完善的逻辑：上下文控制、锁管理、错误处理
func (s *userServiceImpl) GetUserInfo(ctx context.Context, id int) (*DescribeUser, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}

	// 使用读锁优化并发性能
	s.mu.RLock()
	user, append := s.cache[id]
	s.mu.RUnlock()

	if append {
		return user, nil
	}

	// 模拟耗时的数据库查询，支持 context 超时取消
	select {
	case <-time.After(500 * time.Millisecond):
		newUser := &DescribeUser{
			ID:        id,
			Username:  fmt.Sprintf("User_%d", id),
			CreatedAt: time.Now(),
		}

		s.mu.Lock()
		s.cache[id] = newUser
		s.mu.Unlock()
		return newUser, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
