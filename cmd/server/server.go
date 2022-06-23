// Пакет main
// Проект GoNews
// Автор: Егор Логинов (GO-11) по заданию SkillFactory в модуле 31.3

package main

import (
	"fmt"
	"gonews/pkg/api"
	"gonews/pkg/storage"
	"gonews/pkg/storage/postgres"
	"log"
	"net/http"
)

// Параметры подключения к БД Postgres.
const (
	DBHost     = "89.223.121.125"
	DBPort     = "5432"
	DBName     = "gonews"
	DBUser     = "gn_external"
	DBPassword = "Tdf_p9EXa9n"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Объект сервера
	var srv server

	// Объекты баз данных
	//
	// БД в памяти.
	//db := memdb.New()

	// Реляционная БД PostgreSQL.
	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName))
	if err != nil {
		log.Fatal(err)
	}
	// Инициация – создание таблиц
	err = db.CreateTables("././schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	/*	// Документная БД MongoDB.
		db3, err := mongo.New("mongodb://server.domain:27017/")
		if err != nil {
			log.Fatal(err)
		}
		_, _ = db2, db3
	*/

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}

// Тестирование:
//
// Добавление нового документа:
// curl -X POST -H "Content-Type: application/json" -d '{"ID":3,"Title":"New post added for Test","AuthorID":1,"Content":"Interesting.. The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine."}' http://localhost:8080/posts
// curl -X PUT -H "Content-Type: application/json" -d '{"ID":1,"Title":"Updated post - Author changed","AuthorID":2,"Content":"Interesting.. The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine."}' http://localhost:8080/posts
// curl -X DELETE -H "Content-Type: application/json" -d '{"ID":1,"Title":"Updated post - Author changed","AuthorID":2,"Content":"Interesting.. The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine."}' http://localhost:8080/posts
