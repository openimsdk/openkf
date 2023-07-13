<p align="center"> <img src="assets/logo-gif/openkf-logo-white.gif" width="60%" height="30%"/> </p> <h3 align="center" style="border-bottom: none"> ⭐️  OpenKF（Open Knowledge Flow）是一款在线智能客服系统。⭐️ <br> <h3> <p align=center> <a href="https://goreportcard.com/report/github.com/OpenIMSDK/OpenKF"><img src="https://goreportcard.com/badge/github.com/OpenIMSDK/OpenKF" alt="A+"></a> <a href="https://github.com/OpenIMSDK/OpenKF/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A"good+first+issue""><img src="https://img.shields.io/github/issues/OpenIMSDK/OpenKF/good first issue?logo="github"" alt="good first"></a> <a href="https://github.com/OpenIMSDK/OpenKF"><img src="https://img.shields.io/github/stars/OpenIMSDK/OpenKF.svg?style=flat&logo=github&colorB=deeppink&label=stars"></a> <a href="https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg"><img src="https://img.shields.io/badge/Slack-100%2B-blueviolet?logo=slack&amp;logoColor=white"></a> <a href="https://github.com/OpenIMSDK/OpenKF/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-green"></a> <a href="https://golang.org/"><img src="https://img.shields.io/badge/Language-Go-blue.svg"></a> </p>

</p>

<p align="center"> <a href="./README.md"><b>English</b></a> • <a href="./README_zh-CN.md"><b>中文</b></a> </p>

</p> <br>

## 🧩 强大的功能

1. [OpenKF](https://github.com/OpenIMSDK/OpenKF) 是一款基于 [OpenIM](https://github.com/OpenIMSDK) 的开源客服系统。
2. 支持 LLM（本地知识库）客服。
3. 支持多渠道客服，并能与第三方系统轻松集成。
4. 部署简单，便于二次开发。

## 🛫 快速开始

> **注意**：您可以通过以下步骤快速开始使用 OpenKF。

### 📦 安装

```bashCopy Code
git clone https://github.com/OpenIMSDK/OpenKF openkf && export openkf=$(pwd)/openkf && cd $openkf && make
```

### 🚀 运行

> **注意**：首先需要运行后端服务器

```
make build
```

> 打开另一个终端窗口，运行以下命令:

```
make dev
cd web
npm run dev
```

### 📖 贡献者入门指南

熟练使用 Makefile 可以确保项目的质量。

```
用法: make <TARGETS> ...

目标:
  all                          构建所有必要的目标。🏗️
  build                        默认情况下构建二进制文件。🛠️
  go.build                     构建指定平台的二进制文件。👨‍💻
  build-multiarch              为多个平台构建二进制文件。🌍
  tidy                         整理 go.mod 📦
  style                        代码样式 -> fmt,vet,lint 🎨
  fmt                          运行 go fmt 格式化代码。✨
  vet                          运行 go vet 静态检查代码。🔍
  generate                     运行 go generate 自动生成代码和文档。✅
  lint                         运行 go lint 静态分析代码。🔎
  test                         运行单元测试 ✔️
  cover                        运行带覆盖率的单元测试。🧪
  docker-build                 使用管理器构建 Docker 镜像。🐳
  docker-push                  使用管理器推送 Docker 镜像。🔝
  docker-buildx-push           使用管理器和 buildx 推送 Docker 镜像。🚢
  copyright-verify             验证文件的版权声明。📄
  copyright-add                为所有文件添加版权声明。📝
  swagger                      生成 Swagger 文档。📚
  serve-swagger                提供 Swagger 规范和文档。🌐
  clean                        清除所有构建结果。🧹
  help                         显示此帮助信息。ℹ️
```

> **注意**：强烈建议在提交代码之前运行 `make all`。🚀

```
make all
```

## 🕋 架构图

![架构](./assets/images/architecture.png)

**MVC 架构设计:**

![MVC](./assets/images/mvc.png)

## 🤖 文件目录说明

目录结构规范化设计:

```
.
├── assets
│   └── images
├── build
├── deploy
├── docs
├── kf_plugins # 含有 LLM 的本地知识库
│   ├── chat
│   ├── config
│   ├── data
│   ├── logs
│   ├── model
│   └── utils
├── scripts
│   ├── githooks
│   └── LICENSE
├── server # OpenKF 后端
│   ├── cmd
│   ├── data
│   ├── docs
│   ├── examples
│   ├── internal
│   ├── logs
│   ├── pkg
│   ├── test
│   └── tools
└── web # OpenKF 前端
    ├── public
    ├── scripts
    └── src
```

## 🗓️ 社区会议

我们欢迎任何人参与我们的社区，我们会提供礼物和奖励，并欢迎您加入我们的每周四晚上的活动。

我们的会议在 [OpenIM Slack](https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg) 🎯 的 `openkf` 频道中，您可以搜索 openkf 频道加入。

我们会将每次 [双周会议](https://github.com/orgs/OpenIMSDK/discussions/categories/meeting) 的记录保存在 [GitHub discussions](https://github.com/OpenIMSDK/OpenKF/discussions/categories/meeting) 中，会议纪要会写在 [Google Docs](https://docs.google.com/document/d/1nx8MDpuG74NASx081JcCpxPgDITNTpIIos0DS6Vr9GU/edit?usp=sharing) 中。

## 🤼‍ 贡献与开发

OpenIMSDK 的目标是打造一个顶级的开源社区。我们有一套标准，在 [Community repository](https://github.com/OpenIMSDK/community) 中有详细说明。

如果您想为 OpenKF 贡献代码，请阅读我们的 [贡献者文档](https://github.com/OpenIMSDK/OpenKF/blob/main/CONTRIBUTING.md)。

在开始之前，请确保您的更改是需要的。最好先创建一个 [新的讨论](https://github.com/OpenIMSDK/OpenKF/discussions/new/choose) 或通过 [Slack 进行沟通](https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg)，或者如果您发现问题，请首先 [报告问题](https://github.com/OpenIMSDK/OpenKF/issues/new/choose)。

## 🚨 许可证

OpenIMSDK 使用 Apache 2.0 许可证。完整的许可证文本请参见 [LICENSE](https://github.com/OpenIMSDK/OpenKF/tree/main/LICENSE)。

此存储库中显示的 OpenKF 标志（包括其变体和动画版本），存储在 [OpenKF](https://github.com/OpenIMSDK/openkf) 的 [assets/logo](./assets/logo) 和 [assets/logo-gif](./assets/logo-gif) 目录下，受版权法保护。

## 🔮 感谢我们的贡献者！

<a href="https://github.com/OpenIMSDK/OpenKF/graphs/contributors"> <img src="https://contrib.rocks/image?repo=OpenIMSDK/OpenKF" /> </a>
