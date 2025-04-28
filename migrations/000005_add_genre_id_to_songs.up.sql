-- migrations/<timestamp>_add_genre_id_to_songs.up.sql

-- Добавляем столбец genre_id в таблицу songs
ALTER TABLE songs ADD COLUMN genre_id INTEGER;

-- Устанавливаем внешний ключ для genre_id
ALTER TABLE songs
    ADD CONSTRAINT fk_genre
    FOREIGN KEY (genre_id)
    REFERENCES genres(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE;
