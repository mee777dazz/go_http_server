package mocks

import "http_server/domain"

type Repository struct{
	Data map[string]string
	Err  error
}

func (m *Repository) Get(key string) (*string, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	value, exists := m.Data[key]
	if !exists {
		return nil, nil
	}
	return &value, nil
}

func (m *Repository) Put(key string, value string) error {
	if m.Err != nil {
		return m.Err
	}
	m.Data[key] = value
	return nil
}

func (m *Repository) Post(key string, value string) error {
	if m.Err != nil {
		return m.Err
	}
	m.Data[key] = value
	return nil
}

func (m *Repository) Delete(key string) error {
	if m.Err != nil {
		return m.Err
	}
	delete(m.Data, key)
	return nil
}

type Sender struct {
	Err error
}

func (m *Sender) Send(obj domain.Object) error {
	return m.Err
}
