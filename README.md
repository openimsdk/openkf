<h1 align="center" style="border-bottom: none">
    <b>
        <a href="https://docker.nsddd.top">OpenKF</a><br>
    </b>
</h1>
<h3 align="center" style="border-bottom: none">
      ⭐️  OpenKF(Open Knowledge Flow) is an online intelligent customer service system. ⭐️ <br>
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

## 🧩 Awesome features

1. [OpenKF](https://github.com/OpenIMSDK/OpenKF) is an opensource customer service system based on [OpenIM](https://github.com/OpenIMSDK).
2. Support LLM(Local Knowledgebase) customer service.
3. Support multi-channel customer service, and easy to integrate with third-party systems.
4. Easy to deploy and secondary development.

## 🛫 Quick start 

> **Note**: You can get started quickly with OpenKF.

### 📦 Installation

```bash
git clone https://https://github.com/OpenIMSDK/OpenKF
```

### 🚀 Run

```bash
go run main.go
```
```bash
npm run dev
```

## 🕋 Architecture diagram

![Architecture](assets/images/architecture.png)

**MVC Architecture Design:**

![MVC](assets/images/mvc.png)

## 🤖 File Directory Description

Catalog standardization design structure:

```bash
.
├── assets
│   └── images
├── build
├── deploy
├── docs
├── kf_plugins # Local knowledgebase with LLM
│   ├── chat
│   ├── config
│   ├── data
│   ├── logs
│   ├── model
│   └── utils
├── scripts
│   ├── githooks
│   └── LICENSE
├── server # OpenKF backend
│   ├── cmd
│   ├── data
│   ├── docs
│   ├── examples
│   ├── internal
│   ├── logs
│   ├── pkg
│   ├── test
│   └── tools
└── web # OpenKF frontend
    ├── public
    ├── scripts
    └── src
```

## 🗓️ community meeting

We want anyone to get involved in our community, we offer gifts and rewards, and we welcome you to join us every Thursday night.

We take notes of each [biweekly meeting](https://github.com/orgs/OpenIMSDK/discussions/categories/meeting) in [GitHub discussions](https://github.com/OpenIMSDK/OpenKF/discussions/categories/meeting), and our minutes are written in [Google Docs](https://docs.google.com/document/d/1nx8MDpuG74NASx081JcCpxPgDITNTpIIos0DS6Vr9GU/edit?usp=sharing).


## 🤼‍ Contributing & Development

OpenIMSDK Our goal is to build a top-level open source community. We have a set of standards, in the [Community repository](https://github.com/OpenIMSDK/community).

If you'd like to contribute to this OpenKF repository, please read our [contributor documentation](https://github.com/OpenIMSDK/OpenKF/blob/main/CONTRIBUTING.md).

Before you start, please make sure your changes are in demand. The best for that is to create a [new discussion](https://github.com/OpenIMSDK/OpenKF/discussions/new/choose) OR [Slack Communication](https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg), or if you find an issue, [report it](https://github.com/OpenIMSDK/OpenKF/issues/new/choose) first.


## 🚨 License

OpenIMSDK is licensed under the  Apache 2.0 license. See [LICENSE](https://github.com/OpenIMSDK/OpenKF/tree/main/LICENSE) for the full license text.


## 🔮 Thanks to our contributors!

<a href="https://github.com/OpenIMSDK/OpenKF/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=OpenIMSDK/OpenKF" />
</a>
