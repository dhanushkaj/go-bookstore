from golang:1.20.0-alpine3.17

RUN mkdir /app
WORKDIR /APP

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build cmd/main/main.go
RUN mv main /usr/local/bin

#default env
ENV APP_ENV=prep

ENTRYPOINT ["main"]

EXPOSE 3000