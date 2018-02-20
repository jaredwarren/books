# Book App

A simple CRUD book storage API


## Build

### Binary
Build using the following command

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -o books -ldflags '-w' .
``` 

### Docker Image
```
docker build -t jlwarren1/books .
```

## Run
```
docker run -d -P jlwarren1/books
```