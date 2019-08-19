package cache

import (
	"fmt"
	"time"

	"gitlab.com/SausageShoot/admin-server/utils/parse"
)

const defaultExpire = time.Hour

func tokenKey(token string) string {
	return fmt.Sprintf("TOKEN_%s", token)
}

// GetPlayerIDByToken : 通过token获取playerID
func (m *memCache) GetPlayerIDByToken(token string) (int, error) {
	playerID, err := m.cc.Get(tokenKey(token))
	if err != nil {
		return 0, err
	}
	return int(parse.Int(playerID)), nil
}

// UpdateToken ： 更新player的token信息
// 如果没有token信息，则直接加入，如果有token信息，则进行更新
func (m *memCache) SetToken(playerID int, token string) error {
	_, err := m.cc.Set(tokenKey(token), playerID, defaultExpire)
	return err
}

// UpdateExpire ： 修改Expire
func (m *memCache) UpdateExpire(token string) error {
	_, err := m.cc.Expire(tokenKey(token), defaultExpire)
	return err
}
