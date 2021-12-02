# Authentication Microservise

To run use
```
docker build -t auth_ms .
docker run -d -P -p 8080:8080  auth_ms
```
In this case you should use ```localhost:8080```

## Endpoints

Post User: ```/user``` \
Get User:```/user```\
Get User by number:```/user/number```\
Get Users with number:```/user/find/number1,number2,...,numbern```
