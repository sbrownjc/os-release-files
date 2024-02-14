# os-release-files

Running `go run main.go` will update this file using the files contained in the [collection dir](./collection).

Files are named after the PRETTY_NAME variable converted to lowercase and all non alphanumeric characters converted to dashes.

i.e. in ZSH: `source $f; name=${PRETTY_NAME:l}; name=${name//[^a-zA-Z0-9]/-}; mv $f collection/$name`

The columns in the table are:

- **COUNT**: Number of distros that contain this field
- **FIELD**: Name of this field
- **SPEC**: Is the field part of the the [os-release spec](https://www.freedesktop.org/software/systemd/man/os-release.html)?
- **PERCENT**: Percentage of distros that contain this field
- **DISTROS**: List of IDs of distros that contain this field

| COUNT |               FIELD                | SPEC | PERCENT |                                                             DISTROS                                                              |
|-------|------------------------------------|------|---------|----------------------------------------------------------------------------------------------------------------------------------|
|    15 | BUG_REPORT_URL                     |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, neon, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    15 | HOME_URL                           |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, neon, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    15 | ID                                 |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, neon, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    15 | NAME                               |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, neon, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    15 | PRETTY_NAME                        |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, neon, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    14 | VERSION_ID                         |  ✓   |     93% | almalinux, amzn, centos, debian, fedora, linuxmint, neon, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu       |
|    13 | VERSION                            |  ✓   |     86% | almalinux, amzn, centos, debian, fedora, linuxmint, neon, ol, opensuse-leap, pop, rhel, rocky, ubuntu                            |
|    12 | ID_LIKE                            |  ✓   |     80% | almalinux, amzn, centos, linuxmint, neon, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu                       |
|    10 | ANSI_COLOR                         |  ✓   |     66% | almalinux, amzn, arch, centos, fedora, ol, opensuse-leap, opensuse-tumbleweed, rhel, rocky                                       |
|     9 | CPE_NAME                           |  ✓   |     60% | almalinux, amzn, centos, fedora, ol, opensuse-leap, opensuse-tumbleweed, rhel, rocky                                             |
|     7 | PLATFORM_ID                        |      |     46% | almalinux, amzn, centos, fedora, ol, rhel, rocky                                                                                 |
|     7 | SUPPORT_URL                        |  ✓   |     46% | arch, debian, fedora, linuxmint, neon, pop, ubuntu                                                                               |
|     6 | VERSION_CODENAME                   |  ✓   |     40% | debian, fedora, linuxmint, neon, pop, ubuntu                                                                                     |
|     5 | DOCUMENTATION_URL                  |  ✓   |     33% | almalinux, arch, fedora, opensuse-tumbleweed, rhel                                                                               |
|     5 | LOGO                               |  ✓   |     33% | almalinux, arch, fedora, opensuse-tumbleweed, pop                                                                                |
|     5 | PRIVACY_POLICY_URL                 |  ✓   |     33% | fedora, linuxmint, neon, pop, ubuntu                                                                                             |
|     4 | REDHAT_SUPPORT_PRODUCT             |      |     26% | almalinux, centos, fedora, rhel                                                                                                  |
|     4 | REDHAT_SUPPORT_PRODUCT_VERSION     |      |     26% | almalinux, centos, fedora, rhel                                                                                                  |
|     4 | UBUNTU_CODENAME                    |      |     26% | linuxmint, neon, pop, ubuntu                                                                                                     |
|     3 | VARIANT                            |  ✓   |     20% | fedora, ol, rhel                                                                                                                 |
|     3 | VARIANT_ID                         |  ✓   |     20% | fedora, ol, rhel                                                                                                                 |
|     2 | REDHAT_BUGZILLA_PRODUCT            |      |     13% | fedora, rhel                                                                                                                     |
|     2 | REDHAT_BUGZILLA_PRODUCT_VERSION    |      |     13% | fedora, rhel                                                                                                                     |
|     2 | SUPPORT_END                        |  ✓   |     13% | amzn, fedora                                                                                                                     |
|     1 | ALMALINUX_MANTISBT_PROJECT         |      |      6% | almalinux                                                                                                                        |
|     1 | ALMALINUX_MANTISBT_PROJECT_VERSION |      |      6% | almalinux                                                                                                                        |
|     1 | BUILD_ID                           |  ✓   |      6% | arch                                                                                                                             |
|     1 | CENTOS_MANTISBT_PROJECT            |      |      6% | centos                                                                                                                           |
|     1 | CENTOS_MANTISBT_PROJECT_VERSION    |      |      6% | centos                                                                                                                           |
|     1 | DEFAULT_HOSTNAME                   |  ✓   |      6% | fedora                                                                                                                           |
|     1 | ORACLE_BUGZILLA_PRODUCT            |      |      6% | ol                                                                                                                               |
|     1 | ORACLE_BUGZILLA_PRODUCT_VERSION    |      |      6% | ol                                                                                                                               |
|     1 | ORACLE_SUPPORT_PRODUCT             |      |      6% | ol                                                                                                                               |
|     1 | ORACLE_SUPPORT_PRODUCT_VERSION     |      |      6% | ol                                                                                                                               |
|     1 | ROCKY_SUPPORT_PRODUCT              |      |      6% | rocky                                                                                                                            |
|     1 | ROCKY_SUPPORT_PRODUCT_VERSION      |      |      6% | rocky                                                                                                                            |
