# clean
Trying a boilerplate for clean-architecture in Golang


1. Create user
```json
curl --location --request POST 'localhost:8000/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "thedevsaddam",
    "password": "12345",
    "profile": {
        "name": "Saddam H",
        "age": 30
    }
}'
```

1. List users

```
curl --location --request GET 'localhost:8000/v1/users'
```

```json
{
    "data": [
        {
            "id": 1,
            "username": "thedevsaddam",
            "type": "",
            "profile": {
                "id": 1,
                "user_id": 1,
                "name": "Saddam H",
                "age": 30,
                "bio": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": null
            },
            "followers": [
                {
                    "login": "soyelmnd",
                    "id": 2678063
                },
                {
                    "login": "shorifulislam00",
                    "id": 10585612
                },
                {
                    "login": "truejoy",
                    "id": 6727845
                },
                ...
            ]
        }
    ]
}
```