-- migrations/<timestamp>_add_artist_id_to_songs.up.sql

-- Добавляем столбец artist_id в таблицу songs
ALTER TABLE songs ADD COLUMN artist_id INTEGER;

-- Устанавливаем внешний ключ, связывающий artist_id с id из таблицы artists
ALTER TABLE songs
    ADD CONSTRAINT fk_artist
    FOREIGN KEY (artist_id)
    REFERENCES artists(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE;
