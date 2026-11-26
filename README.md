# xiaomi-be7000-singbox

面向小米 BE7000 路由器及其 Cortex-A73 CPU 的 sing-box 可执行文件。

源码来自通用代理平台 [SagerNet/sing-box](https://github.com/SagerNet/sing-box)。BE7000 发行包使用本仓库源码构建，基于上游 `testing` 分支的提交 [`a93b9e8a`](https://github.com/SagerNet/sing-box/commit/a93b9e8a8474253dfcab8d76aab647ea98cfcacc)，对应 `1.15.0-alpha.14` 的开发快照。

## 功能

- 构建面向 Armv8.0-A 的 Linux ARM64 静态可执行文件，启动时无需系统提供 glibc 或 musl 加载器。
- 包含内置 Cronet 的 NaiveProxy 出站、Tailscale、sing-tun TCP/IP 栈，以及支持 REALITY 和 uTLS 的 VLESS。
- 早期握手及已配置的 TCP 前缀发送完成后，尝试进入 Go TUN 的 TCP 快速路径。
- UDP 快速路径按需检查代理协议封装。支持的直连和原生 SOCKS5 UDP 继续使用批量收发。

## 快速开始

从成功的 [BE7000 构建](https://github.com/lyleee0083/xiaomi-be7000-singbox/actions/workflows/be7000.yml)或对应的 [GitHub Release](https://github.com/lyleee0083/xiaomi-be7000-singbox/releases) 下载可执行文件和 `SHA256SUMS.txt`，核对校验值并赋予执行权限。

```sh
sha256sum -c SHA256SUMS.txt
chmod +x sing-box-xiaomi-be7000-0.0.4
./sing-box-xiaomi-be7000-0.0.4 version
```

## 配置

省略 TUN 的 `stack` 字段，即可使用 sing-tun TCP/IP 栈及其 TCP 快速路径。旧配置若使用 `mixed` 或 `gvisor`，请在升级前移除 `stack` 字段。其他配置方法见 [sing-box 上游文档](https://sing-box.sagernet.org)。

## 构建与测试

[BE7000 工作流](.github/workflows/be7000.yml) 准备固定版本的 Chromium musl 工具链，并通过[构建脚本](release/build-be7000.sh)编译本仓库源码。工作流核验文件名、校验值、版本、源码提交和静态链接，并在 Ubuntu ARM64 与 Alpine 上检查包含 TUN、SOCKS 和 VLESS REALITY 的示例配置。

## 版本发布

成功的构建直接提供可执行文件和 `SHA256SUMS.txt`。

## 许可证

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
