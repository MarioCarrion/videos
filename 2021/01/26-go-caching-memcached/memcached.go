package main

import (
	"bytes"
	"encoding/gob"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Client struct {
	client *memcache.Client
}

func NewMemcached() (*Client, error) {
	// XXX Assuming environment variable contains only one server
	client := memcache.New(os.Getenv("MEMCACHED"))

	if err := client.Ping(); err != nil {
		return nil, err
	}

	client.Timeout = 100 * time.Millisecond
	client.MaxIdleConns = 100

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetName(nconst string) (Name, error) {
	item, err := c.client.Get(nconst)
	if err != nil {
		return Name{}, err
	}

	b := bytes.NewReader(item.Value)

	var res Name

	if err := gob.NewDecoder(b).Decode(&res); err != nil {
		return Name{}, err
	}

	return res, nil
}

func (c *Client) SetName(n Name) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		return err
	}

	return c.client.Set(&memcache.Item{
		Key:        n.NConst,
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(25 * time.Second).Unix()),
	})
}
