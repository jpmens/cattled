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

### with basic auth (`/secret`)

return a `text/plain` payload which also contains the hostname

```console
$ curl -i -u cow:moo localhost:10000/secret
HTTP/1.1 200 OK
Date: Mon, 02 Oct 2023 10:16:52 GMT
Content-Length: 41
Content-Type: text/plain; charset=utf-8

Here is the magic ...: rabbit.ww.mens.de
```
