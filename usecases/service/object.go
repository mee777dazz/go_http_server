package service

import (
	"fmt"
	"http_server/domain"
	"http_server/repository"
)

type Object struct {
	repo repository.Object
	sender repository.ObjectSender
}

func NewObject(
	repo repository.Object,
	sender repository.ObjectSender,
) *Object {
	return &Object{
		repo: repo,
		sender: sender,
	}
}

func (rs *Object) Get(key string) (*string, error) {
	return rs.repo.Get(key)
}

func (rs *Object) Put(key string, value string) error {
	return rs.repo.Put(key, value)
}

func (rs *Object) Post(key string, value string) error {
	err := rs.sender.Send(domain.Object{Key: key, Value: value})
	if err != nil {
		return fmt.Errorf("sending object: %w", err)
	}
	return rs.repo.Post(key, value)
}

func (rs *Object) Delete(key string) error {
	return rs.repo.Delete(key)
}
