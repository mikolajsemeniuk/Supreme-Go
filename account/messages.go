package account

import (
	"encoding/json"
	"errors"
	"regexp"
)

type Response struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *Response) UnmarshalString(data string) error {
	return json.Unmarshal([]byte(data), r)
}

type Request struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r Request) MarshalString() (string, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r Request) Valid() (bool, error) {
	if len(r.Username) < 4 || len(r.Username) > 50 {
		return false, errors.New("username has to be between 4 and 50 characters")
	}

	re := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !re.Match([]byte(r.Email)) {
		return false, errors.New("email is not valid")
	}

	return true, nil
}
