package BLC

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/vendor/golang.org/x/crypto/ripemd160"
	"fmt"
	"bytes"
)
const version = byte(0x00)
const addressCheckSumLen = 4
const addressBytesLen = 25

//钱包数据结构
type Wallet struct{
	//私钥，
	PrivateKey ecdsa.PrivateKey
	//公钥
	PublicKey []byte
}

//新建钱包
func NewWallet()*Wallet{
	private, public := NewKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

func NewKeyPair()(ecdsa.PrivateKey, []byte){
	curve := elliptic.P256() //椭圆曲线
	private, _ := ecdsa.GenerateKey(curve, rand.Reader)	//获取私钥
	publicKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)  //形成公钥
	return *private, publicKey[:]
}

//getAddress
func (wallet Wallet)GetAddress()[]byte{
	//1、先将公钥进行256hash，然后进行160hash
	pubKeyHash := HashPupKey(wallet.PublicKey)
	//2、加上1个字节的version
	versionPayload := append([]byte{version}, pubKeyHash...)

	checkSum := CheckSum(versionPayload)
	fullPayload := append(versionPayload,checkSum...)
	address := Base58Encode(fullPayload)
	return address
}

func HashPupKey(pubKey []byte)[]byte{

	//先进行256hash
	pubKeySha256 := sha256.Sum256(pubKey)
	//再进行160hash，生成20字节的字节数组
	ripemd160Hasher := ripemd160.New()
	_,err := ripemd160Hasher.Write(pubKeySha256[:])

	if err !=nil{
		fmt.Println("Ripemd160 hash failed...")
	}
	return ripemd160Hasher.Sum(nil)
}

func CheckSum(payload []byte)[]byte{
	firstSha := sha256.Sum256(payload)
	secondSha := sha256.Sum256(firstSha[:])
	return secondSha[:addressCheckSumLen]
}

//判断地址是否有效
func IsValidAddress(address []byte)bool{
	//1、先将地址进行Base58反编码
	fullPayload := Base58Decode(address)
	if len(fullPayload) != addressBytesLen{
		return false
	}
	//2、取出前21个字节，求校验和
	versionPayload := fullPayload[0:len(fullPayload)-addressCheckSumLen]
	calccheckSum :=  CheckSum(versionPayload)
	//3、取出最后四个字节的校验和
	addressCheckSum := fullPayload[len(fullPayload)-addressCheckSumLen:]
	//4、比较两个校验和是否相等
	return bytes.Compare(addressCheckSum, calccheckSum) == 0
}