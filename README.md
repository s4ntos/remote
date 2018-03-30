Remote APP
=======

#### A web application built in [Go](http://golang.org) on top of the [Revel Web Framework](https://revel.github.io) using work done by [richtr](https://github.com/richtr) on  [Baseapp](https://github.com/richtr/baseapp) ####


Remote is a web application that provides web application for web remote desktop.

* Basic pages (Home, About Us, Contact Us, etc) working
* First roles configuration working


To start application:

```bash
export GOPATH=`pwd`
go get -v -u github.com/revel/cmd/revel
export PATH=$PATH:$GOPATH/bin
go get github.com/s4ntos/remote
revel run github.com/s4ntos/remote
```

On windows you will need to install a GCC, this as been tested with [TDM-GCC](http://tdm-gcc.tdragon.net/download)
