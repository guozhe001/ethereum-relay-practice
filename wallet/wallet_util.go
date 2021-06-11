package wallet

import (
  "errors"
  "fmt"
  "github.com/ethereum/go-ethereum/accounts"
  "github.com/ethereum/go-ethereum/accounts/keystore"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  "log"
  "path"
  "runtime"
)

const keystoreDir = ".keystores"

// 全局的保存了已经解锁成功了的钱包 map 集合变量
var ETHUnlockMap map[string]accounts.Account

// 全局的对应 keystore 实例
var UnlockKs *keystore.KeyStore

// CreateWallet，创建新的钱包
func CreateWallet(password string) (string, error) {
  if "" == password {
    return "", fmt.Errorf("please input password！")
  }
  if len(password) < 6 {
    return "", fmt.Errorf("password length must big than 6！")
  }
  ks := keystore.NewKeyStore(getKeystoreDir(), keystore.StandardScryptN, keystore.StandardScryptP)
  account, err := ks.NewAccount(password)
  if err != nil {
    return "", err
  }
  return account.Address.String(), nil
}

// 解锁以太坊钱包，传入钱包地址和对应的 keystore密码
func UnlockETHWallet(address, password string) error {
  if UnlockKs == nil {
    UnlockKs = keystore.NewKeyStore(
      // 服务端存储 keystore 文件的目录
      // 这些配置类的信息可以由配置文件指定
      getKeystoreDir(),
      keystore.StandardScryptN,
      keystore.StandardScryptP)
    if UnlockKs == nil {
      return errors.New("ks is nil")
    }
  }
  unlock := accounts.Account{Address: common.HexToAddress(address)}
  // ks.Unlock 调用 keystore.go 的解锁函数，解锁出的私钥将存储在它里面的变量中
  if err := UnlockKs.Unlock(unlock, password); nil != err {
    return errors.New("unlock err ： " + err.Error())
  }
  if ETHUnlockMap == nil {
    ETHUnlockMap = map[string]accounts.Account{}
  }
  ETHUnlockMap[address] = unlock // 解锁成功，存储
  return nil
}

// 签名交易数据结构体 types.Transaction
func SignETHTransaction(address string, transaction *types.Transaction) (*types.Transaction, error) {
  if UnlockKs == nil {
    return nil, errors.New("you need to init keystore first")
  }
  account := ETHUnlockMap[address]
  if !common.IsHexAddress(account.Address.String()) {
    // 判断当前地址钱包是否解锁了
    return nil, errors.New("account need to unlock first")
  }
  return UnlockKs.SignTx(account, transaction, transaction.ChainId()) // 调用签名函数
}

// getKeystoreDir 获取Keystore所在的路径
func getKeystoreDir() string {
  keyStoreDir := path.Join(getCurrentPath(), keystoreDir)
  log.Println("path=", keyStoreDir)
  return keyStoreDir
}

// getCurrentPath 获取当前的路径
func getCurrentPath() string {
  _, filename, _, ok := runtime.Caller(1)
  var cwdPath string
  if ok {
    cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
  } else {
    cwdPath = "./"
  }
  return cwdPath
}
