# Authentication Microservise

To run use
```
docker build -t auth_ms .
docker run -d -P -p 1024:1024  auth_ms
```
In this case you should use ```localhost:1024```

## Endpoints

Post User: ```/user``` \
Get User:```/user```\
Get User by number:```/user/number```\
Get Users with number:```/user/find/number1,number2,...,numbern```
