package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Marshaler interface {
	MarshalString() (string, error)
}

type Unmarshaler interface {
	UnmarshalString(string) error
}

type Storage struct {
	client *redis.Client
}

func (s *Storage) List() error {
	iterator := s.client.Scan(0, "prefix:*", 0).Iterator()
	for iterator.Next() {
		fmt.Println("keys", iterator.Val())
	}
	if err := iterator.Err(); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Get(key string, unmarshaler Unmarshaler) error {
	document, err := s.client.Get(key).Result()
	if err != nil {
		return err
	}
	return unmarshaler.UnmarshalString(document)
}

func (s *Storage) Set(key string, marshaler Marshaler, duration time.Duration) error {
	content, err := marshaler.MarshalString()
	if err != nil {
		return err
	}
	return s.client.Set(key, content, duration).Err()
}

func NewStorage(client *redis.Client) Storage {
	return Storage{
		client: client,
	}
}
