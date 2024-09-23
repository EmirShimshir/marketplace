package redis

import (
	"encoding/json"
	"github.com/EmirShimshir/marketplace/internal/adapter/authProvider/port"
	"github.com/go-redis/redis/v7"
	"time"
)

type SessionStorage struct {
	redisClient *redis.Client
}

func NewSessionStorage(redisClient *redis.Client) *SessionStorage {
	return &SessionStorage{redisClient: redisClient}
}

func (s *SessionStorage) Get(refreshToken string) (port.AuthSession, error) {
	sessionJson, err := s.redisClient.Get(refreshToken).Bytes()
	if err != nil {
		return port.AuthSession{}, err
	}

	var session port.AuthSession
	err = json.Unmarshal(sessionJson, &session)
	if err != nil {
		return port.AuthSession{}, err
	}

	return session, nil
}

func (s *SessionStorage) Put(refreshToken string, session port.AuthSession,
	expireTime time.Duration) error {
	sessionJson, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return s.redisClient.Set(refreshToken, sessionJson, expireTime).Err()
}

func (s *SessionStorage) Delete(refreshToken string) error {
	return s.redisClient.Del(refreshToken).Err()
}
