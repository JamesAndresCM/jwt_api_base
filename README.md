# jwt_api_base

- create public and private keys (inside keys folder)
```
openssl genrsa -out private.rsa 3072
```

```
openssl rsa -in private.rsa -pubout -out public.rsa.pub
```
- set env file (env_sample) as example

## available endpoints
- create user payload:
-  POST `localhost:8080/api/users/`
````
{
    "username": "test",
    "email": "test@domain.com",
    "fulname": "User Test",
    "password": "12345678",
    "confirmPassword": "12345678"
}
````

- login user payload
- POST `localhost:8080/api/login`
````
{
    "email": "test@domain.com",
    "password": "12345678"
}
````
