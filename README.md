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
settimeout [-addr=:80] [-tcpaddr=:5103]
```

