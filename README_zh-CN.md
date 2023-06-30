<h1 align="center" style="border-bottom: none">
    <b>
        <a href="https://docker.nsddd.top">OpenKF</a><br>
    </b>
</h1>
<h3 align="center" style="border-bottom: none">
      ⭐️  OpenKF（开放知识流）是一个在线客服系统。 ⭐️ <br>
<h3>


<p align=center>
<a href="https://goreportcard.com/report/github.com/OpenIMSDK/OpenKF"><img src="https://goreportcard.com/badge/github.com/OpenIMSDK/OpenKF" alt="A+"></a>
<a href="https://github.com/OpenIMSDK/OpenKF/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22good+first+issue%22"><img src="https://img.shields.io/github/issues/OpenIMSDK/OpenKF/good%20first%20issue?logo=%22github%22" alt="good first"></a>
<a href="https://github.com/OpenIMSDK/OpenKF"><img src="https://img.shields.io/github/stars/OpenIMSDK/OpenKF.svg?style=flat&logo=github&colorB=deeppink&label=stars"></a>
<a href="https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg"><img src="https://img.shields.io/badge/Slack-100%2B-blueviolet?logo=slack&amp;logoColor=white"></a>
<a href="https://github.com/OpenIMSDK/OpenKF/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-green"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/badge/Language-Go-blue.svg"></a>
</p>

</p>

<p align="center">
    <a href="./README.md"><b>English</b></a> •
    <a href="./README_zh-CN.md"><b>中文</b></a>
</p>

</p>

----

## 🧩 特性

1. [OpenKF](https://github.com/OpenIMSDK/OpenKF) 是基于 [OpenIM](https://github.com/OpenIMSDK) 的开源客服系统。
2. 支持 LLM（本地知识库）客服。
3. 支持多渠道客服，并易于与第三方系统集成。
4. 易于部署和二次开发。

## 🛫 快速开始

> **注意**：您可以通过以下方式快速开始使用 OpenKF。

### 📦 安装

```bash
git clone https://https://github.com/OpenIMSDK/OpenKF
```

### 🚀 运行

```bash
```

## 🕋 架构图

![架构图](assets/images/architecture.png)

**MVC 架构设计：**

![MVC](assets/images/mvc.png)

## 🤖 文件目录说明

目录规范化设计结构：

```bash
.
├── assets
│   └── images
├── build
├── deploy
├── docs
├── kf_plugins # 带有 LLM 的本地知识库
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

我们希望任何人都能参与到我们的社区中来，我们提供礼品和奖励，并欢迎您在每个星期四的晚上加入我们。

我们在 [GitHub 讨论区](https://github.com/orgs/OpenIMSDK/discussions/categories/meeting) 中记录了每一次[双周会议](https://github.com/orgs/OpenIMSDK/discussions/categories/meeting)，我们的会议纪要写在 [Google 文档](https://docs.google.com/document/d/1nx8MDpuG74NASx081JcCpxPgDITNTpIIos0DS6Vr9GU/edit?usp=sharing) 中。

## 🤼‍ 贡献与开发

OpenIMSDK 的目标是建立一个顶级的开源社区。我们有一套标准，可以在 [Community 仓库](https://github.com/OpenIMSDK/community) 中找到。

如果您想对此 OpenKF 仓库进行贡献，请阅读我们的 [贡献者文档](https://github.com/OpenIMSDK/OpenKF/blob/main/CONTRIBUTING.md)。

在开始之前，请确保您的更改是需要的。最好的方式是创建一个 [新的讨论](https://github.com/OpenIMSDK/OpenKF/discussions/new/choose) 或使用 [Slack 交流](https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg)，或者如果您发现了问题，请先进行 [报告](https://github.com/OpenIMSDK/OpenKF/issues/new/choose)。

## 🚨 许可证

OpenIMSDK 在 Apache 2.0 许可证下发布。完整的许可证文本请参阅 [LICENSE](https://github.com/OpenIMSDK/OpenKF/tree/main/LICENSE)。

## 🔮 感谢我们的贡献者！

<a href="https://github.com/OpenIMSDK/OpenKF/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=OpenIMSDK/OpenKF" />
</a>
