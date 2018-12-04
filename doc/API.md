# API


## `200` OK
``` json
{
    "id": 123,
    "name": "jiang"
}
```
## `400` Client Error
``` json
{
    "errors": [
        {
            "field": "username",
            "type": "Invalid",
            "message": "username is required, and must be 6~30 chars"
        },
        {
            "field": "nickname",
            "type": "Occupied",
            "message": "the nickname is occupied"
        }
    ]
}
```
## `401` Unauthorized
## `500` Server Error

