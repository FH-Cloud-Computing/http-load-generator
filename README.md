# Load generating HTTP server

This is a HTTP server that creates CPU load for every request
to `/load`. It also includes a health check on `/health`. It listens on port 8080.

Run it:

```
docker run -p 8080:8080 janoszen/http-load-generator
```

Alternatively, you can [download the binary from the releases section](https://github.com/janoszen/http-load-generator/releases) and run it manually.
