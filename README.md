## Simple Go HTTP server

Simple HTTP server that doesn't do anything else aside from showing in logs requests received. It randomizes the HTTP server response.

### Example

From one terminal:

```shell
$ make run
go build -o simple-http-server -v
./simple-http-server
2023/05/09 12:36:00 Server listening on 127.0.0.1:8080...
```

From another:

```shell
curl -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer my-bearer-token" \
  -d '{"name": "Raúl Barroso", "email": "raul@meroxa.io"}' \
  'http://127.0.0.1:8080/mypath?param1=value1&param2=value2'
```

Expected response:

```shell
👋
```


Expected server logs:

```shell
2023/05/09 12:36:02 Request path: /mypath
2023/05/09 12:36:02 Request query parameters:
2023/05/09 12:36:02 param1: value1
2023/05/09 12:36:02 param2: value2
2023/05/09 12:36:02 Request headers:
2023/05/09 12:36:02 Content-Length: 52
2023/05/09 12:36:02 User-Agent: curl/7.87.0
2023/05/09 12:36:02 Accept: */*
2023/05/09 12:36:02 Content-Type: application/json
2023/05/09 12:36:02 Authorization: Bearer my-bearer-token
2023/05/09 12:36:02 Request body: {"name": "Raúl Barroso", "email": "raul@meroxa.io"}
```

## Other options

You could specify the name of a specific service in case you want to customize the response. e.g.:

```shell
make run SERVICE=discord # discord is yet to be implemented
```