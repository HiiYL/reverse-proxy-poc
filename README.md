# Reverse Proxy PoC

The project contains:
1. File upload destination service written in flask.
2. Proxy service in Go.

# Prereq
To run, Postman or any http client is required for posting file to the proxy.

# How to use
1. run flask server. `python server.py`
2. run go proxy. `go run *.go`
3. use Postman to post a file to proxy service at `localhost:5001`.

No elevated memory usage should be observed
