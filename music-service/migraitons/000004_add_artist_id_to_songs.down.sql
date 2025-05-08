-- migrations/<timestamp>_add_artist_id_to_songs.down.sql

-- Удаляем внешний ключ
ALTER TABLE songs
    DROP CONSTRAINT fk_artist;

-- Удаляем столбец artist_id
ALTER TABLE songs
    DROP COLUMN artist_id;
