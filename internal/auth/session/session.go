package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v7"
)

var ErrSessionNotFound = errors.New("session not found")

type SessionClaims struct {
	UserID     int       `json:"user_id,omitempty"`
	ExternalID string    `json:"external_id"`
	CreatedAt  time.Time `json:"created_at"`
	IDToken    string    `json:"id_token"`
}

// Sessions store config
type Config struct {
	Addr       string
	Password   string
	DB         int
	SessionTTL int `default:30` // session max age in minutes
}

type Session struct {
	sessionTTL time.Duration
	client     *redis.Client
}

// New
func New(cfg *Config) *Session {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Session{
		client:     client,
		sessionTTL: time.Minute * time.Duration(cfg.SessionTTL),
	}
}

var Unauthorized = &SessionClaims{UserID: 0, CreatedAt: time.Unix(0, 0)}

func (s *Session) Get(key string) (*SessionClaims, error) {
	bytes, err := s.client.Get(s.skey(key)).Bytes()
	if err == redis.Nil {
		return nil, ErrSessionNotFound
	}
	if err != nil {
		return nil, err
	}

	claims := &SessionClaims{}
	err = json.Unmarshal(bytes, claims)
	if err != nil {
		return nil, err
	}

	resetTime, err := s.resetTime(claims.ExternalID)
	if err != nil {
		return nil, err
	}

	if claims.CreatedAt.Before(resetTime) {
		// ignore err
		s.Drop(key)
		return nil, ErrSessionNotFound
	}

	return claims, nil
}

func (s *Session) Set(key string, claims *SessionClaims) error {
	bytes, err := json.Marshal(claims)
	if err != nil {
		return nil
	}

	return s.client.Set(s.skey(key), bytes, s.sessionTTL).Err()
}

func (s *Session) Drop(key string) error {
	return s.client.Del(s.skey(key)).Err()
}

func (s *Session) DropAllByExternalID(ctx context.Context, externalID string) error {
	return s.client.Set(s.rkey(externalID), time.Now().Unix(), s.sessionTTL).Err()
}

func (s *Session) resetTime(key string) (time.Time, error) {
	timestamp, err := s.client.Get(s.rkey(key)).Int64()
	if err == redis.Nil {
		return time.Unix(0, 0), nil
	}
	if err != nil {
		return time.Unix(0, 0), err
	}
	return time.Unix(timestamp, 0), nil
}

// skey returns session:key
func (s *Session) skey(key string) string {
	return "session:" + key
}

// rkey returns reset:key
func (s *Session) rkey(key string) string {
	return "reset:" + key
}
