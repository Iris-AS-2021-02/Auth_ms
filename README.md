# Authentication Microservise

To run use
```
docker build -t auth_ms .
docker run -d -P -p 8080 auth_ms
```
Use the port used by the machine.
```
docker ps
CONTAINER ID   IMAGE     COMMAND      CREATED         STATUS         PORTS                     NAMES
e8bc6917d2c8   auth_ms   "./auth &"   9 seconds ago   Up 4 seconds   0.0.0.0:51620->8080/tcp   reverent_sanderson
```
In this case you should use ```localhost:51620```

## Endpoints

Post User.
```user```

Get User.

```/user```
Get User by number.

```//user/number```
Get Users with number.

```/user/find/number1,number2,...,numbern```
