目標讓 yak 本身可以自我更新

```
 # yak install yak
```

example yak.yml

```
name: yak
summary: Yak application kit
version: 1.0

backend: plain
source: http://mirror.easystack.io/yak/yak-1.0.gz
sha256sum: 87428fc522803d31065e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7

description: |
  Yak application kit
```

## Backend: Plain
放置位置可能為FTP或HTTP(S)提供直接下載,以checksum作為確認

## Specific Hook or Plugin

### check_os
確認系統是否為 ostree or yak-core or classic

## 變數
### PATH
加上 YAK_BIN_PATH, 可能不在yak本身設定需要加進相關shell設定檔案

### YAK_BIN_PATH
執行檔放置位置 系統需要可讀寫

### YAK_RECIPE_PATH
yak recipe 系統需要可讀寫