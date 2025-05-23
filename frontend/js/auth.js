// --- Универсальные функции ---
function showError(text) {
  const errorEl = document.getElementById('error');
  errorEl.textContent = text;
  errorEl.style.display = 'block';
}

function hideError() {
  const errorEl = document.getElementById('error');
  errorEl.style.display = 'none';
}

function showMessage(text) {
  const messageEl = document.getElementById('message');
  if (messageEl) {
    messageEl.textContent = text;
  }
}

// --- Регистрация ---
async function registerUser(event) {
  event.preventDefault();
  hideError();
  showMessage('');

  const username = document.getElementById('username').value.trim();
  const email = document.getElementById('email').value.trim();
  const password = document.getElementById('password').value;

  if (!username || !email || !password) {
    showError('Заполните все поля.');
    return;
  }

  try {
    const response = await fetch('http://localhost:8082/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, email, password }),
    });
    if (!response.ok) {
      const err = await response.text();
      throw new Error(err || 'Ошибка регистрации');
    }

    showMessage('Регистрация успешна! Теперь войдите.');
  } catch (e) {
    showError(e.message);
  }
}

// --- Вход ---
async function loginUser(event) {
  event.preventDefault();
  hideError();

  const username = document.getElementById('username').value.trim();
  const password = document.getElementById('password').value;

  if (!username || !password) {
    showError('Заполните все поля.');
    return;
  }

  try {
    const response = await fetch('http://localhost:8082/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      const err = await response.text();
      throw new Error(err || 'Ошибка входа');
    }

    const data = await response.json();
    // Сохраняем токен в localStorage
    localStorage.setItem('token', data.token);
    // Переходим на главную страницу
    window.location.href = 'index.html';
  } catch (e) {
    showError(e.message);
  }
}

// --- Инициализация страниц регистрации и входа ---
function initAuthPage() {
  const registerForm = document.getElementById('register-form');
  if (registerForm) {
    registerForm.addEventListener('submit', registerUser);
  }

  const loginForm = document.getElementById('login-form');
  if (loginForm) {
    loginForm.addEventListener('submit', loginUser);
  }
}

initAuthPage();
