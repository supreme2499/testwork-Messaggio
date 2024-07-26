# HTTP-API-SERVER

Это выполнение моего тестового задания:

***rest-api-serv***

**Для создания этого репозитория я использовал несколько библеотек:**
1. github.com/julienschmidt/httprouter - удобный http роутер(мультиплексор)
2. github.com/sirupsen/logrus - удобный и красивый логер
3. github.com/ilyakaznacheev/cleanenv - минималистичный и простой конфигуратор
4. go get github.com/jackc/pgx/v5 - драйвер постгрес

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

# Postges in docker:

Для развёртывания БД в контейнере я использовал офицальный драйвер Postgres для docker

```docker pull postgres```

***создал контейнер:***
```
docker run --name post_container -e POSTGRES_PASSWORD=verysecret -p 5438:5432 -d postgres

docker exec -ti post_container createdb -U postgres testwork

docker exec -ti post_container psql -U postgres
```

***таблица:***

```SQL
    CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL
    );  
```
и получил вот это
```Docker
Databases name: testwork
Username: postgres
Password: verysecret
port: 5438->5432
```
