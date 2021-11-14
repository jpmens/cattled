## cattled

```console
$ ansible alice -m uri -a "url=http://localhost:10000/cow"
alice | SUCCESS => {
    "changed": false,
    "connection": "close",
    "content_length": "52",
    "content_type": "application/json",
    "cookies": {},
    "cookies_string": "",
    "date": "Sun, 14 Nov 2021 09:45:18 GMT",
    "elapsed": 0,
    "json": {
        "breed": "Irish Moiled",
        "id": 30,
        "origin": "Ireland"
    },
    "msg": "OK (52 bytes)",
    "redirected": false,
    "status": 200,
    "url": "http://localhost:10000/cow"
}
```
