# yak application keeper

Refer detail information in docs. https://goyak.io

## How to build

```
 # install golang
 $ sudo yum install golang
 # set GOPATH
 # echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
 $ export GOPATH=$HOME/go

 # clone yak
 $ go get github.com/goyak/yak
 # get govendor
 $ go get -u github.com/kardianos/govendor

 # switch to yak dir
 $ cd $GOPATH/src/github.com/goyak/yak

 # sync
 $ govendor sync

 # build yak
 $ make build
```
---

Project sponsored by EasyStack Inc.
