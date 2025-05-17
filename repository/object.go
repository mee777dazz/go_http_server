package repository

type Object interface {
	Get(key string) (*string, error)
	Put(key string, value string) error
	Post(key string, value string) error
	Delete(key string) error
}
