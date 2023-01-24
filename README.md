# os-release-files

Running `go run main.go` will update this file using the files contained in the [collection dir](./collection).

Files are named after the PRETTY_NAME variable converted to lowercase and all non alphanumeric characters converted to dashes.

i.e. in ZSH: `source $f; name=${PRETTY_NAME:l}; name=${name//[^a-zA-Z0-9]/-}; mv $f $name`

The columns in the table are:

- **COUNT**: Number of distros that contain this field
- **FIELD**: Name of this field
- **SPEC**: Is the field part of the the [os-release spec](https://www.freedesktop.org/software/systemd/man/os-release.html)?
- **PERCENT**: Percentage of distros that contain this field
- **DISTROS**: List of IDs of distros that contain this field

| COUNT |               FIELD                | SPEC | PERCENT |                                                          DISTROS                                                           |
|-------|------------------------------------|------|---------|----------------------------------------------------------------------------------------------------------------------------|
|    14 | HOME_URL                           |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    14 | ID                                 |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    14 | NAME                               |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    14 | PRETTY_NAME                        |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu |
|    13 | BUG_REPORT_URL                     |  ✓   |     92% | almalinux, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu       |
|    13 | VERSION_ID                         |  ✓   |     92% | almalinux, amzn, centos, debian, fedora, linuxmint, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu       |
|    12 | VERSION                            |  ✓   |     85% | almalinux, amzn, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu                            |
|    11 | ID_LIKE                            |  ✓   |     78% | almalinux, amzn, centos, linuxmint, ol, opensuse-leap, opensuse-tumbleweed, pop, rhel, rocky, ubuntu                       |
|    10 | ANSI_COLOR                         |  ✓   |     71% | almalinux, amzn, arch, centos, fedora, ol, opensuse-leap, opensuse-tumbleweed, rhel, rocky                                 |
|     9 | CPE_NAME                           |  ✓   |     64% | almalinux, amzn, centos, fedora, ol, opensuse-leap, opensuse-tumbleweed, rhel, rocky                                       |
|     7 | PLATFORM_ID                        |      |     50% | almalinux, amzn, centos, fedora, ol, rhel, rocky                                                                           |
|     6 | SUPPORT_URL                        |  ✓   |     42% | arch, debian, fedora, linuxmint, pop, ubuntu                                                                               |
|     5 | DOCUMENTATION_URL                  |  ✓   |     35% | almalinux, arch, fedora, opensuse-tumbleweed, rhel                                                                         |
|     5 | LOGO                               |  ✓   |     35% | almalinux, arch, fedora, opensuse-tumbleweed, pop                                                                          |
|     5 | VERSION_CODENAME                   |  ✓   |     35% | debian, fedora, linuxmint, pop, ubuntu                                                                                     |
|     4 | PRIVACY_POLICY_URL                 |  ✓   |     28% | fedora, linuxmint, pop, ubuntu                                                                                             |
|     4 | REDHAT_SUPPORT_PRODUCT             |      |     28% | almalinux, centos, fedora, rhel                                                                                            |
|     4 | REDHAT_SUPPORT_PRODUCT_VERSION     |      |     28% | almalinux, centos, fedora, rhel                                                                                            |
|     3 | UBUNTU_CODENAME                    |      |     21% | linuxmint, pop, ubuntu                                                                                                     |
|     3 | VARIANT                            |  ✓   |     21% | fedora, ol, rhel                                                                                                           |
|     3 | VARIANT_ID                         |  ✓   |     21% | fedora, ol, rhel                                                                                                           |
|     2 | REDHAT_BUGZILLA_PRODUCT            |      |     14% | fedora, rhel                                                                                                               |
|     2 | REDHAT_BUGZILLA_PRODUCT_VERSION    |      |     14% | fedora, rhel                                                                                                               |
|     1 | ALMALINUX_MANTISBT_PROJECT         |      |      7% | almalinux                                                                                                                  |
|     1 | ALMALINUX_MANTISBT_PROJECT_VERSION |      |      7% | almalinux                                                                                                                  |
|     1 | BUILD_ID                           |  ✓   |      7% | arch                                                                                                                       |
|     1 | CENTOS_MANTISBT_PROJECT            |      |      7% | centos                                                                                                                     |
|     1 | CENTOS_MANTISBT_PROJECT_VERSION    |      |      7% | centos                                                                                                                     |
|     1 | DEFAULT_HOSTNAME                   |  ✓   |      7% | fedora                                                                                                                     |
|     1 | ORACLE_BUGZILLA_PRODUCT            |      |      7% | ol                                                                                                                         |
|     1 | ORACLE_BUGZILLA_PRODUCT_VERSION    |      |      7% | ol                                                                                                                         |
|     1 | ORACLE_SUPPORT_PRODUCT             |      |      7% | ol                                                                                                                         |
|     1 | ORACLE_SUPPORT_PRODUCT_VERSION     |      |      7% | ol                                                                                                                         |
|     1 | ROCKY_SUPPORT_PRODUCT              |      |      7% | rocky                                                                                                                      |
|     1 | ROCKY_SUPPORT_PRODUCT_VERSION      |      |      7% | rocky                                                                                                                      |
|     1 | SUPPORT_END                        |  ✓   |      7% | fedora                                                                                                                     |
