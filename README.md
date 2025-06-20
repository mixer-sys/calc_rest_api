# calc_rest_api
Калькулятор с REST API

# Запуск в контейнере
Создай .env файл. См. env.example.

Создай config.yaml.

Выполни команды
```bash
docker build -t app .
docker run  -p 8080:8080 --env-file .env app
```
# Документация
```
http://localhost:8080/swagger/
```
![alt text](<Screenshot from 2025-06-19 20-20-02.png>)