
REST сервис вопросов и ответов на Go с чистой архитектурой.

## Быстрый запуск

```bash
docker-compose up
```

Приложение будет доступно по адресу: `http://localhost:8080`

## Описание проекта

Сервис предоставляет API для управления вопросами и ответами с валидацией данных, логированием и полным тестовым покрытием.

### Основные возможности

- Создание, просмотр и удаление вопросов
- Добавление ответов к вопросам
- Валидация данных (длина текста 5-200 символов)
- Структурированное логирование
- Миграции базы данных через Goose
- Health checks для сервисов

## Архитектура

```
internal/
├── entity/      # Бизнес-сущности
├── repository/  # Работа с данными (GORM)
├── usecase/     # Бизнес-логика
├── controller/  # HTTP обработчики
└── pkg/         # Вспомогательные пакеты
```

## API Endpoints

### Вопросы

| Метод | Endpoint | Описание |
|-------|----------|-----------|
| GET | `/question` | Все вопросы |
| GET | `/question/{id}` | Вопрос по ID |
| POST | `/question` | Создать вопрос |
| DELETE | `/question/{id}` | Удалить вопрос |

### Ответы

| Метод | Endpoint | Описание |
|-------|----------|-----------|
| GET | `/answer/{id}` | Ответ по ID |
| POST | `/question/{id}/answer` | Создать ответ |
| DELETE | `/answer/{id}` | Удалить ответ |

## Примеры запросов

### Создание вопроса
```bash
curl -X POST http://localhost:8080/question \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "text": "Как работает этот сервис?"
  }'
```

### Создание ответа
```bash
curl -X POST http://localhost:8080/question/1/answer \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "b2c3d4e5-f6g7-8901-bcde-f23456789012",
    "text": "Это сервис вопросов и ответов с REST API"
  }'
```

### Получение всех вопросов
```bash
curl http://localhost:8080/question
```

## База данных

Используется PostgreSQL с автоматическими миграциями. Таблицы:
- `questions` - вопросы
- `answers` - ответы

## Тестирование

```bash
# Запуск всех тестов
go test ./...

# Unit-тесты
go test ./internal/usecase/...

# Интеграционные тесты
go test ./internal/controller/...
```

## Переменные окружения

| Переменная | По умолчанию | Описание |
|------------|--------------|-----------|
| DB_HOST | localhost | Хост БД |
| DB_PORT | 5432 | Порт БД |
| DB_USER | postgres | Пользователь БД |
| DB_PASSWORD | postgres | Пароль БД |
| DB_NAME | testovoe | Имя БД |

## Ручная установка (без Docker)

```bash
# Установка зависимостей
go mod download

# Запуск миграций
go run main.go migrate up

# Запуск приложения
go run cmd/main/main.go
```

## Технологии

- Go 1.25+
- PostgreSQL
- GORM
- Zap (логирование)
- Goose (миграции)
- Docker
```
