# creamy-static

Barebones web-accessible file directory

## Building

```sh
go get
go build
```

## Running

```sh
FILE_PATH=./foo/bar \
PORT=80 \
WATERMARK="creamy-videos was here" \
./creamy-static
```

## Running with Docker

```sh
docker run --rm \
-e FILE_PATH=/static \
-p 80:80 \
-v $PWD/static:/static \
albinodrought/creamy-static
```
