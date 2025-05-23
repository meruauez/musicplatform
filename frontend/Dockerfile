# Используем официальный nginx образ
FROM nginx:alpine

# Копируем все файлы фронтенда в папку nginx для раздачи
COPY . /usr/share/nginx/html

# nginx по умолчанию слушает 80 порт
EXPOSE 80
