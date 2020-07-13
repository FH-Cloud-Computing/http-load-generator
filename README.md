# Load generating HTTP server

This is a HTTP server that creates CPU load for every request
to `/load`. It also includes a health check on `/health`.

Run it:

```
docker run -p 8080:8080 janoszen/http-load-generator
```
