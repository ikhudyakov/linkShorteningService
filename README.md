# linkShorteningService
 
Сервис для сокращения ссылок
Golang, PostgreSQL

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
