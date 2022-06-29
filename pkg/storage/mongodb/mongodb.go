// Пакет mongodb
//
// Проект GoNews
// Автор: Егор Логинов (GO-11) по заданию SkillFactory в модуле 31.3
//
// Реализует CRUD-операции с хранилищем на основе MongoDB

package mongodb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gonews/pkg/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client   *mongo.Client
	dbname   string
	collname string
}

// Конструктор клиента БД MongoDB с параметрами подключения по адресу uri
// с имменем БД dn и коллекцией cn
func New(uri, dn, cn string) (*Store, error) {

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
		client:   c,
		dbname:   dn,
		collname: cn,
	}

	return &s, nil
}

// Метод закрывает подключение клиента к БД mongo
func (s *Store) Close() {
	s.client.Disconnect(context.Background())
}

// Реализуем интерфейс Interface из Storage: метод Posts()
func (s *Store) Posts() ([]storage.Post, error) {

	// Получаем указатель на коллекцию документов в БД Mongo
	coll := s.client.Database(s.dbname).Collection(s.collname)
	// Объект фильтра на документы ("все документы")
	f := bson.D{}
	// Получаем выборку (курсор) документов из БД
	c, err := coll.Find(context.Background(), f)
	if err != nil {
		return nil, err
	}
	// Закрываем курсор
	defer c.Close(context.Background())

	// Парсим документы из БД в структуру Post
	var pset []storage.Post
	for c.Next(context.Background()) {
		var p storage.Post
		err = c.Decode(&p)
		if err != nil {
			return nil, err
		}
		pset = append(pset, p)
	}

	return pset, nil
}

// Реализуем интерфейс Interface из Storage: AddPost()
func (s *Store) AddPost(p storage.Post) error {

	// Получаем указатель на коллекцию документов в БД Mongo
	coll := s.client.Database(s.dbname).Collection(s.collname)
	// Преобразуем структуру в документ json
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// Поиск такого документа по ID
	f := bson.D{{"ID", p.ID}}
	c, err := coll.Find(context.Background(), f)
	if err != nil {
		return err
	}
	// Если документ уже присуствует, возвращаем ошибку
	if c.TryNext(context.Background()) {
		return errors.New(fmt.Sprint("Document %s is already exists", p.ID))
	}

	// Добавляем новый документ в коллекцию БД
	_, err = coll.InsertOne(context.Background(), d)
	if err != nil {
		return err
	}

	return nil
}

// Реализуем интерфейс Interface из Storage: UpdatePost()
func (s *Store) UpdatePost(p storage.Post) error {

	// Получаем указатель на коллекцию документов в БД Mongo
	coll := s.client.Database(s.dbname).Collection(s.collname)
	// Поиск нужного документа по ID
	f := bson.D{{"ID", p.ID}}
	u, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// Обновляем документ
	_, err = coll.UpdateOne(context.Background(), f, u)
	if err != nil {
		return err
	}
	return nil
}

// Реализуем интерфейс Interface из Storage: DeletePost()
func (s *Store) DeletePost(p storage.Post) error {

	// Получаем указатель на коллекцию документов в БД Mongo
	coll := s.client.Database(s.dbname).Collection(s.collname)
	// Поиск нужного документа по ID
	f := bson.D{{"ID", p.ID}}
	_, err := coll.DeleteOne(context.Background(), f)
	if err != nil {
		return err
	}

	return nil
}
