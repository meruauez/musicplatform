const token = localStorage.getItem('token');
if (!token) {
  // Если токена нет — перенаправляем на страницу входа
  window.location.href = 'login.html';
}

let editingSongId = null;

function showLoading(text = 'Загрузка...') {
  const loadingEl = document.getElementById('loading');
  loadingEl.textContent = text;
  loadingEl.style.display = 'block';
}

function hideLoading() {
  document.getElementById('loading').style.display = 'none';
}

function showError(text) {
  const errorEl = document.getElementById('error');
  errorEl.textContent = text;
  errorEl.style.display = 'block';
}

function hideError() {
  const errorEl = document.getElementById('error');
  errorEl.style.display = 'none';
}

function clearMessage() {
  document.getElementById('message').textContent = '';
}

async function fetchArtists() {
  showLoading('Загрузка артистов...');
  hideError();
  try {
    const response = await fetch('http://localhost:8082/artists', {
      headers: { Authorization: 'Bearer ' + token }
    });
    if (!response.ok) throw new Error('Ошибка загрузки артистов');

    const artists = await response.json();
    const artistSelect = document.getElementById('artist-select');
    artistSelect.innerHTML = '<option value="">Выберите артиста</option>';
    artists.forEach(artist => {
      const option = document.createElement('option');
      option.value = artist.id;
      option.textContent = artist.name;
      artistSelect.appendChild(option);
    });
  } catch (e) {
    showError(e.message);
  } finally {
    hideLoading();
  }
}

async function fetchGenres() {
  showLoading('Загрузка жанров...');
  hideError();
  try {
    const response = await fetch('http://localhost:8082/genres', {
      headers: { Authorization: 'Bearer ' + token }
    });
    if (!response.ok) throw new Error('Ошибка загрузки жанров');

    const genres = await response.json();
    const genreSelect = document.getElementById('genre-select');
    genreSelect.innerHTML = '<option value="">Выберите жанр</option>';
    genres.forEach(genre => {
      const option = document.createElement('option');
      option.value = genre.id;
      option.textContent = genre.name;
      genreSelect.appendChild(option);
    });
  } catch (e) {
    showError(e.message);
  } finally {
    hideLoading();
  }
}

async function fetchSongs() {
  showLoading('Загрузка песен...');
  hideError();
  clearMessage();
  try {
    const response = await fetch('http://localhost:8082/songs', {
      headers: { Authorization: 'Bearer ' + token }
    });
    if (!response.ok) throw new Error('Ошибка загрузки песен');

    const songs = await response.json();
    renderSongs(songs);
  } catch (e) {
    showError(e.message);
  } finally {
    hideLoading();
  }
}

function renderSongs(songs) {
  const container = document.getElementById('songs-container');
  container.innerHTML = '';

  songs.forEach(song => {
    const div = document.createElement('div');
    div.innerHTML = `
      <b>${song.name}</b> (Артист: ${song.artistName}, Жанр: ${song.genreName}) 
      <button onclick="startEditSong(${song.id})">Редактировать</button>
      <button onclick="deleteSong(${song.id})">Удалить</button>
    `;
    container.appendChild(div);
  });
}

async function startEditSong(songId) {
  editingSongId = songId;
  showLoading('Загрузка данных песни...');
  hideError();
  clearMessage();

  try {
    const response = await fetch(`http://localhost:8082/songs/${songId}`, {
      headers: { Authorization: 'Bearer ' + token }
    });
    if (!response.ok) throw new Error('Ошибка при загрузке песни');

    const song = await response.json();

    document.getElementById('song-name').value = song.name;
    document.getElementById('artist-select').value = song.artistId;
    document.getElementById('genre-select').value = song.genreId;
    document.getElementById('submit-button').textContent = 'Сохранить изменения';
  } catch (e) {
    showError(e.message);
  } finally {
    hideLoading();
  }
}

async function addOrEditSong(event) {
  event.preventDefault();
  hideError();
  clearMessage();

  const name = document.getElementById('song-name').value.trim();
  const artistId = document.getElementById('artist-select').value;
  const genreId = document.getElementById('genre-select').value;

  if (!name || !artistId || !genreId) {
    showError('Пожалуйста, заполните все поля формы.');
    return;
  }

  showLoading(editingSongId ? 'Сохранение изменений...' : 'Добавление песни...');

  try {
    let url = 'http://localhost:8082/songs';
    let method = 'POST';

    if (editingSongId) {
      url += `/${editingSongId}`;
      method = 'PUT';
    }

    const response = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + token,
      },
      body: JSON.stringify({ name, artistId, genreId }),
    });

    if (!response.ok) {
      const err = await response.text();
      throw new Error(err || 'Ошибка при сохранении песни');
    }

    document.getElementById('message').textContent = editingSongId ? 'Песня успешно обновлена!' : 'Песня успешно добавлена!';
    document.getElementById('add-song-form').reset();
    document.getElementById('submit-button').textContent = 'Добавить песню';
    editingSongId = null;

    await fetchSongs();
  } catch (e) {
    showError(e.message);
  } finally {
    hideLoading();
  }
}

async function deleteSong(songId) {
  if (!confirm('Вы уверены, что хотите удалить эту песню?')) return;

  showLoading('Удаление песни...');
  hideError();
  clearMessage();

  try {
    const response = await fetch(`http://localhost:8082/songs/${songId}`, {
      method: 'DELETE',
      headers: { Authorization: 'Bearer ' + token }
    });
    if (!response.ok) throw new Error('Ошибка при удалении песни');

    document.getElementById('message').textContent = 'Песня удалена.';
    await fetchSongs();
  } catch (e) {
    showError(e.message);
  } finally {
    hideLoading();
  }
}

function logout() {
  localStorage.removeItem('token');
  window.location.href = 'login.html';
}

async function init() {
  await Promise.all([fetchArtists(), fetchGenres()]);
  await fetchSongs();

  document.getElementById('add-song-form').addEventListener('submit', addOrEditSong);
  document.getElementById('logout-button').addEventListener('click', logout);
}

init();
