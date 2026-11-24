# xiaomi-be7000-singbox

A sing-box package for the Xiaomi BE7000 router and its Cortex-A73 CPU.

The source code comes from [SagerNet/sing-box](https://github.com/SagerNet/sing-box), a universal proxy platform.
The imported source baseline is commit [`039d829`](https://github.com/SagerNet/sing-box/commit/039d829d20c7bfb963dd73113afa9f2d1f080a70) on the upstream `testing` branch.
The BE7000 package uses the full-featured static ARM64 executable from the upstream [`v1.14.4` release](https://github.com/SagerNet/sing-box/releases/tag/v1.14.4) at commit [`c005f35`](https://github.com/SagerNet/sing-box/commit/c005f35300ce76427dc285481ec5695f5b179cd2).

## Features

- Uses the official static Linux ARM64 executable targeting Armv8.0-A. It starts without a system glibc or musl loader.
- Includes the upstream feature set, including NaiveProxy outbound, Tailscale, mixed TUN, and VLESS with REALITY and uTLS. The static executable includes Cronet.

## Quick Start

Clone this repository and run the following from its root with GitHub CLI and Python 3.11 or later.

```sh
gh release download v1.14.4 --repo SagerNet/sing-box --pattern 'sing-box-1.14.4-linux-arm64-musl.tar.gz' --dir build
python3 release/package-be7000.py build/sing-box-1.14.4-linux-arm64-musl.tar.gz
```

The archive is written to `dist/be7000/sing-box-xiaomi-be7000-0.0.1.tar.gz` alongside its SHA-256 file. The archive bytes, directory names, executable and license are identical to the official `sing-box-1.14.4-linux-arm64-musl.tar.gz`. The outer filename carries the BE7000 package version `0.0.1`. The executable reports upstream version `1.14.4`.

## Configuration

See the [upstream sing-box documentation](https://sing-box.sagernet.org).

## Build and Test

The [BE7000 workflow](.github/workflows/be7000.yml) handles the `linux-arm64-musl` package. It verifies the upstream SHA-256, checks static linking and validates a sample mixed TUN, SOCKS and VLESS REALITY configuration on Ubuntu and Alpine. Successful runs provide the tarball and checksum as an Actions artifact.

## License

```
Copyright (C) 2022 by nekohasekai <contact-sagernet@sekai.icu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

In addition, no derivative work may use the name or imply association
with this application without prior consent.
```
