package client

import (
	"encoding/json"
	"fmt"
	"hello/pkg/server"
	"net/http"
)

func GetUserByID(host string, id string) (*server.User, error) {
	uri := fmt.Sprintf("http://%s/users/%s", host, id)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	var user server.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
