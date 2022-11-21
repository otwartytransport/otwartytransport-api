# otwartytransport-api

## Run

To run the application specify environment using `APP_ENV` (possible: `dev`, `prod`).

```
APP_ENV=dev go run cmd/server.go
```

## Hot reload

If you want to use hot reload first install [cosmtrek/air](https://github.com/cosmtrek/air) in your system.
Then simply type

```
air
```

in the root of the project.

## Build Docker image

- Build the application

```
CGO_ENABLED=0 go build -o ./build/server cmd/server.go
```

- Build Docker image

```
docker build --tag otwartytransport-api .
```
