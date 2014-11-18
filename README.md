## settimeout

HTTP and TCP server that allows a client to specify how long the server should
sleep before responding. A live demo is at [http://settimeout.io/](http://settimeout.io/)

### Building

To build settimeout you just need to run:
```
go build
```

If you want to change the assets, you'll need [go-bindata](http://github.com/jteeuwen/go-bindata/).
After installing, you'll need to compile the assets after every change:
```
go-bindata assets/
```
Then you can build as usual.

### Usage

```
settimeout [-addr=:80] [-tcpaddr=:5103] [-statsaddr=127.0.0.1:5104] [-httpsaddr= -httpscrt= -httpskey=]
```
To disable HTTP, HTTPS, or TCP server, or the stats socket, just set their respective address to nothing.

To see stats:
```
(echo info && sleep 0.1) | nc 127.0.0.1 5104
```

### Other stuff

If you're getting an error such as:
```
Failed to start HTTP server: listen tcp :80: bind: permission denied
```
you're not running settimeout as root and therefore you need to make it listen on a higher port number.
You can then use iptables to redirect port 80 to your new port (example uses port 1080):

```
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 80 -j REDIRECT --to-ports 1080
```

If you want to run settimeout behind nginx you need to setup settimeout to listen on a different port
such as 1080. The addr can be `127.0.0.1:1080` then you just need to add the following to your nginx conf:
```
proxy_pass http://127.0.0.1:1080;
proxy_pass_header Server;
```
`proxy_pass_header` is optional and just tells nginx to not overwrite the Server header with "nginx".