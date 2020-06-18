[English](./README.md) | 简体中文

# LastOrder

使用[Golang](https://github.com/golang/go)编写的网关服务

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

<a href="https://www.jetbrains.com/?from=LastOrder" target="_blank"><img src="./images/jetbrains-variant-4.png" width="250" align="middle"/></a>

## FAQ

### 命名由来？

> 动漫作品《[魔法禁书目录](https://baike.baidu.com/item/魔法禁书目录/25423)》中的人物。[御坂网络](https://baike.baidu.com/item/御坂网络/8582829?fr=aladdin)中的上位个体，检体番号 20001 号，是所有御坂妹妹的司令塔，是为了防止“妹妹们”反叛、失控而制作出来的安全装置。和御坂网络中的其他个体不同，具有较丰富的表情，头顶上有一根呆毛。

<img src="https://gss2.bdstatic.com/-fo3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike150%2C5%2C5%2C150%2C50/sign=609b31fe047b020818c437b303b099b6/bf096b63f6246b6002aeab2fe5f81a4c500fa2cc.jpg" width="200" height="400" />

## License

[MIT](https://tldrlegal.com/license/mit-license)
