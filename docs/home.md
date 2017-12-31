# Yak/Yakety app kit for DevOps. (WIP)
目前我们透过container与ostree等技术,让布署方面可以达到便利且稳定.
然而升级或是服务级设定需求依然无法轻易的达成.
我们将参考现行的一些服务设计一套meta-data (recipe) based的 app kit.
以application的角度提供DevOps便捷性.

**recipe = sources + hooks + config**

## 问题 - 处理 版本异动 与 设定
 * 版本异动 往往需要 相对应之环境配置
  * ostree (atomic) postcompose.sh
  * docker run 必须带许多参数 bind network, storage and config
 * 延伸 依据环境或设备 有不同配置条件
  - upgrade, rollback
  - try 用假资料 或 dump 与 load进去新的版本
  - prod or staging
 * 状态回报 health check
 * image 讯息不易显出, 如果一个image 提供多个以上 commands (目前docker run 预设只有一个command), 我们不容易了解image内部配置
 * CI/CD recipe, testing

## 解决方案- meta data
 * atomic core 管理 <规划系统可运作于所有可以执行docker或kubernetes环境, 但在atomic环境下可以管理机制>
  - cloud-init: api-key, client-server key exchange
 * image 与版本关联性 
  - PACKAGE or SERVICES PACK:USER or GROUP:DEVICE or ENV
  - 版本
    - revision (hash) pull from ostree or docker
  - update policies
 * 基本配置 (developer 预先规划初始值)与设定(user 调整)
 * 企业级使用 搭配 Kubernetes 需要service 与 service 的关联资料

## 自我測試

## Roadmap
### recipe design
* [ ]  recipe spec
* [ ]  configs
* [ ]  - auto load config from add-files
* [ ]  - config in recipe (default value)
* [ ]  -- need to support = :=
* [ ]  git action
* [ ]  - git commit as changelog
* [ ]  - version as git tag
* [ ]  add-files template

### yak host status monitor
* [ ] host info

### ostree backend recipes
* [ ]  deploy
* [ ]  image

### ostree backend (10/15)
* [ ] first target - atomic host
* [ ] yakd - status control
* [ ] local