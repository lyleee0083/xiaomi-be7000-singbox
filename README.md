# xiaomi-be7000-singbox

A sing-box executable for the Xiaomi BE7000 router and its Cortex-A73 CPU.

The source code comes from [SagerNet/sing-box](https://github.com/SagerNet/sing-box), a universal proxy platform.
The imported source baseline is commit [`039d829`](https://github.com/SagerNet/sing-box/commit/039d829d20c7bfb963dd73113afa9f2d1f080a70) on the upstream `testing` branch.
The BE7000 package builds the full-featured static ARM64 executable from the upstream [`v1.14.4` release](https://github.com/SagerNet/sing-box/releases/tag/v1.14.4) at commit [`c005f35`](https://github.com/SagerNet/sing-box/commit/c005f35300ce76427dc285481ec5695f5b179cd2).

## Features

- Builds a static Linux ARM64 executable targeting Armv8.0-A. It starts without a system glibc or musl loader.
- Includes the upstream feature set, including NaiveProxy outbound, Tailscale, mixed TUN, and VLESS with REALITY and uTLS. The static executable includes Cronet.

## Quick Start

Download the executable and `SHA256SUMS.txt` from [GitHub Releases](https://github.com/lyleee0083/xiaomi-be7000-singbox/releases/latest). Check the digest and grant execute permission before running it.

```sh
sha256sum -c SHA256SUMS.txt
chmod +x sing-box-xiaomi-be7000-0.0.2
./sing-box-xiaomi-be7000-0.0.2 version
```

## Configuration

See the [upstream sing-box documentation](https://sing-box.sagernet.org).

## Build and Test

The [BE7000 workflow](.github/workflows/be7000.yml) prepares the pinned upstream Chromium musl toolchain and invokes [the build script](release/build-be7000.sh). It verifies the executable filename, checksum, version and static linking. It also validates a sample mixed TUN, SOCKS and VLESS REALITY configuration on Ubuntu ARM64 and Alpine. Successful runs provide the executable and `SHA256SUMS.txt` as direct files.

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
