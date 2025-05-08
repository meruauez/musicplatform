-- migrations/<timestamp>_add_genre_id_to_songs.down.sql

-- Удаляем внешний ключ
ALTER TABLE songs
    DROP CONSTRAINT fk_genre;

-- Удаляем столбец genre_id
ALTER TABLE songs
    DROP COLUMN genre_id;
