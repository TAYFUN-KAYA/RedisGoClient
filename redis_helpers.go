package main

import (
	"errors"
	"fmt"
	"net"
	"redis_client_example/commands"
)

func (c *RedisConfig) Connect() (*RedisConfig, error) {
	stream, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Address, c.Port))
	if err != nil {
		return nil, err
	}

	c.Connection = RedisConnection{
		Stream: stream,
	}

	if c.Password != "" {

	}

	return c, nil
}

func (c *RedisConfig) Auth() (*RedisConfig, error) {
	buffer := make([]byte, 0, 4096)
	tmp := make([]byte, 256)
	authCommand := fmt.Sprintf("*%d\r\n$4\r\nAUTH\r\n$%d\r\n%s\r\n", 2, len(c.Password), c.Password)

	command := []byte(authCommand)

	_, err := c.Connection.Stream.Write(command)

	if err != nil {
		return nil, err
	}
	_, err = c.Connection.Stream.Read(tmp)
	if err != nil {
		return nil, err
	}

	for _, buffertmp := range tmp {
		buffer = append(buffer, buffertmp)
	}
	result := string(buffer)
	if string(result[0]) == "-" {
		return nil, errors.New("Authentication failed")
	}
	return c, nil
}

func (c *RedisConfig) Info() (*RedisResponse, error) {
	buffer := make([]byte, 0, 4096)
	tmp := make([]byte, 256)
	infoCommand := fmt.Sprintf("*%d\r\n%d\r\n%s\r\n", 1, len(commands.INFO), commands.INFO)

	command := []byte(infoCommand)

	_, err := c.Connection.Stream.Write(command)

	if err != nil {
		return nil, err
	}
	_, err = c.Connection.Stream.Read(tmp)
	if err != nil {
		return nil, err
	}

	for _, buffertmp := range tmp {
		buffer = append(buffer, buffertmp)
	}
	result := string(buffer)
	if string(result[0]) == "-" {
		return &RedisResponse{
			Message: result,
		}, errors.New("Info Command Field")
	}
	return &RedisResponse{
		Message: result,
		Success: true,
	}, nil
}
