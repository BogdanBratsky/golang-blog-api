# Документация API для блога

###### Основные возможности

***

## Эндпоинты для работы с пользователями

***

#### POST api/register
>Создаёт нового пользователя.

**METHOD**: POST
**URL**: api/register

**request:** 

>POST api/register

>Content-type: application/json
```
{
    "userName": "User_Name",
    "userEmail": "user@email.com",
    "userPassword": "ExampleOfPassword2048"
}
```

**response (201):**
>Content-type: application/json

```
{
    "success": true,
    "message": "Пользователь успешно создан",
    "user": {
        "userId": 1,
        "userName": "@User_Name",
        "userEmail": "user@email.com",
        "createdAt": "2024-02-25T12:00:00Z"
    }
}
```

**response (400):**
>Content-type: application/json

```
{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**
>Content-type: application/json

```
{
    "success": false,
    "message": "Ресурс не найден"
}
```

**response (500):**
>Content-type: application/json

```
{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```
---
#### POST api/login
>Эндпоинт для авторизации.

**METHOD**: POST
**URL**: api/login



**request:** 
```
POST api/login
```
>Content-type: application/json
```
{
    "userName": "User_Name",
    "userPassword": "ExampleOfPassword2048"
}
```

**response (200):**
>Content-type: application/json

```
{
    "success": true,
    "message": "Пользователь успешно авторизован",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "user": {
        "userId": 1,
        "userName": "@User_Name",
        "userEmail": "user@email.com",
        "createdAt": "2024-02-25T12:00:00Z"
    }
}
```

**response (400):**
>Content-type: application/json

```
{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**
>Content-type: application/json

```
{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**
>Content-type: application/json

```
{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```
---

#### GET api/users
>Возвращает список пользователей. Если не указать в квери-параметрах 'page' и 'per_page', то в ответе будет возвращено 50 последних зарегистрировавшихся пользователей. В 'per_page' можно передать от 1 до 100. Если в этот параметр передано больше 100, то в ответе вернётся ошибка 404.

**METHOD**: GET
**URL**: api/users



**request:**
```
GET api/users
```
**or**
```
GET api/users?page=1&per_page=100
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "items": [
        {
            "userId": 1,
            "userName": "@User_Name",
            "createdAt": "2024-02-25T12:00:00Z",
        },
        ...
        // прочие пользователи
    ],
    "pagination": {
        "page": 1,
        "perPage": 100,
        "pagesCount": 50,
    },
    "totalCount": 5000
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---
#### GET api/users/:id
>Возвращает конкретного пользователя по указанному в запросе идентификатору. 

**METHOD**: GET
**URL**: api/users/:id



**request:**
```
GET api/users/1
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "user": {
        "userId": 1,
        "userName": "@User_Name",
        "createdAt": "2024-02-25T12:00:00Z",
    }
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---

## Эндпоинты для работы с записями

---

#### POST api/articles
>Создаёт новую запись, если запрос выполняется авторизованным пользователем.

**METHOD**: POST
**URL**: api/articles

**request:**
```
POST api/articles
```

>Authorization: Bearer {{token}}
>Content-type: application/json
```
{
    "categoryId": 1,
    "articleTitle": "Title",
    "articleDescription": "Description",
    "articleText": "Text",
}
```
**response (201):**

>Content-type: application/json
```

{
    "success": true,
    "message": "Запись успешно создана",
    "article":  {
        "articleId": 3,
        "authorId": 1,
        "categoryId": 1,
        "articleTitle": "Title",
        "articleDescription": "Description",
        "articleText": "Text",
        "articleViews": 125,
        "timeToRead": 3000,
        "edited": false,
        "createdAt": "2024-02-25T12:00:00Z",
    },
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (401):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Необходимо авторизоваться"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---
#### PATCH api/articles/:id
>Редактирует конкретную существующую запись, если запрос выполняется авторизованным пользователем и если запись создана этим пользователем. Редактировать можно поля "articleTitle", "articleDescription" и "articleText".

**METHOD**: PATCH
**URL**: api/articles/:id



**request:**
```
PATCH api/articles/3
```
>Authorization: Bearer {{token}}
>Content-type: application/json
```
{
    "articleTitle": "Title Edited",
    "articleDescription": "Description Edited",
    "articleText": "Text Edited",
},
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "message": "Запись с id=3 успешно отредактирована",
    "article":  {
        "articleId": 3,
        "authorId": 1,
        "categoryId": 1,
        "articleTitle": "Title Edited",
        "articleDescription": "Description Edited",
        "articleText": "Text Edited",
        "articleViews": 125,
        "timeToRead": 3000,
        "edited": true,
        "createdAt": "2024-02-25T12:00:00Z",
    }
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (401):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Необходимо авторизоваться"
}
```
**response (403):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Нет прав на это действие"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---
#### DELETE api/articles/:id
>Удаляет конкретную существующую запись, если запрос выполняется авторизованным пользователем и если запись создана этим пользователем.

**METHOD**: DELETE
**URL**: api/articles/:id



**request:**
```
DELETE api/articles/5
```
>Authorization: Bearer {{token}}

**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "message": "Запись с id=5 успешно удалена"
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (401):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Необходимо авторизоваться"
}
```
**response (403):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Нет прав на это действие"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---
#### GET api/articles
>Возвращает список записей безотносительно к авторству этих записей. 
>>Есть параметр order, принимающий значения asc и desc. Если не указать этот параметр, то по умолчанию будут выдаваться записи, начиная с самых новых. 
>>
>>Также есть параметры 'page' и 'per_page', реализующие пагинацию. Если не указать эти параметры, то в ответе будут возвращены 50 последних публикаций (если нет параметра order со значением asc). В 'per_page' можно передать значение от 1 до 100. Если в этот параметр передано значение больше 100, то в ответе вернётся ошибка 404.

**METHOD**: GET
**URL**: api/articles



**request:**
```
GET api/articles
```
**or**
```
GET api/articles?order=asc
```
**or**
```
GET api/articles?page=1&per_page=100
```
**or**
```
GET api/articles?page=1&per_page=100&order=desc
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "items": [
        {
            "articleId": 99,
            "authorId": 1,
            "categoryId": 1,
            "articleTitle": "Title",
            "articleDescription": "Description",
            "articleText": "Text",
            "articleViews": 125,
            "timeToRead": 2000,
            "edited": false,
            "createdAt": "2024-02-25T12:00:00Z",
        },
        ...
        // прочие записи
    ],
    "pagination": {
        "page": 1,
        "perPage": 100,
        "pagesCount": 50,
    },
    "totalCount": 5000
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---
#### GET api/articles/:id
>Возвращает конкретную запись по идентификатору.

**METHOD**: GET
**URL**: api/articles/:id

**request:**
```
GET api/articles/3
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "articles":  {
        "articleId": 3,
        "authorId": 1,
        "categoryId": 1,
        "articleTitle": "Title",
        "articleDescription": "Description",
        "articleText": "Text",
        "articleViews": 125,
        "timeToRead": 3000,
        "edited": true,
        "createdAt": "2024-02-25T12:00:00Z",
    },
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```
---
#### GET api/users/:id/articles
>Возвращает список записей конкретного пользователя по идентификатору этого пользователя. 
>>Есть параметр order, принимающий значения asc и desc. Если не указать этот параметр, то по умолчанию будут выдаваться записи, начиная с самых новых. 
>>
>>Также есть параметры 'page' и 'per_page', реализующие пагинацию. Если не указать эти параметры, то в ответе будут возвращены 50 последних публикаций (если нет параметра order со значением asc). В 'per_page' можно передать значение от 1 до 100. Если в этот параметр передано значение больше 100, то в ответе вернётся ошибка 404.

**METHOD**: GET
**URL**: api/users/:id/articles

**request:**
```
GET api/users/4/articles
```
**or**
```
GET api/users/4/articles?order=asc
```
**or**
```
GET api/users/4/articles?page=1&per_page=100
```
**or**
```
GET api/users/4/articles?page=1&per_page=100&order=desc
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "items": [
        {
            "articleId": 3,
            "authorId": 4,
            "categoryId": 1,
            "articleTitle": "Title",
            "articleDescription": "Description",
            "articleText": "Text",
            "articleViews": 125,
            "timeToRead": 2000,
            "edited": false,
            "createdAt": "2024-02-25T12:00:00Z",
        },
        ...
        // прочие записи
    ],
    "pagination": {
        "page": 1,
        "perPage": 100,
        "pagesCount": 50,
    },
    "totalCount": 5000
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---
#### GET api/categories/:id/articles
>Возвращает список записей конкретной категории по идентификатору этой категории. 
>>Есть параметр order, принимающий значения asc и desc. Если не указать этот параметр, то по умолчанию будут выдаваться записи, начиная с самых новых. 
>>
>>Также есть параметры 'page' и 'per_page', реализующие пагинацию. Если не указать эти параметры, то в ответе будут возвращены 50 последних публикаций (если нет параметра order со значением asc). В 'per_page' можно передать значение от 1 до 100. Если в этот параметр передано значение больше 100, то в ответе вернётся ошибка 404.

**METHOD**: GET
**URL**: api/categories/:id/articles

**request:**
```
GET api/categories/7/articles
```
**or**
```
GET api/categories/7/articles?order=asc
```
**or**
```
GET api/categories/7/articles?page=1&per_page=100
```
**or**
```
GET api/categories/7/articles?page=1&per_page=100&order=desc
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "items": [
        {
            "articleId": 3,
            "authorId": 1,
            "categoryId": 7,
            "articleTitle": "Title",
            "articleDescription": "Description",
            "articleText": "Text",
            "articleViews": 125,
            "timeToRead": 2000,
            "edited": false,
            "createdAt": "2024-02-25T12:00:00Z",
        },
        ...
        // прочие записи
    ],
    "pagination": {
        "page": 1,
        "perPage": 100,
        "pagesCount": 50,
    },
    "totalCount": 5000
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---

### Эндпоинты для работы с категориями

***

#### GET api/categories
>Возвращает список всех категорий.

**METHOD**: GET
**URL**: api/categories

**request:**
```
GET api/categories
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "items":  [
        {
            category_id: 1,
            category_name: "Политика",
        },
        {
            ...
        },
        ...,
        {
            category_id: 10,
            category_name: "IT",
        },
    ]
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---

#### GET api/categories/:id
>Возвращает конкретную категорию по её идентификатору.

**METHOD**: GET
**URL**: api/categories/:id

**request:**
```
GET api/categories/10
```
**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "category": {
        {
            category_id: 10,
            category_name: "IT",
        }
    }
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---

### Эндпоинты для работы с подписками

***

#### POST api/subscriptions

>Создаёт подписку, если пользователь авторизован.

**METHOD**: POST
**URL**: api/subscriptions

**request:**
```
POST api/subscriptions
```
>Content-type: application/json
>Authorization: Bearer {{token}}
```
{
    "subscriber_id": 1,
    "target_user_id": 2
}
```
**response (201):**

>Content-type: application/json
```

{
    "success": true,
    "message": "Подписка успешно создана",
    "subscription": {
        "subscription_id": 1,
        "subscriber_id": 1,
        "target_user_id": 2,
    }
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (401):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Необходимо авторизоваться"
}
```
**response (403):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Нет прав на это действие"
}
```
**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---

#### DELETE api/subscriptions

>Удаляет конкретную подписку, если пользователь авторизован и если эта подписка создана этим пользователем.

**METHOD**: DELETE
**URL**: api/subscriptions/:id

**request:**
```
DELETE api/subscriptions/1
```
>Content-type: application/json
>Authorization: Bearer {{token}}

**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "message": "Подписка с id=1 успешно отменена",
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (401):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Необходимо авторизоваться"
}

```
**response (403):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Нет прав на это действие"
}
```

**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```
{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---

#### GET api/users/:id/subscriptions

>Возвращает список подписок конкретного пользователя по идентификатору этого пользователя.

**METHOD**: GET
**URL**: api/users/:id/subscriptions

**request:**
```
GET api/users/2/subscriptions
```
>Content-type: application/json
>Authorization: Bearer {{token}}

**response (200):**

>Content-type: application/json
```

{
    "success": true,
    "message": "",
    "items": [
        {
            "subscription_id": 2,
            "subscriber_id": 2,
            "target_user_id": 1,
        }
        ...
        // прочие подписки
    ],
    "pagination": {
        "page": 1,
        "perPage": 100,
        "pagesCount": 50,
    },
    "totalCount": 5000
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (401):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Необходимо авторизоваться"
}

```
**response (403):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Нет прав на это действие"
}
```

**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```
{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---

#### GET api/users/:id/followers

>Возвращает список подписчиков конкретного пользователя по идентификатору этого пользователя.

**METHOD**: GET
**URL**: api/users/:id/followers

**request:**
```
GET api/users/2/followers
```
>Content-type: application/json
>Authorization: Bearer {{token}}

**response (200):**

>Content-type: application/json
```
{
    "success": true,
    "message": "",
    "items": [
        {
            "subscription_id": 2,
            "subscriber_id": 1,
            "target_user_id": 2,
        }
        ...
        // прочие подписки
    ],
    "pagination": {
        "page": 1,
        "perPage": 100,
        "pagesCount": 50,
    },
    "totalCount": 5000
}
```

**response (400):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Некорректный запрос"
}
```
**response (401):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Необходимо авторизоваться"
}

```
**response (403):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Нет прав на это действие"
}
```

**response (404):**

>Content-type: application/json
```

{
    "success": false,
    "message": "Ресурс не найден"
}
```
**response (500):**

>Content-type: application/json
```
{
    "success": false,
    "message": "Внутренняя ошибка сервера"
}
```

---