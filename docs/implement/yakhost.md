# yak host


UPX – the Ultimate Packer for eXecutables
--brute

client 部份用 cobra
deamon 部份用 flag 原生


cmds:

cmd/snap/cmd_snap_op.go

> text template replace added-files

## system commands
 - alias

### yak install core, command alias

example command: core-info
```
 $ core-info
Yak: 0.01
```


## yakd or supervisord - process and status control

 * docker runwrapper
 * local yak store

## yak rest-api

## yak web ui

## yak cli

### yak Commands

* yak fetch: 下载 recipes 至本机cache

```
 # yak fetch
 # yak xxx
```

* yak info core:

```
 # yak info core
 Yak is not run on a ostree based host

 # yak info core
name:      core
summary:   "snapd runtime environment"
publisher: canonical
contact:   snappy-canonical-storeaccount@canonical.com
description: |
  The core runtime environment for snapd
type:        core
snap-id:     99T7MUlRhtI3U0QFgl5mXXESAiSwt776
tracking:    stable
installed:   16-2.27.6 (2844) 85MB core
refreshed:   2017-09-07 17:09:21 +0800 CST
channels:                                    
  stable:    16-2.27.6                (2844) 85MB -
  candidate: 16-2.27.6                (2898) 85MB -
  beta:      16-2.27.6                (2898) 85MB -
  edge:      16-2.27.6+git365.f5f40b0 (2908) 87MB -
```

## 使用场景规划

### 单机 ostree
```
# yak fetch # update 
```

### user 单机

```
# yak install gitea # docker pull gitea/HASH_b6a57956d1be7548aa625b38c1382fc8

# yak config gitea # 秀出所有所有变数

# yak config gitea set DATA_PATH=/srv/gitea # 设定变数

# systemctl gitea start # 设定systemd config

# gitea help # docker run gitea/HASH_b6a57956d1be7548aa625b38c1382fc8 help 另外带一些设定进去

# yak try (upgrade) --channel beta gitea # 升级至 beta channel 版本

# yak host register #
```

### 企业集规划场景

```
# kubectl cluster-info

Kubernetes master is running at http://host01:8080

# yak env create staging --type=kubernetes CLUSTER_SERVICE=http://host01:8080

# yak install gitea --env staging --pod=5 #

# yak upgrade gitea --env staging
Container OS

# yak config core set XXX=abc

# yak env create staging --type=group cluster_a

# yak env add-host node-1 node-2 node-3

# yak upgrade core --env cluster_a
```
### Store 与 登入

```
# yak repo add fedora http://xxx

# yak repo list

default: http://yak.easystack.io

fedora: http://xxx

# yak find mysql

mysql <default>

beta: 1.2.123 mysql/b6a57956d1be7548aa625b38c1382fc

stable: 1.2 mysql/b6a57956d1be7548aa625b38c1382fc

mysql <fedora>

beta: 1.2.555 mysql/b6a57956d1be7548aa625b38c1382fc

stable: 1.2 mysql/b6a57956d1be7548aa625b38c1382fc

# yak login shawn
```
