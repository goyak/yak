# Yakety application keeper

Refer detail information in docs. https://gitlab.com/EasyStack/yakety/tree/master/docs
Or quick overview slide: https://gitpitch.com/Easystack/yakety/master?grs=gitlab&p=docs/talks

## How to build

```
 # install golang
 $ sudo yum install golang
 # set GOPATH
 # echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
 $ export GOPATH=$HOME/go

 # clone yak
 $ go get gitlab.com/EasyStack/yakety
 # get govendor
 $ go get -u github.com/kardianos/govendor

 # switch to yakety dir
 $ cd $GOPATH/src/gitlab.com/EasyStack/yakety

 # sync
 $ govendor sync

 # build yak
 $ make build
```