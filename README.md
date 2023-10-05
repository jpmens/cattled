## cattled

a miniscule http server for Ansible and AWX trainings in order to test the `uri` module and custom credentials.

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
$ curl -i -u cow:milk localhost:10000/secret
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 05 Oct 2023 14:49:37 GMT
Content-Length: 163

{"url":"https://en.wikipedia.org/wiki/Hey_Diddle_Diddle","rhyme":"Hey diddle diddle, The cat and the fiddle, The cow jumped over the moon","instance":"rabbit"}
```
