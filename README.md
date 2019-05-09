# http_jsondumper

## Build and run

```shell
go build
./http_jsondumper
```

## Use

```shell
curl -X POST http://localhost:8080 -H "Content-Type:application/json" -d "$(cat 0000.json)"
```