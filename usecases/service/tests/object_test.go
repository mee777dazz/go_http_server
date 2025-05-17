package service

import (
	"fmt"
	"http_server/usecases/service"
	"testing"

	"http_server/usecases/service/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestObject_Get(t *testing.T) {
	repo := &mocks.Repository{
		Data: map[string]string{"key1": "value1"},
	}
	obj := service.NewObject(repo, nil)

	value, err := obj.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", *value)

	value, err = obj.Get("key2")
	assert.NoError(t, err)
	assert.Nil(t, value)
}

func TestObject_Put(t *testing.T) {
	repo := &mocks.Repository{
		Data: make(map[string]string),
	}

	obj := service.NewObject(repo, nil)
	err := obj.Put("key1", "value1")
	assert.NoError(t, err)

	value, err := repo.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", *value)
}

func TestObject_Post(t *testing.T) {
	repo := &mocks.Repository{
		Data: make(map[string]string),
	}
	sender := &mocks.Sender{}
	obj := service.NewObject(repo, sender)

	err := obj.Post("key1", "value1")
	assert.NoError(t, err)

	value, err := repo.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", *value)

	sender.Err = fmt.Errorf("send error")
	err = obj.Post("key2", "value2")
	assert.Error(t, err)
	assert.Equal(t, fmt.Errorf("sending object: %w", sender.Err), err)
}

func TestObject_Delete(t *testing.T) {
	repo := &mocks.Repository{
		Data: map[string]string{"key1": "value1"},
	}
	obj := service.NewObject(repo, nil)

	err := obj.Delete("key1")
	assert.NoError(t, err)

	value, err := repo.Get("key1")
	assert.NoError(t, err)
	assert.Nil(t, value)
}
