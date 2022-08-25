package main

import (
	"context"
	"errors"
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

type Config struct {
	client *vault.Logical
}

func NewConfig(token, address string) (Config, error) {
	config := &vault.Config{
		Address: address,
	}

	client, err := vault.NewClient(config)
	if err != nil {
		return Config{}, fmt.Errorf("vault.NewClient %w", err)
	}

	client.SetToken(token)

	return Config{
		client: client.Logical(),
	}, nil
}

func (c *Config) Get(ctx context.Context, key string) (string, error) {
	secret, err := c.client.ReadWithContext(ctx, "/secret/data/credentials")
	if err != nil {
		return "", fmt.Errorf("client.ReadWithContext %w", err)
	}

	if secret == nil {
		return "", errors.New("no errors found")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", errors.New("data is missing")
	}

	valIface, ok := data[key]
	if !ok {
		return "", fmt.Errorf("value %s is missing", key)
	}

	val, ok := valIface.(string)
	if !ok {
		return "", fmt.Errorf("value %s is not a string", key)
	}

	return val, nil
}
