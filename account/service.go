package account

import (
	"encoding/json"
	"errors"
	"net/http"
	"supreme-go/redis"
	"time"

	goredis "github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Service struct {
	redis redis.Storage
}

func (s *Service) HandleList(c *fiber.Ctx) error {
	return c.SendString("read list")
}

func (s *Service) HandleRead(c *fiber.Ctx) error {
	// Read params
	id, exists := c.AllParams()["id"]
	if !exists {
		c.SendStatus(http.StatusBadRequest)
		return errors.New(`{ "error": "cannot get id from route" }`)
	}

	// Validate Params
	if _, err := uuid.Parse(id); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return errors.New(`{ "error": "cannot parse id as UUID" }`)
	}

	// Read from redis
	response := Response{}
	switch err := s.redis.Get(id, &response); err {
	case goredis.Nil:
		c.SendStatus(http.StatusNotFound)
		return c.SendString(`{ "error": "account not found" }`)
	case nil:
		return c.JSON(response)
	default:
		c.SendStatus(http.StatusServiceUnavailable)
		return c.SendString(`{ "error": "` + err.Error() + `"}`)
	}
}

func (s *Service) HandleCreate(c *fiber.Ctx) error {
	// Read body
	request := Request{}
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return c.SendString(`{ "error": "cannot read json body" }`)
	}

	// Validate body
	if valid, err := request.Valid(); !valid {
		c.SendStatus(http.StatusBadRequest)
		return c.SendString(`{ "error": "` + err.Error() + `"}`)
	}

	// Write to redis
	request.Id = uuid.New().String()
	if err := s.redis.Set(request.Id, request, 0); err != nil {
		c.SendStatus(http.StatusServiceUnavailable)
		return c.SendString(`{ "error": "cannot write to redis" }`)
	}

	return c.SendString("account created")
}

func (s *Service) HandleUpdate(c *fiber.Ctx) error {
	// Read params
	id, exists := c.AllParams()["id"]
	if !exists {
		c.SendStatus(http.StatusBadRequest)
		return errors.New(`{ "error": cannot get id from route" }`)
	}

	// Validate params
	if _, err := uuid.Parse(id); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return errors.New(`{ "error": "cannot parse id as UUID" }`)
	}

	// Read body
	request := Request{}
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return c.SendString(`{ "error": "cannot read json body" }`)
	}

	// Validate body
	if valid, err := request.Valid(); !valid {
		c.SendStatus(http.StatusBadRequest)
		return c.SendString(`{ "error": "` + err.Error() + `"}`)
	}

	// get by id
	response := Response{}
	switch err := s.redis.Get(id, &response); err {
	case goredis.Nil:
		c.SendStatus(http.StatusNotFound)
		return c.SendString(`{ "error": "account not found" }`)
	case nil:

	default:
		c.SendStatus(http.StatusServiceUnavailable)
		return c.SendString(`{ "error": "` + err.Error() + `"}`)
	}

	// Write to redis
	request.Id = id
	if err := s.redis.Set(request.Id, request, 0); err != nil {
		c.SendStatus(http.StatusServiceUnavailable)
		return c.SendString(`{ "error": "cannot write to redis" }`)
	}

	return c.SendString("update account")
}

func (s *Service) HandleRemove(c *fiber.Ctx) error {
	// Read params
	id, exists := c.AllParams()["id"]
	if !exists {
		c.SendStatus(http.StatusBadRequest)
		return errors.New(`{ "error": cannot get id from route" }`)
	}

	// Validate params
	if _, err := uuid.Parse(id); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return errors.New(`{ "error": "cannot parse id as UUID" }`)
	}

	// Write to redis
	if err := s.redis.Set(id, Request{}, time.Millisecond*1); err != nil {
		c.SendStatus(http.StatusServiceUnavailable)
		return c.SendString(`{ "error": "cannot write to redis" }`)
	}

	return c.SendString("remove account")
}

func NewService(redis redis.Storage) Service {
	return Service{
		redis: redis,
	}
}
