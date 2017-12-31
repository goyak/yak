# 相似solutions

cloud-init

## package based

### package manager - yum and apt
spec
```
%description
administrate and deploy your gandi resources

%prep
%setup -n gandi.cli-%{version}

%build
%{__python2} setup.py build

%install
%{__python2} setup.py install --skip-build --root $RPM_BUILD_ROOT
rst2man --no-generator gandicli.man.rst | gzip -9 > gandi.1.gz
install -d -m 0755 %{buildroot}/usr/share/man/man1
install -m 0644 gandi.1.gz %{buildroot}/usr/share/man/man1

%clean
rm -r %{buildroot}
```

### snap

```
channels:                        
  - latest:                           
      stable:    3.1.5      (46) 7MB  -
      candidate: ^                    
      beta:      ^                    
      edge:      3.1+master (50) 10MB -
  - 2.3:                              
      stable:    2.3.8      (22) 5MB  -
      candidate: 2.3.8      (39) 5MB  -
      beta:      ^                    
      edge:      ^                    
  - 3.1:                              
      stable:    3.1.5      (46) 7MB  -
      candidate: ^                    
      beta:      ^                    
      edge:      3.1+master (50) 10MB -
```

### atomicapp
![AtomicApp](https://raw.githubusercontent.com/projectatomic/atomicapp/master/docs/images/logo.png)
### flatpak
 - OSTree uses hardlinks, but has otherwise very little requirements on the underlying filesystem.
 - ![Filesystem](https://github.com/flatpak/flatpak/wiki/Filesystem)

## whole distro
### atomic # atomic 基本有提供 atomic install 的command 安装 docker services 至系统
### meta-updater
 - ostree for Yocto
 - solutions focus on different hardware device (arm)
 - [FOSDEM17_OSTree](https://archive.fosdem.org/2017/schedule/event/updates_with_ostree/attachments/slides/1663/export/events/attachments/updates_with_ostree/slides/1663/FOSDEM_OSTree.pdf)

## image
### projectatomic/skopeo
https://github.com/projectatomic/skopeo

container image

 * docker registries
 * the Atomic registry
 * private registries
 * local directories

## Orchestration
### rancher
 - similar with juju
 - ranchar yml
 - 著重于 service 与 service 之间关联
 - https://github.com/rancher/agent > go, 提供host 基本资讯
 - ![rancher demo](https://player.vimeo.com/video/212646077)

### docker-compose
### kubernetes
 - k8s.io/client-go
### juju
![juju](https://jujucharms.com/static/img/home/laptop.png)
## other
### ansible - Simple, agentless and powerful open source IT automation.
 - ansible-exec
 - [ansible playbook yml](http://docs.ansible.com/ansible/latest/playbooks.html)
  - loop design
 - Facts.d
 - [tower demo](https://www.ansible.com/tower-demo)

### landscape
* Automate day-to-day tasks
* Receive alerts to update machines you specify
* Keep secure with the latest security patches
* Manage up to 40,000 machines with a single interface
* Create your own software repositories
* Extend and customise Landscape via our API

