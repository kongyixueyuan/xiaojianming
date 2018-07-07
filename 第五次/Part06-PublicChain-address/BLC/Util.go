package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
	"encoding/json"
)

func IntToHex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil{
		log.Panic(err)
	}
	return buff.Bytes()
}

func JSONToArray(jsonstring string)[]string{
	var sArr []string
	if err := json.Unmarshal([]byte(jsonstring), &sArr); err !=nil{
		log.Panic("Json转换为string数组失败",err)
	}
	//fmt.Println(sArr)
	return sArr
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}