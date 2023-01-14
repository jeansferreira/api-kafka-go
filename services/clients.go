package services

import (
	"collect-server/env"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientAuth struct {
	Username string
	Password string
}

type ClientsService struct {
	auth      ClientAuth
	ttl       time.Duration
	timestamp time.Time
	cache     map[string]Client
}

func NewClientsService(auth ClientAuth) ClientsService {
	return ClientsService{
		cache:     make(map[string]Client),
		auth:      auth,
		ttl:       5 * time.Minute,
		timestamp: time.Now(),
	}
}

type Client struct {
	Debug string            `json:"debug"`
	Info  map[string]string `json:"info"`
}

func (c *ClientsService) ValidateCache() bool {
	if time.Now().Sub(c.timestamp) > c.ttl {
		c.cache = make(map[string]Client)
		c.timestamp = time.Now()

		return true
	}

	return false
}

func (c *ClientsService) GetClient(apiKey string) (Client, error) {
	client, err := c.getClientFromCache(apiKey)

	if err == nil {
		return client, nil
	}

	client, err = c.getClientFromAPI(apiKey)

	if err == nil {
		c.cache[apiKey] = client
	}

	return client, err
}

func (c *ClientsService) getClientFromCache(apiKey string) (Client, error) {
	if c.ValidateCache() {
		return Client{}, errors.New("Cache expired")
	}

	if client, ok := c.cache[apiKey]; ok {
		return client, nil
	}

	return Client{}, errors.New("Not found in cache")
}

func (c *ClientsService) getClientFromAPI(apiKey string) (Client, error) {
	client := Client{}

	res, err := http.Get("https://" + c.auth.Username + ":" + c.auth.Password + "@" + env.CLIENTS_URL + "/clients/" + apiKey)

	if err != nil {
		return client, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return client, err
	}

	err = json.Unmarshal(body, &client)

	if err != nil {
		return client, err
	}

	return client, nil
}

func (c *ClientsService) SetAuth(username, password string) {
	c.auth = ClientAuth{Username: username, Password: password}
}
