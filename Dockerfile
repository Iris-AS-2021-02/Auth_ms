#grab base image
FROM golang:1.17

WORKDIR /arqui/project
COPY ./arqui/project . 

RUN go mod tidy
RUN go build -o auth
RUN chmod +777 auth
RUN export GIN_MODE=release

EXPOSE 8080

CMD ["./auth"]

