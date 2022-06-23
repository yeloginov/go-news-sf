/* 
    Схема БД

    Проект GoNews
    Автор: Егор Логинов (GO-11) по заданию SkillFactory начиная с модуля 30.8
*/

DROP TABLE IF EXISTS posts, authors;
GRANT usage ON SCHEMA public TO public;

-- Авторы
CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Публикации
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now())
);

INSERT INTO authors (name) VALUES ('Pavel');
INSERT INTO authors (name) VALUES ('Mariya');
--INSERT INTO posts (author_id, title, content, created_at) VALUES (0, 'Статья', 'Содержание статьи', 0);