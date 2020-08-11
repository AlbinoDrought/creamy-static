# Creamy Static

<a href="https://hub.docker.com/r/albinodrought/creamy-static">
<img alt="albinodrought/creamy-static Docker Pulls" src="https://img.shields.io/docker/pulls/albinodrought/creamy-static">
</a>
<a href="https://github.com/AlbinoDrought/creamy-static/blob/master/LICENSE"><img alt="AGPL-3.0 License" src="https://img.shields.io/github/license/AlbinoDrought/creamy-static"></a>


No-fluff [`pkg/net/http` FileServer](https://golang.org/pkg/net/http/#FileServer) wrapper
for serving a local directory on the web

## Usage

- `FILE_PATH`: location to load files from, defaults to `./static` (effectively `/static` in container)

- `PORT`: port to listen on, defaults to `80`

- `WATERMARK`: text to show on HTML pages as a watermark, disabled by default

- `BASIC_AUTH_USERNAME`: enable basic auth with this username

- `BASIC_AUTH_PASSWORD`: enable basic auth with this password

### With Docker

```sh
docker run --rm -it \
    -p 80:80 \
    -v $(pwd)/foo/bar:/static \
    -e WATERMARK="creamy-static was here" \
    albinodrought/creamy-static
```

### Without Docker

```sh
FILE_PATH=./foo/bar \
PORT=80 \
WATERMARK="creamy-static was here" \
./creamy-static
```

## Building

### With Docker

```sh
docker build -t albinodrought/creamy-static .
```

### Without Docker

```sh
go get && go build
```
