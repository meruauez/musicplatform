// Здесь прописываем URL сервисов из docker-compose
const USER_SERVICE_URL = 'http://user-service:8081';
const MUSIC_SERVICE_URL = 'http://music-service:8082';

// Пример функции для получения песен
async function fetchSongs() {
  const response = await fetch(`${MUSIC_SERVICE_URL}/songs`);
  if (!response.ok) {
    throw new Error('Ошибка при загрузке песен');
  }
  return await response.json();
}

// Пример функции для добавления новой песни
async function addSong(songData) {
  const response = await fetch(`${MUSIC_SERVICE_URL}/songs`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(songData),
  });
  if (!response.ok) {
    throw new Error('Ошибка при добавлении песни');
  }
  return await response.json();
}
