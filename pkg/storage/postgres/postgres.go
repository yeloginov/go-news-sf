// Пакет postgres
//
// Проект GoNews
// Автор: Егор Логинов (GO-11) по заданию SkillFactory в модуле 31.3
//
// Реализует CRUD-операции с хранилищем на основе Postgres

package postgres

import (
	"context"
	"gonews/pkg/storage"
	"io/ioutil"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных
type Store struct {
	pool *pgxpool.Pool
}

// Конструктор объекта БД
func New(p string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), p)
	if err != nil {
		return nil, err
	}

	s := Store{
		pool: db,
	}
	return &s, nil
}

// Создание таблиц БД на основе файла схемы schm
func (s *Store) CreateTables(schm string) error {

	// Читаем файл SQL-запроса со схемой БД
	buf, err := ioutil.ReadFile(schm)
	if err != nil {
		return err
	}

	// Выполняем SQL-запрос создания структуры БД
	sql := string(buf)
	_, err = s.pool.Exec(context.Background(), sql)
	if err != nil {
		return err
	}

	return nil
}

// Реализуем интерфейс Interface из Storage: метод Posts()
func (s *Store) Posts() ([]storage.Post, error) {
	sql := `
		SELECT posts.id, posts.author_id, posts.title, posts.content, posts.created_at, authors.name
		FROM posts INNER JOIN authors
		ON posts.author_id = authors.id
		ORDER BY id;
	`
	rows, err := s.pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	var pset []storage.Post
	// Сканирование строк результата запроса в структуру
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreatedAt, &p.AuthorName)
		if err != nil {
			return nil, err
		}
		pset = append(pset, p)
	}

	// TODO: из комментария Автора модуля - разобраться с rows.Err() (как может проявитсья)
	return pset, rows.Err()
}

// Реализуем интерфейс Interface из Storage: AddPost()
func (s *Store) AddPost(p storage.Post) error {

	sql := `INSERT INTO posts (title, content, author_id) 
			VALUES ($1, $2, $3)`
	_, err := s.pool.Exec(context.Background(), sql, p.Title, p.Content, p.AuthorID)

	return err
}

// Реализуем интерфейс Interface из Storage: UpdatePost()
func (s *Store) UpdatePost(p storage.Post) error {

	sql := `UPDATE posts
	SET title = $2, content = $3, author_id = $4
	WHERE id = $1;`
	_, err := s.pool.Exec(context.Background(), sql, p.ID, p.Title, p.Content, p.AuthorID)

	return err
}

// Реализуем интерфейс Interface из Storage: DeletePost()
func (s *Store) DeletePost(p storage.Post) error {

	sql := `DELETE FROM posts WHERE id = $1;`
	_, err := s.pool.Exec(context.Background(), sql, p.ID)

	return err
}
