# ETHEREUM-RELAY 以太坊中继服务练习

## DApp的分层

### 以太坊可能的应用分层方式1:

这种方式不是完全的DApp，因为上面的web和application层都部署一个中心化的服务。


![DApp应用分层1](https://gitee.com/guozhe001/images/raw/master/DApp%E5%BA%94%E7%94%A8%E5%88%86%E5%B1%821.jpg)

### 以太坊可能的应用分层方式2:

这种分层方式是把一个web3的前端部署到swarm平台，这样没有中间的任何的需要私有部署的应用，只有一个web3.0的web服务，是一个完全的去中心化的服务。

![DApp应用分层2](https://gitee.com/guozhe001/images/raw/master/DApp%E5%BA%94%E7%94%A8%E5%88%86%E5%B1%822-20210330184425534.jpg)

当前项目是根据第一种分层方式开发的以太坊中继服务，主要是练习以太坊JSON RPC API的调用以及go-ethereum工具的使用。

## 运行当前项目

当前项目还无法单独启动，但是单元测试是可以调用本地或者Ropsten测试网络的，需要自己添加`.evn`文件，把助记词放进去。


## 遗留问题和TODO
* eth_call中data参数的格式？
* gas费如何赋值？
* JSON RPC API的所有方法过一遍

