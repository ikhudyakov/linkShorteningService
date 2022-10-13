# linkShorteningService
 
# Сервис для сокращения ссылок
_Golang, PostgreSQL_

# Установка

- Склонировать репозиторий _```git clone https://github.com/ikhudyakov/linkShorteningService.git```_
- Перейти в папку _linkShorteningService_
- Выполнить команду ```docker-compose up -d```

# Использование

После запуска сервис доступен по адресу ```localhost:8001```
Для получения короткой ссылки в теле POST запроса отправляем JSON следующего вида:

```sh
{
    "link" : "https://vk.com"
    "domain" : 1
}
```

В ответ получаем JSON:

```sh
{
    "link": "https://vk.com",
    "shortlink": "localhost:8001/WI7U0i9e99"
}
```
