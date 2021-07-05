package config

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	yaml "gopkg.in/yaml.v3"
)

type GithubConfig struct {
	FullName string `yaml:"full_name"`
	User     string `yaml:"user"`
	Owner    string `yaml:"owner"`
	Repo     string `yaml:"repo"`
	Token    string `yaml:"token"`
}

func ReadConfig(r io.Reader) (*GithubConfig, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to read from file: %s", err.Error()))
	}

	var conf GithubConfig
	err = yaml.Unmarshal(buf.Bytes(), &conf)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to read config: %s", err.Error()))
	}
	return &conf, nil
}

func ParseConfig() context.Context {
	home, _ := os.UserHomeDir()
	f, err := os.OpenFile(path.Join(home, ".salmon", "config.yaml"), os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	conf, err := ReadConfig(f)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	configMap := map[string]string{
		"name":  conf.FullName,
		"user":  conf.User,
		"owner": conf.Owner,
		"repo":  conf.Repo,
		"token": conf.Token,
	}
	return contextWithMap(ctx, configMap)
}

func contextWithMap(ctx context.Context, data map[string]string) context.Context {
	for key, value := range data {
		ctx = context.WithValue(ctx, key, value)
	}
	return ctx
}
