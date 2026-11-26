# xiaomi-be7000-singbox

A sing-box executable for the Xiaomi BE7000 router and its Cortex-A73 CPU.

The source code comes from [SagerNet/sing-box](https://github.com/SagerNet/sing-box), a universal proxy platform.
The BE7000 package builds this fork from upstream `testing` commit [`491b845be`](https://github.com/SagerNet/sing-box/commit/491b845be7ac4bdaa3e3f9443505eb57e061a28c), a development snapshot targeting `1.15.0-alpha.14`.

## Features

- Builds a static Linux ARM64 executable targeting Armv8.0-A. It starts without a system glibc or musl loader.
- Includes NaiveProxy outbound with Cronet, Tailscale, the sing-tun TCP/IP stack, and VLESS with REALITY and uTLS.
- Tries the Go TUN TCP fast path after early handshakes and any configured TCP prefix have been sent.

## Quick Start

Download the executable and `SHA256SUMS.txt` from a successful [BE7000 build](https://github.com/lyleee0083/xiaomi-be7000-singbox/actions/workflows/be7000.yml) or its [GitHub Release](https://github.com/lyleee0083/xiaomi-be7000-singbox/releases). Check the digest and grant execute permission before running it.

```sh
sha256sum -c SHA256SUMS.txt
chmod +x sing-box-xiaomi-be7000-0.0.3
./sing-box-xiaomi-be7000-0.0.3 version
```

## Configuration

Omit the TUN `stack` field to use the sing-tun TCP/IP stack and its TCP fast path. See the [upstream sing-box documentation](https://sing-box.sagernet.org).

## Build and Test

The [BE7000 workflow](.github/workflows/be7000.yml) prepares the pinned Chromium musl toolchain and builds the fork checkout with [the build script](release/build-be7000.sh). It verifies the executable filename, checksum, version, source commit and static linking. It also validates a sample TUN, SOCKS and VLESS REALITY configuration on Ubuntu ARM64 and Alpine. Successful runs provide the executable and `SHA256SUMS.txt` as direct files.

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
