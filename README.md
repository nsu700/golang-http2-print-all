# golang-http2-print-all

## How to run the docker image
```
root@test:~/go# docker run -d -P -e KEY_FILE=/key.pem -e CERT_FILE=/cert.pem -v /root/go/cert.pem:/cert.pem -v /root/go/key.pem:/key.pem my-golang-http2-server
a593e28cf1b68e6cb3479590885b501e9a3423d79703c975a623ba3e1bb0eaa7
```

## Expected output
```
root@test:~/go# docker ps
CONTAINER ID   IMAGE                    COMMAND   CREATED         STATUS         PORTS                                           NAMES
a593e28cf1b6   my-golang-http2-server   "./app"   2 seconds ago   Up 2 seconds   0.0.0.0:32771->8080/tcp, [::]:32771->8080/tcp   busy_cerf
root@test:~/go# curl -Ik https://localhost:32771
HTTP/2 200
content-type: text/plain; charset=utf-8
content-length: 14
date: Sat, 16 Nov 2024 17:31:23 GMT

root@test:~/go# curl -Ik https://localhost:32771 -H "Host:abc.com"
HTTP/2 200
content-type: text/plain; charset=utf-8
content-length: 14
date: Sat, 16 Nov 2024 17:31:30 GMT

root@test:~/go# docker logs a593e28cf1b6
REQUEST:
HEAD / HTTP/2.0
Host: localhost:32771
Accept: */*
User-Agent: curl/7.74.0

REQUEST:
HEAD / HTTP/2.0
Host: abc.com
Accept: */*
User-Agent: curl/7.74.0
```
