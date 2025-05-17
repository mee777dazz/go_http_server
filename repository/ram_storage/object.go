package ram_storage

import (
	"errors"
	"http_server/repository"
)

type Object struct {
	data map[string]string
}

func NewObject() *Object {
	return &Object{
		data: make(map[string]string),
	}
}

func (rs *Object) Get(key string) (*string, error) {
	value, exists := rs.data[key]
	if !exists {
		return nil, repository.NotFound
	}
	return &value, nil
}

func (rs *Object) Put(key string, value string) error {
	rs.data[key] = value
	return nil
}

func (rs *Object) Post(key string, value string) error {
	if _, exists := rs.data[key]; exists {
		return errors.New("key already exists")
	}
	rs.data[key] = value
	return nil
}

func (rs *Object) Delete(key string) error {
	if _, exists := rs.data[key]; !exists {
		return errors.New("key not found")
	}
	delete(rs.data, key)
	return nil
}
