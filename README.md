# API для работы со списком объявлений
Данный сервис реализован на языке Go с использованием библиотеки HTTP. Для работы с PostgreSQL использовался драйвер pgx. Для маршрутизации использовался фреймворк gorilla/mux. Для создания объявлений неободимо авторизоваться и ввести полученный JWT токен в поле заголовка, для отображения не нужно авторизовываться. Если пользователь авторизован при отображении списка, то выведутся только созданные им объявления. Для данных из файла конфигурации используется Viper.
Данные принимаются в JSON формате, далее сохраняются в базу данных PostgreSQL. Вывод данных происходит тоже в формате JSON. Реализованы операции сохранения и получения.
Для создания объявлений необходима авторизация, для отображения нет. При авторизации при отображении ленты объявлений отображается их принадлежность пользователю.
Для постраничной навигации необходимо передать значение page в ссылке-запросе. Также, можно передать значения для желаемой сортировки order и направление orderby, если она отлична от сортировки по самой свежей дате добавления.

## Есть 2 способа запуска микросервиса:
### 1. Локально.
   Необходимо в файле конфигурации configs/config.yaml указать
   ```
   host = "localhost"
   ```
   Запустить контейнер postgres
   ```
   docker run --name=db -e POSTGRES_PASSWORD='54321' -p 5432:5432 -d postgres
   ```
   При наличии утилиты golang-migrate после первого запуска контейнеров прописать команду
   ```
   make migrate
   ```
   Она создает необходимые для работы сервиса таблицы в контейнере базы данных. Данную команду нужно прописать всего 1 раз, и далее при запуске тех же контейнеров таблицы сохранятся.

   Ввести в консоль команду
   ```
   go run cmd/main.go
   ```
   При необходимости, можно заменить названия контейнера базы данных, пароль и порты. Соответствующие параметры для такого же изменения находятся в configs/config.yaml.
### 2. Локально при использовании docker-compose.
   
   Необходимо в файле конфигурации configs/config.yaml указать (По умолчанию в проекте стоит такое значение)
   ```
   host = "db"
   ```
   Ввести в консоль команду
   ```
   make build && make run
   ```
   При наличии утилиты golang-migrate после первого запуска контейнеров прописать команду
   ```
   make migrate
   ```
Информация по утилите golang-migrate находится в репозитории https://github.com/golang-migrate/migrate
## Пользование сервисом
### 1. Авторизация и регистрация
#### Для регистрации необходимо выполнить запрос
```
curl --location  --request POST 'http://localhost:8000/api/auth/sign-up' \
--header 'Content-Type: application/json' \
--data '{
    "username": "{username}",
    "password": "{password}"
}'
```
Вместо username вводится желаемый username, в поле password соответственно желаемый пароль. 
#### Для авторизации необходимо выполнить запрос
```
curl --location  --request POST 'http://localhost:8000/api/auth/sign-in' \
--header 'Content-Type: application/json' \
--data '{
    "username": "{username}",
    "password": "{password}"
}'
```
Вместо username вводится выбранный нами при регистрации username, в поле password соответственно пароль.
В ответ на данный запрос нам выдастся токен, который нужно сохранить и использовать во всех следующих запросах. В программе Postman имеется функционал, который позволяет один раз указать токен и выполнять все дальнейшие запросы уже с ним. В командной строке с каждым запросом придется указывать вручную заголовок.
Проверка токена в сервисе выполняется при помощи методов в Middleware.
### 2. Объявления
#### Для получения списка объявлений необходимо ввести запрос
```
curl --location --request GET 'http://localhost:8000/api/sellings?order=date&sortby=name&page=1' \
--header 'Authorization: Bearer {token}' \
--data ''
```
Вместо Token в заголовке можно ввести личный токен, полученный при авторизации, либо убрать заголовок. При указании токена отобразятся только созданные пользователем объявления. Также, можно указать поле, по которому отобразится отсортированный список, и направление по возрастанию и убыванию. Возможные значения для order это Title, Price, Date для sortby ASC и DESC. Значение page нужно указывать в численном формате. Если не указывать эти три параметра по умолчанию будет сортировка по дате по убыванию и страница будет 1.
#### Для создания объявления необходимо ввести запрос
```
curl --location --request POST 'http://localhost:8000/api/create-selling' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token} \
--data '{
    "title": {"title"},
    "description": {"description"},
    "url": {"url"},
    "price": {1}
}'
```
Здесь нужно обязательно указать токен, иначе не получится создать объявление. В полях title, description, url, price нужно указать данные о названии, описании, ссылке на изображение, цене соответственно. Поле времени создания автоматически заполняется текущим временем. 

## Обработка ошибок
Для различных методов и вызовов функций реализована обработка ошибок, в зависимости от категории ошибки, выдается текст и код ошибки. Присуствуют коды 4хх и 5хх.
