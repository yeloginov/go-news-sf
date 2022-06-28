// Пакет mongodb
//
// Проект GoNews
// Автор: Егор Логинов (GO-11) по заданию SkillFactory в модуле 31.3
//
// Реализует CRUD-операции с хранилищем на основе MongoDB

package mongodb

import (
	"context"
	"gonews/pkg/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// in progress...

type Store struct {
	client *mongo.Client
}

// Конструктор клиента БД MongoDB с параметрами подключения по адресу uri
func New(uri string) (*Store, error) {

	// Подключение к БД
	mOpts := options.Client().ApplyURI(uri)
	c, err := mongo.Connect(context.Background(), mOpts)
	if err != nil {
		return nil, err
	}

	err = c.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	s := Store{
		client: c,
	}

	return &s, nil
}

// Метод закрывает подключение клиента к БД mongo
func (s *Store) Close() {
	s.client.Disconnect(context.Background())
}

// Реализуем интерфейс Interface из Storage: метод Posts()
func (s *Store) Posts() ([]storage.Post, error) {
	var pset []storage.Post
	return pset, nil
}

// Реализуем интерфейс Interface из Storage: AddPost()
func (s *Store) AddPost(p storage.Post) error {
	return nil
}

// Реализуем интерфейс Interface из Storage: UpdatePost()
func (s *Store) UpdatePost(p storage.Post) error {
	return nil
}

// Реализуем интерфейс Interface из Storage: DeletePost()
func (s *Store) DeletePost(p storage.Post) error {
	return nil
}
