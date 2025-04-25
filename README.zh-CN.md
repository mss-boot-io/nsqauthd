# NSQ 授权服务器
`nsqauthd`是一个轻量级的 HTTP 服务，用于为 NSQ 集群提供身份验证和权限控制功能。通过配置 nsqd 的 -auth-http-address 参数，您可以将身份验证逻辑委托给 nsqauthd，从而实现对发布和订阅操作的精细控制。该项目基于 mss-boot-http 构建，采用 Go 语言开发，具备高性能和易于扩展的特点。

## 📌 项目简介
`nsqauthd`旨在为 NSQ 的 nsqd 实例提供授权支持。通过配置 nsqd 的 -auth-http-address 参数，您可以将身份验证逻辑委托给 nsqauthd，从而实现对发布和订阅操作的精细控制。该项目基于 mss-boot-http 构建，采用 Go 语言开发，具备高性能和易于扩展的特点。

## ⚙️ 功能特性
- **身份验证 / Authentication**：支持通过 HTTP 请求验证客户端身份。 / Supports client identity verification via HTTP requests.
- **权限控制 / Authorization**：基于身份信息，定义客户端对特定主题和频道的访问权限。 / Defines client access permissions for specific topics and channels based on identity information.
- **灵活配置 / Flexible Configuration**：支持通过配置文件或环境变量进行参数设置。 / Allows parameter settings through configuration files or environment variables.
- **高性能 / High Performance**：采用 Go 语言开发，具备高并发处理能力。 / Developed in Go, capable of handling high concurrency.
- **易于集成 / Easy Integration**：与 NSQ 集群无缝对接，简化授权管理。 / Seamlessly integrates with NSQ clusters, simplifying authorization management.

## 🚀 快速开始
1. 克隆项目
```shell
git clone https://github.com/mss-boot-io/nsqauthd.git
cd nsqauthd
```
2. 构建项目
```shell
make build
```
3. 启动服务
```shell
./nsqauthd --http-address=:4181
```
4. 配置 NSQ
```shell
--auth-http-address=127.0.0.1:4181
```

## 📄 配置说明

## 🔧 构建与开发