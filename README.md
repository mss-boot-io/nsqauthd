# NSQ Auth Server
`nsqauthd` is a lightweight HTTP service that provides authentication and authorization for NSQ clusters. By configuring the -auth-http-address parameter in nsqd, you can delegate authentication logic to nsqauthd, enabling fine-grained control over publish and subscribe operations. This project is built on mss-boot-http and developed in Go, offering high performance and extensibility.

## 📌 Project Overview
`nsqauthd` aims to provide authorization support for NSQ's nsqd instances. By configuring the -auth-http-address parameter in nsqd, you can delegate authentication logic to nsqauthd, enabling fine-grained control over publish and subscribe operations. This project is built on mss-boot-http and developed in Go, offering high performance and extensibility.

## ⚙️ Features
- **Authentication**: Supports client identity verification via HTTP requests.
- **Authorization**: Defines client access permissions for specific topics and channels based on identity information.
- **Flexible Configuration**: Allows parameter settings through configuration files or environment variables.
- **High Performance**: Developed in Go, capable of handling high concurrency.
- **Easy Integration**: Seamlessly integrates with NSQ clusters, simplifying authorization management.

## 🚀 Quick Start
1. Clone the Project 
```shell
git clone https://github.com/mss-boot-io/nsqauthd.git
cd nsqauthd
```
2. Build the Project
```shell
make build
```
3. Start the Project
```shell
./nsqauthd --http-address=:4181
```
4. Configure NSQ
```shell
--auth-http-address=127.0.0.1:4181
```

## 📄 Configuration Details

## 🔧 Build & Development

