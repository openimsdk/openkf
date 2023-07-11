# Changelog

All notable changes to this project will be documented in this file.

+ [https://github.com/OpenIMSDK/OpenKF/releases](https://github.com/OpenIMSDK/OpenKF/releases)

## command

```bash
git-chglog --tag-filter-pattern 'v2.0.*'  -o CHANGELOG-2.0.md
```

## create next tag

```bash
git-chglog --next-tag 2.0.0 -o CHANGELOG.md
git commit -am "release 2.0.0"
git tag 2.0.0
```

## Release version logs

+ [OpenIM CHANGELOG-V0.1](CHANGELOG-0.1.md)