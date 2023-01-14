# Snippetbox

### Routing requests

| Method | Pattern            | Handler       | Action                     |
| ------ | ------------------ | ------------- | -------------------------- |
| ANY    | /                  | home          | Display the home page      |
| ANY    | /snippet/view?id=1 | snippetView   | Display a specific snippet |
| POST   | /snippet/create    | snippetCreate | Create a new snippet       |

```bash
$ curl -i -X POST http://localhost:4000/snippet/create
HTTP/1.1 200 OK
Date: Sat, 14 Jan 2023 20:54:58 GMT
Content-Length: 23
Content-Type: text/plain; charset=utf-8

Create a new snippet...
```

```bash
$ curl -i -X GET http://localhost:4000/snippet/create
HTTP/1.1 405 Method Not Allowed
Allow: POST
Date: Sat, 14 Jan 2023 21:04:56 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

Method Not Allowed
```

