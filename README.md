# REST-API-SERVER
Это личный репозиторий в котором я учуть работать с API. Этот репозиторий не содержит никакой полезной информации
просто в нём я учусь работать с API в своё удовольствие

***rest-api-serv***

**Для создания этого репозитория я использовал несколько библеотек:**
1. github.com/julienschmidt/httprouter - удобный http роутер(мультиплексор)
2. github.com/sirupsen/logrus - удобный и красивый логер
3. github.com/ilyakaznacheev/cleanenv - минималистичный и простой конфигуратор
4. go get github.com/jackc/pgx/v5 - драйвер постгрес
___
# Логика сервера
Сервер может отвечать на заданные API запросы  

POST /message --> create users
- Этот запрос получает сообщение, записывает его в бд (создаёт там user) и может вернуть: 204, 4XX, hendler location: url(ссылка на созданного пользователя)

GET /statistics --> get statistics 
- Этот метод выводит статистику по обработанным сообщениям

___
# Логирование:
В качестве логов я использовал обёртку для логрус у чела с ютуба.

Её так же можно использовать и в других, последующих проектах.


Для её использования нужно лишь передать её в нужное нам место с помощью:

`logger := logging.GetLogger()`

и  использовать с помощью:

 `logger.Info("logs")`

вместо инфо можно написать любой другой уровень логирования поддерживаемый логрус.

Данная обёртка выведет лог в консоль и запишет в нужный нам файл, в данном случае в `all.log` 

# Конфигурация
Конфиг написан с помощью библеотеки cleanenv и с его помощью можно запускать сервер как через **tcp**
соединение так и с помощью **сокета**


# Postgres
***Таблица users***:
- поля: 
```SQL

ID           string `json:"id" bson:"_id,omitempty"` - primery key
Email        string `json:"email" bson:"email"` - VARCHAR (100) NOT NULL
Username     string `json:"username" bson:"username"` - VARCHAR (100) NOT NULL
PasswordHash string `json:"-" bson:"password" - VARCHAR (100) NOT NULL

```

CREATE TABLE messages (
id SERIAL PRIMARY KEY,
content TEXT NOT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP NOT NULL
);

CREATE TABLE processed_messages (
id SERIAL PRIMARY KEY,
message_id INTEGER REFERENCES messages(id),
processed_at TIMESTAMP NOT NULL
);


{"content": "test2", "status":"yraaaaa"} {"content": "test3", "status":"yraaaaa"} {"content": "test4", "status":"yraaaaa"}{"content": "test5", "status":"yraaaaa"}
