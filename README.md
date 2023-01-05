# os-release-files

Running `go run main.go` will update this file using the files contained in the [collection dir](./collection).

The columns in the table are:

- **COUNT**: Number of distros that contain this field
- **FIELD**: Name of this field
- **SPEC**: Is the field part of the the [os-release spec](https://www.freedesktop.org/software/systemd/man/os-release.html)?
- **PERCENT**: Percentage of distros that contain this field
- **DISTROS**: List of IDs of distros that contain this field

| COUNT |               FIELD                | SPEC | PERCENT |                                                DISTROS                                                |
|-------|------------------------------------|------|---------|-------------------------------------------------------------------------------------------------------|
|    13 | HOME_URL                           |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu |
|    13 | ID                                 |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu |
|    13 | NAME                               |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu |
|    13 | PRETTY_NAME                        |  ✓   |    100% | almalinux, amzn, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu |
|    12 | BUG_REPORT_URL                     |  ✓   |     92% | almalinux, arch, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu       |
|    12 | VERSION                            |  ✓   |     92% | almalinux, amzn, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu       |
|    12 | VERSION_ID                         |  ✓   |     92% | almalinux, amzn, centos, debian, fedora, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu       |
|    10 | ID_LIKE                            |  ✓   |     76% | almalinux, amzn, centos, linuxmint, ol, opensuse-leap, pop, rhel, rocky, ubuntu                       |
|     9 | ANSI_COLOR                         |  ✓   |     69% | almalinux, amzn, arch, centos, fedora, ol, opensuse-leap, rhel, rocky                                 |
|     8 | CPE_NAME                           |  ✓   |     61% | almalinux, amzn, centos, fedora, ol, opensuse-leap, rhel, rocky                                       |
|     7 | PLATFORM_ID                        |      |     53% | almalinux, amzn, centos, fedora, ol, rhel, rocky                                                      |
|     6 | SUPPORT_URL                        |  ✓   |     46% | arch, debian, fedora, linuxmint, pop, ubuntu                                                          |
|     5 | VERSION_CODENAME                   |  ✓   |     38% | debian, fedora, linuxmint, pop, ubuntu                                                                |
|     4 | DOCUMENTATION_URL                  |  ✓   |     30% | almalinux, arch, fedora, rhel                                                                         |
|     4 | LOGO                               |  ✓   |     30% | almalinux, arch, fedora, pop                                                                          |
|     4 | PRIVACY_POLICY_URL                 |  ✓   |     30% | fedora, linuxmint, pop, ubuntu                                                                        |
|     4 | REDHAT_SUPPORT_PRODUCT             |      |     30% | almalinux, centos, fedora, rhel                                                                       |
|     4 | REDHAT_SUPPORT_PRODUCT_VERSION     |      |     30% | almalinux, centos, fedora, rhel                                                                       |
|     3 | UBUNTU_CODENAME                    |      |     23% | linuxmint, pop, ubuntu                                                                                |
|     3 | VARIANT                            |  ✓   |     23% | fedora, ol, rhel                                                                                      |
|     3 | VARIANT_ID                         |  ✓   |     23% | fedora, ol, rhel                                                                                      |
|     2 | REDHAT_BUGZILLA_PRODUCT            |      |     15% | fedora, rhel                                                                                          |
|     2 | REDHAT_BUGZILLA_PRODUCT_VERSION    |      |     15% | fedora, rhel                                                                                          |
|     1 | ALMALINUX_MANTISBT_PROJECT         |      |      7% | almalinux                                                                                             |
|     1 | ALMALINUX_MANTISBT_PROJECT_VERSION |      |      7% | almalinux                                                                                             |
|     1 | BUILD_ID                           |  ✓   |      7% | arch                                                                                                  |
|     1 | CENTOS_MANTISBT_PROJECT            |      |      7% | centos                                                                                                |
|     1 | CENTOS_MANTISBT_PROJECT_VERSION    |      |      7% | centos                                                                                                |
|     1 | DEFAULT_HOSTNAME                   |  ✓   |      7% | fedora                                                                                                |
|     1 | ORACLE_BUGZILLA_PRODUCT            |      |      7% | ol                                                                                                    |
|     1 | ORACLE_BUGZILLA_PRODUCT_VERSION    |      |      7% | ol                                                                                                    |
|     1 | ORACLE_SUPPORT_PRODUCT             |      |      7% | ol                                                                                                    |
|     1 | ORACLE_SUPPORT_PRODUCT_VERSION     |      |      7% | ol                                                                                                    |
|     1 | ROCKY_SUPPORT_PRODUCT              |      |      7% | rocky                                                                                                 |
|     1 | ROCKY_SUPPORT_PRODUCT_VERSION      |      |      7% | rocky                                                                                                 |
|     1 | SUPPORT_END                        |  ✓   |      7% | fedora                                                                                                |
