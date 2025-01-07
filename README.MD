#                               **DeNet REST Api**

### Для запуска приложения:
Сначала скачайте проект к себе в окружение:
```
git clone https://github.com/Njrctr/DeNet_test && cd DeNet_test
```
Запуск:
```
make build && make run
```
Альтернативный запуск без утилиты make:
```
docker-compose build denet-backend && docker-compose up denet-backend
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate.up
```

### В данном проекте был реализован REST Api сервис для работы с кошельками:
* Весь функционал можно протестировать в SWAGGER документации: http://localhost:8080/swagger/
* Реализован функционал Аутентификации на основе JWT токена
* Реализован функционал создания Заданий, ввода реферального кода, выполнения заданий с наградой рефералу