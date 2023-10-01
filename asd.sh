#!/bin/bash
echo "подтягиваем репы"
docker pull druiddb/best_friends_bot

echo "удаляем старые сервисы"
docker compose down

echo "запускаем новые"
docker compose up

echo "завершили обновление"