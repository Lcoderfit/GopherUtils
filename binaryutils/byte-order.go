/*
1.什么是大端序，什么是小端序
2.大端序和小端序内存排列情况
3.Go语言如何判断系统是大端序还是小端序
4.为什么计算机会出现字节序
*/
package binaryutils

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

// 获取操作系统整型占字节数
// 32位操作系统占4个字节，64位操作系统占8个字节
const INTSIZE = int(unsafe.Sizeof(0))

func SystemEndian() {
	var i = 0x01020304
	fmt.Println("&i:", &i)
	// 用来存储整型i的字节数组
	b := (*[INTSIZE]byte)(unsafe.Pointer(&i))

	if b[0] == 0x04 {
		fmt.Println("system endian is little endian")
	} else {
		fmt.Println("system endian is big endian")
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("temp: 0x%x,%v\n", b[i], &b[i])
	}
}

// 测试大端序情况下的数据输出和存储形式
func TestBigEndian() {
	var testInt int32 = 0x01020304
	fmt.Printf("%d use big endian: \n", testInt)

	testBytes := make([]byte, 4)
	// 将整型数通过大端序的方式放入字节数组中
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)
	// 以十六进制的方式打印
	fmt.Printf("int32 to bytes: %x \n", testBytes)

	// 根据大端序打印字节数组的int值
	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

// 测试小端序情况下的数据输出和存储形式

//
//
//func testBigEndian() {
//
//	var testInt int32 = 0x01020304
//	fmt.Printf("%d use big endian: \n", testInt)
//	testBytes := make([]byte, 4)
//	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
//	fmt.Println("int32 to bytes:", testBytes)
//	fmt.Printf("int32 to bytes: %x \n", testBytes)
//
//	convInt := binary.BigEndian.Uint32(testBytes)
//	fmt.Printf("bytes to int32: %d\n\n", convInt)
//}
//
//func testLittleEndian() {
//
//	var testInt int32 = 0x01020304
//	fmt.Printf("%x use little endian: \n", testInt)
//	testBytes := make([]byte, 4)
//	binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
//	fmt.Printf("int32 to bytes: %x \n", testBytes)
//
//	convInt := binary.LittleEndian.Uint32(testBytes)
//	fmt.Printf("bytes to int32: %d\n\n", convInt)
//}
//
//func main() {
//	systemEdian()
//	fmt.Println("")
//	testBigEndian()
//	testLittleEndian()
//}package main
//
