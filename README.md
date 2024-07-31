# HTTP-API-SERVER

Это выполнение моего тестового задания:

# Инструкция по использования:
для отправки сообщений используйте curl запрос следующего вида:

``` 
curl -X POST http://localhost:8080/message \
     -H "Content-Type: application/json" \
     -d '{"content": "message"}'
```
или в тело запроса добавьте json данного формата. так же в можете отправить сразу несколько сообщений отправив сразу несколько json друг за другом `{"content": "message1"}{"content": "message2"}{"content": "message3"}`

для получения статистики по обработанным сообщения 
```
http://localhost:8080/statistics
```

***rest-api-serv***

**Для создания этого репозитория я использовал несколько библеотек:**
1. go get github.com/julienschmidt/httprouter - удобный http роутер(мультиплексор)
2. go get github.com/sirupsen/logrus - удобный и красивый логер
3. go get github.com/ilyakaznacheev/cleanenv - минималистичный и простой конфигуратор
4. go get github.com/jackc/pgx/v5 - драйвер постгрес
5. go get github.com/IBM/sarama для работы с kafka

___
# Логика сервера  
***Endpoints:***

POST /message --> post content
- Этот запрос получает сообщение в формате json, получаемое вместе с запросом. {"content":"какое-то сообщение"}, полученное сообщение запишется в Postgres, а далее отправится в kafka. Там оно просто поменяет статус на "обработано"

GET /statistics --> get statistics 
- Этот метод выводит статистику по обработанным сообщениям

___
# Логирование:
В качестве логов я использовал обёртку которую я уже ранее использовал.

Данная обёртка выведет лог в консоль и запишет в нужный нам файл, в данном случае в `all.log` 

# Конфигурация
Конфиг написан с помощью библеотеки cleanenv, в конфиг я передаю необходимые мне параметры для подключения к бд


***таблица:***

```SQL
    CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created TIMESTAMP NOT NULL
    );  

        CREATE TABLE worker (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created TIMESTAMP NOT NULL
    ); 
```
и получил вот это
```Docker
Databases name: testwork
Username: postgres
Password: verysecret
port: 5432->5432

```

После получения сообщения оно записывается в бд(постгрес)
затем отправляется в кафку, там оно записывается в другую таблицу(обработанных сообщений), от туда берётся кол-во обработанных сообщений

# Dockerfile
Для создания докер контейнера используйте данные команды
```
docker-compose up --build
docker-compose down -v
```

# Баги и ошибки
Я знаю что криво обработал ошибки, но у меня уже не было времени  это исправлять.

Кое где я в функциях возвращаю ошибки, но из-за не внимательности я их вывожу в лог фатал, вместо того что бы прокинуть их выше... 