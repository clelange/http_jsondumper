# http_jsondumper

## Build and run

```shell
go build
./http_jsondumper
```

## Use

```shell
curl -X POST http://localhost:8080 -H "Content-Type:application/json" -d "$(cat 0000.json)"
find /Volumes/Higgs4l/higgs4lbucket/jsondata/eventselection/cms_run2012b_doublemuparked_aod_22jan2013-v1 -iname '*.json'|parallel --bar -j 50 curl -s -X POST http://localhost:8080 -H "Content-Type:application/json"  -d@{}
ab -p 0000.json -T application/json -c 200 -n 30000 http://localhost:8080/
```