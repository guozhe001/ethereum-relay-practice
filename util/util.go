package util

import (
  "fmt"
  "github.com/guozhe001/ethereum-relay-practice/constant"
  "log"
  "math/big"
  "os"
  "path"
  "path/filepath"
  "strings"
)

func init() {
  GetProjectAbsPath()
}

// WeiToEth 单位转换，把wei转换成eth
func WeiToEth(hexWei string) (string, error) {
  ten, _ := new(big.Int).SetString(hexWei[2:], 16)
  log.Println(ten)
  return ten.String(), nil
}

// GetZero 获取指定数量的0的字符串；主要是用在调用以太坊时组装data
func GetZero(number int) (string, error) {
  if number < 1 || number > 64 {
    return "", fmt.Errorf("只能获取1到64个0组成的字符串")
  }
  return constant.ZeroOf64[:number], nil
}

// HexToBigInt 把带"0x"前缀的十六进制的字符串转换成十进制的big.Int
func HexToBigInt(hex string) (*big.Int, bool) {
  return new(big.Int).SetString(hex[2:], 16)
}

// BigIntToHex 把big.Int转换成十六进制的字符串
func BigIntToHex(num big.Int) string {
  return ""
}

// GetRealDecimalValue 根据代币的 decimal 得出乘上 10^decimal 后的值
// value 是包含浮点数的。例如 0.5 个 ETH
func GetRealDecimalValue(value string, decimal int) string {
  if strings.Contains(value, ".") {
    // 小数
    arr := strings.Split(value, ".")
    if len(arr) != 2 {
      return ""
    }
    num := len(arr[1])
    left := decimal - num
    return arr[0] + arr[1] + strings.Repeat("0", left)
  } else {
    // 整数
    return value + strings.Repeat("0", decimal)
  }
}

func GetProjectAbsPath() (projectAbsPath string) {
  programPath, _ := filepath.Abs(os.Args[0])

  fmt.Println("programPath:", programPath)

  projectAbsPath = path.Dir(path.Dir(programPath))

  fmt.Println("PROJECT_ABS_PATH:", projectAbsPath)

  return projectAbsPath
}
