package main

import (
  "fmt"
  "github.com/ethereum/go-ethereum/rpc"
)

type ETHRPCClient struct {
  NodeUrl string      // 节点的url
  client  *rpc.Client // rpc客户端句柄实例
}

func NewETHRPCClient(nodeUrl string) *ETHRPCClient {
  e := &ETHRPCClient{NodeUrl: nodeUrl}
  e.initRpc()
  return e
}

func (e *ETHRPCClient) initRpc() {
  http, err := rpc.DialHTTP(e.NodeUrl)
  if err != nil {
    panic(fmt.Errorf("初始化rpcClient失败 %s", err.Error()).Error())
  }
  e.client = http
}

func (e *ETHRPCClient) GetRpc() *rpc.Client {
  if e.client == nil {
    e.initRpc()
  }
  return e.client
}
