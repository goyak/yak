# Application/Core - 主要作業系統本身

Container OS 目前與k8s有相當多的整合, 所以預設k8s包含在整個運作系統上

```
name: core
summary: Easystack Container Linux
version: '1.0'

# ${url}:${branch}@${from}
backend: ostree
source: http://mirror.centos.org/centos/7/atomic/x86_64/repo
branch: centos-atomic-host/7/x86_64/standard
backend: ostree
commit: 173278f2ccba80c5cdda4b9530e6f0388177fb6d27083dec9d61bbe40e22e064

description: |
  Easystack Container Linux
```

## Backend: Ostree, Docker(LinuxKit)
第一階段僅支援Ostree, 未來可以整併 Docker 或者應該叫LinuxKit


## Features
 * bcache(?)
 * kernel
 * busybox

## Specific Hook or Plugin
### build_image
### k8s related config