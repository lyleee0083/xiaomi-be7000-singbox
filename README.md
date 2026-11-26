# xiaomi-be7000-singbox

面向小米 BE7000 路由器及其 Cortex-A73 CPU 的 sing-box 可执行文件。

源码来自通用代理平台 [SagerNet/sing-box](https://github.com/SagerNet/sing-box)。BE7000 发行包使用本仓库源码构建，基于上游 `testing` 分支的提交 [`a93b9e8a`](https://github.com/SagerNet/sing-box/commit/a93b9e8a8474253dfcab8d76aab647ea98cfcacc)，对应 `1.15.0-alpha.14` 的开发快照。

## 功能

- 构建面向 Armv8.0-A 的 Linux ARM64 静态可执行文件，启动时无需系统提供 glibc 或 musl 加载器。
- 使用 sing-tun TCP/IP 栈，支持 VLESS REALITY、uTLS、Vision TCP 和 XUDP。
- 早期握手及已配置的 TCP 前缀发送完成后，尝试进入 Go TUN 的 TCP 快速路径。
- Go TUN 直连 UDP 满足条件时使用批量收发。VLESS UDP 直接进入 XUDP 转发，跳过原生 UDP 套接字的快速路径检查。

## 快速开始

从成功的 [BE7000 构建](https://github.com/lyleee0083/xiaomi-be7000-singbox/actions/workflows/be7000.yml)或对应的 [GitHub Release](https://github.com/lyleee0083/xiaomi-be7000-singbox/releases) 下载可执行文件和 `SHA256SUMS.txt`，核对校验值并赋予执行权限。

```sh
sha256sum -c SHA256SUMS.txt
chmod +x sing-box-xiaomi-be7000-0.0.5
./sing-box-xiaomi-be7000-0.0.5 version
```

## 配置

本发行版面向本地 BE7000 OpenWrt 规则，支持以下配置类型。

| 用途 | 类型 |
| --- | --- |
| 入站 | `tun`、`socks`、`direct` |
| 出站 | `direct`、`selector`、`vless` |
| DNS | `udp`、`https`、`hosts`、`local` |

省略 TUN 的 `stack` 字段，使用 sing-tun TCP/IP 栈。仍使用 `mixed` 的 OpenWrt 打包配置需先适配对应的 TUN 和防火墙集成。VLESS 通过 TCP 连接代理服务器，启用 REALITY、uTLS 与 Vision 时沿用对应握手。UDP 使用 XUDP 在 VLESS 连接内转发，`packet_encoding` 可省略或设为 `xudp`。

## 构建与测试

[BE7000 工作流](.github/workflows/be7000.yml) 在 GitHub Actions ARM64 runner 上用 musl 编译[静态可执行文件](release/build-be7000.sh)。工作流检查模块依赖、运行 Go 测试，核验文件校验值、版本、源码提交和静态链接，并在 ARM64 Ubuntu 与 Alpine 上检查不含真实节点凭据的示例配置。

## 仓库结构

- `include/` 注册 BE7000 支持的入站、出站和 DNS 类型。
- `protocol/`、`dns/` 与 `route/` 包含对应的转发、解析和路由实现。
- [release/be7000.json](release/be7000.json) 记录发行版本和上游提交。

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
