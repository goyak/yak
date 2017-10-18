 $ yak install core

## Ostree 特性
* Transactional upgrades and rollback for the system
* Replicating content incrementally over HTTP via GPG signatures and "pinned TLS" support
* Support for parallel installing more than just 2 bootable roots
* Binary history on the server side (and client)
* Introspectable shared library API for build and deployment systems
* Flexible support for multiple branches and repositories, supporting projects like flatpak which use libostree for applications, rather than hosts.


### 可讀寫區域

### 開機特性


> rpm-ostree
> atomic
> rpm-ostree-toolbox

```
/* Locate kernel/initramfs in the tree; the current standard is to look in
 * /usr/lib/modules/$kver/vmlinuz first.
 *
 * Originally OSTree defined kernels to be found underneath /boot
 * in the tree.  But that means when mounting /boot at runtime
 * we end up masking the content underneath, triggering a warning.
 *
 * For that reason, and also consistency with the "/usr defines the OS" model we
 * later switched to defining the in-tree kernels to be found under
 * /usr/lib/ostree-boot. But since then, Fedora at least switched to storing the
 * kernel in /usr/lib/modules, which makes sense and isn't ostree-specific, so
 * we prefer that now. However, the default Fedora layout doesn't put the
 * initramfs there, so we need to look in /usr/lib/ostree-boot first.
 */
static gboolean
get_kernel_from_tree (int                  deployment_dfd,
                      OstreeKernelLayout **out_layout,
                      GCancellable        *cancellable,
                      GError             **error)
```