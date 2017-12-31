# Application recipes (or spec, playbook)

ref: snapd/snap/info_snap_yaml.go
gopkg.in/yaml.v2
https://github.com/go-yaml/yaml/tree/v2

Recipes 儲存方式以git, version 作為 git tag, commit log 作為 release note

![Recipes](https://g.gravizo.com/svg?
  @startuml;
  ;
  frame "backend" {;
    [plain];
    [ostree];
    [docker];
    [k8s];
  };
  [plain] ..> [recipes];
  [ostree] ..> [recipes];
  [docker] ..> [recipes];
  [k8s] ..> [recipes];
  [recipes] -> recipe;
  frame "files" {;
    package "hooks" {;
      [do_bootloader];
      [configure];
      [preinst_hook];
      [install];
      [postinst_hook];
      [prerm_hook];
      [remove];
      [postrm_hook];
    };
    package "added_files" {;
      [file_a];
      [file_b];
      [file_c];
    };
    frame backend_specific {;
      package "backend_specific_files" {;
        [abc];
        [foo];
      };
      package "backend_specific_hooks" {;
        [f_abc];
        [f_foo];
      };
    };
  };
  ; 
  frame config.yml {;
    package "common_config" {;
      [COMMON_ABC];
      [COMMON_FOO];
    };
    package "config_table" {;
      [ABC];
      [FOO];
    };
  };
  frame yak.yml {;
    [ver];
    [desc];
  };
  files .> config_table;
  common_config -> config_table;
  ;
  backend_specific_hooks .> hooks;
  backend_specific_files .> added_files;
  recipe -> yak.yml;
  recipe -> config.yml;
  recipe -> files;
  ;
  @enduml;
)

## file structure
```
 ├─ yak.yml
 ├─ config.yml
 └─┬─ (d) hooks
   └─ (d) added-files
```

### tag log
```
 $ git log --no-walk --tags --pretty="%h %d %s" --decorate=full
3713f3f  (tag: refs/tags/1.0.0, tag: refs/tags/0.6.0, refs/remotes/origin/master, refs/heads/master) SP-144/ISP-177: Updating the package.json with 0.6.0 version and the README.md.
00a3762  (tag: refs/tags/0.5.0) ISP-144/ISP-205: Update logger to save files with optional port number if defined/passed: Version 0.5.0
d8db998  (tag: refs/tags/0.4.2) ISP-141/ISP-184/ISP-187: Fixing the bug when loading the app with Gulp and Grunt for 0.4.2
3652484  (tag: refs/tags/0.4.1) ISP-141/ISP-184: Missing the package.json and README.md updates with the 0.4.1 version
c55eee7  (tag: refs/tags/0.4.0) ISP-141/ISP-184/ISP-187: Updating the README.md file with the latest 1.3.0 version.
6963d0b  (tag: refs/tags/0.3.0) ISP-141/ISP-184: Add support for custom serializers: README update
4afdbbe  (tag: refs/tags/0.2.0) ISP-141/ISP-143/ISP-144: Fixing a bug with the creation of the logs
e1513f1  (tag: refs/tags/0.1.0) ISP-141/ISP-143: Betterr refactoring of the Loggers, no dependencies, self-configuration for missing settings.
```

## yak.yml
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
## backend
不同的backend在針對每個動作會有不同的
 - plain: upgrade yak itself
 - ostree: atomic, systemd plugin
 - docker: wrapper, docker-run plugin
 - kubernetes: pod, k8s plugin
- plugin:
  - host status check

### backend_specific_files or application_specific_files

### ostree
 * add-files
 * treecompose-post.sh
 * *.{repo}
 * image
  * config.ini
  * xx.tdl
  * es-atomic-host-7.ks

### ostree single deployment
use follow command can clean up rollback deployments, but if we need to rollback, we
need to use the recipes function.  
  ostree cleanup -r 

默认rpm-ostree如果只有一个deploy时， rpm-ostree rollback 将失败。  
所以应该考虑同过proxy来直接访问到ostree的code。 C   


## hooks
   - do_bootloader
   - configure
   - preinst_hook
   - install
   - postinst_hook
   - prerm_hook
   - remove
   - postrm_hook

## recipe verification
