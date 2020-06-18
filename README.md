[![](https://api.travis-ci.com/MisakaSystem/LastOrder.svg?branch=master)](https://travis-ci.com/MisakaSystem/LastOrder)
[![](https://img.shields.io/github/license/MisakaSystem/LastOrder)](https://opensource.org/licenses/MIT)
![](https://img.shields.io/github/v/release/MisakaSystem/LastOrder)
[![Go Report Card](https://goreportcard.com/badge/github.com/MisakaSystem/LastOrder)](https://goreportcard.com/report/github.com/MisakaSystem/LastOrder)

English | [简体中文](./README_CN.md)

# LastOrder

A Gateway written by [Golang](https://github.com/golang/go)

Use etcd store serve config and routing config

1. [Website]()
2. [Release]()
3. [Roadmap](https://github.com/MisakaSystem/LastOrder-roadmap)

## Feature

1. Routing


## Set up config
### 1.1 Install owl
use owl store config into the ectd
- use go install owl
```shell script
go install github.com/gsxhnd/owl
```
- download binary
```shell script
wget https://github.com/gsxhnd/owl/releases/download/v0.3.0/owl-0.3.0-linux64-amd64
mv owl-0.3.0-linux64-amd64 /usr/local/bin/owl
chmod +x /usr/local/bin/owl
```

### 1.2 upload/update config into ectd
```shell script
owl put -e "local_dev:2379" /conf/gateway.yaml ./conf/gateway.yaml
```

### 1.3 check current config in ectd
```shell script
owl get -e "local_dev:2379" /conf/cdn.yaml
```

## Install LastOrder
```bash
wget https://github.com/gsxhnd/owl/releases/download/v0.3.0/owl-0.3.0-linux64-amd64
wget https://github.com/gsxhnd/owl/releases/download/v0.3.0/owl-0.3.0-linux64-amd64
wget https://github.com/gsxhnd/owl/releases/download/v0.3.0/owl-0.3.0-linux64-amd64
```
## Run Server
```bash
last_order run --etcds="127.0.0.1:2379" /conf/gateway.yaml
```

## JetBrains OS licenses

`LastOrder` had been being developed with GoLand under the **free JetBrains Open Source license(s)** granted by JetBrains s.r.o., hence I would like to express my thanks here.

<a href="https://www.jetbrains.com/?from=LastOrder" target="_blank"><img src="https://github.com/gsxhnd/archive/blob/master/jetbrains-variant-4.png?raw=true" width="250" align="middle"/></a>

## License

[MIT](https://tldrlegal.com/license/mit-license)
