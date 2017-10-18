# Yak/Yakety application kit for DevOps. (WIP)

Refer detail information in docs. https://gitlab.com/EasyStack/yakety/wikis
Or quick overview slide: https://gitpitch.com/Easystack/yakety/master?grs=gitlab&p=docs/talks

## How to build

```
 # install golang
 $ sudo yum install golang
 # set GOPATH
 # echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
 $ export GOPATH=$HOME/go

 # clone yak
 $ mkdir -p $GOPATH/src/gitlab.com/EasyStack/
 $ cd $GOPATH/src/gitlab.com/EasyStack/
 $ git clone git@gitlab.com:EasyStack/yakety

 # get govendor

 $ go get -u github.com/kardianos/govendor

 # sync

 $ govendor sync

 # build yak
 $ make build
```

## Roadmap

 * 实做 local keep single commit
   - ostree-go

## govendor

Dependencies are handled via `govendor`. Get it via:

    go get -u github.com/kardianos/govendor

And then, run:

    govendor sync
