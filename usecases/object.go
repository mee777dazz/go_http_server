package usecases

type Object interface {
	Get(key string) (*string, error)
	Put(key string, value string) error // Put(domain.Object) error
	Post(key string, value string) error
	Delete(key string) error
}
