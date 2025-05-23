// При загрузке страницы показываем список песен
document.addEventListener('DOMContentLoaded', async () => {
  const container = document.getElementById('songs-container');
  try {
    const songs = await fetchSongs();
    container.innerHTML = songs.map(s => `<div>${s.title} — ${s.artist}</div>`).join('');
  } catch (error) {
    container.innerHTML = `<p style="color:red">${error.message}</p>`;
  }
});
