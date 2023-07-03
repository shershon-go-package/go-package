package test

import (
	"shershon1991/go-package/app/bufiopkg"
	"testing"
)

func TestBufIORead(t *testing.T) {
	// 创建读取器
	bufiopkg.Read1() // 创建默认缓冲大小
	//bufiopkg.Read2() // 创建指定缓冲大小

	// 读取指定字节
	//bufiopkg.ReadSomeByte1()

	// 读取一个字节
	//bufiopkg.ReadByte1()

	// 读取一行
	//bufiopkg.ReadLine1() // 当行超过了缓冲缓存时
	//bufiopkg.ReadLine2() // 当行小于缓冲缓存时

	// 读取到指定的字符
	//bufiopkg.ReadString1()
	//bufiopkg.ReadString2()

	// 每次读取前几个字节
	//bufiopkg.ReadPeek1()
	bufiopkg.ReadPeek2()
}

func TestBufIOWrite(t *testing.T) {
	// 创建写入器
	//bufiopkg.Write1() // 创建默认缓冲大小
	//bufiopkg.Write2() // 创建指定缓冲大小

	// 当写入内容大于缓冲区时
	//bufiopkg.WriteContent1()
	// 当写入内容小于缓冲区时
	bufiopkg.WriteContent2()
}
