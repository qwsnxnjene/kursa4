#!/bin/bash

# Папки вашего проекта
FRONTEND_DIR="."
BACKEND_DIR="./backend"

# Проверка и установка зависимостей фронтенда
if [ ! -d "$FRONTEND_DIR/node_modules" ]; then
  echo "Устанавливаем зависимости фронтенда..."
  cd "$FRONTEND_DIR" && npm install
  cd - > /dev/null
else
  echo "Зависимости фронтенда уже установлены."
fi

# Проверка и установка зависимостей бэкенда (Go)
if [ ! -f "$BACKEND_DIR/go.sum" ]; then
  echo "Устанавливаем зависимости бэкенда (Go)..."
  cd "$BACKEND_DIR" && go mod tidy
  cd - > /dev/null
else
  echo "Зависимости бэкенда уже установлены."
fi

# Запуск фронтенда и бэкенда параллельно
echo "Запускаем фронтенд и бэкенд..."

cd "$BACKEND_DIR" && go run main.go &   # Замените main.go на вашу точку входа
BACKEND_PID=$!
cd "$FRONTEND_DIR" && npm run serve &       # Или npm run dev, если у вас другой скрипт
FRONTEND_PID=$!

# Ожидание завершения обоих процессов
wait $BACKEND_PID
wait $FRONTEND_PID