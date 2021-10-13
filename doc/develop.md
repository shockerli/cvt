# develop

## test in 32bit

> use `i386/golang:alpine` image for testing 32bit system, detail in `Dockerfile.32bit`

- build image

```shell
docker build -t cvt-test -f Dockerfile.32bit .
```

- run container

```shell
docker run -it --name cvt-test -v "$PWD":/usr/src/myapp -w /usr/src/myapp cvt-test bash
```

- run test

```shell
go test ./...
```
