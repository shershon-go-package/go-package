/**
 * @Author: shershon
 * @Description:
 * @Date: 2023/06/15 20:31
 */

package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Read1() {
	// 创建默认缓冲大小
	reader := bufio.NewReader(strings.NewReader("hello world!"))
	fmt.Printf("默认大小:%v \n", reader.Size())
}

func Read2() {
	reader := bufio.NewReaderSize(strings.NewReader("hello world"), 40)
	fmt.Printf("大于16字节:%v \n", reader.Size())
	reader2 := bufio.NewReaderSize(strings.NewReader("hello world"), 4)
	fmt.Printf("小于16字节:%v \n", reader2.Size())
}

func ReadSomeByte1() {
	file, err := os.Open("./public/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		// 一次读取21个字节
		content := make([]byte, 21)
		_, err := reader.Read(content)
		// 读取完毕，则跳出
		if err == io.EOF {
			break
		}
		fmt.Printf("读取内容:%s", content)
	}
}

func ReadByte1() {
	reader := bufio.NewReader(strings.NewReader("Go,word!"))
	for {
		// 一次读取1个字节
		readByte, err := reader.ReadByte()
		// 读取完毕，则跳出
		if err == io.EOF {
			break
		}
		fmt.Printf("读取内容:%s \n", string(readByte))
	}
}

func ReadLine1() {
	// 创建字符串，每行超过16个字节
	str := strings.Repeat("Hello world, nice to meet you!\n", 3)
	// 基于字符串创建一个缓冲区=16字节的读取器
	reader := bufio.NewReaderSize(strings.NewReader(str), 16)
	for {
		// 每次读取一行
		line, prefix, err := reader.ReadLine()
		// 读取完毕，则跳出
		if err == io.EOF {
			break
		}
		fmt.Printf("line:%s, isPrefix:%t \n", line, prefix)
	}
}

func ReadLine2() {
	// 每行不超过16个字节
	str := strings.Repeat("Hello world!\n", 3)
	// 基于字符串创建读取器
	reader := bufio.NewReaderSize(strings.NewReader(str), 16)
	for {
		// 每次读取一行
		line, prefix, err := reader.ReadLine()
		// 读取完毕，则跳出
		if err == io.EOF {
			break
		}
		fmt.Printf("line:%s isPrefix:%t \n", line, prefix)
	}
}

func ReadString1() {
	reader := bufio.NewReaderSize(strings.NewReader("Go,PHP,Java,Python,C"), 16)
	for {
		s, err := reader.ReadString(',')
		// 读取完毕，则跳出
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("err:%s", err.Error())
			break
		}
		fmt.Printf("%s\n", s)
	}
}

func ReadString2() {
	// 设置每行超过缓冲区大小
	repeat := strings.Repeat("Hello,Hello,Hello,Hello\n", 3)
	reader := bufio.NewReaderSize(strings.NewReader(repeat), 16)
	for {
		// 读取
		s, err := reader.ReadString('\n')
		// 读取完毕，则跳出
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("err:%s", err.Error())
			break
		}
		fmt.Printf("%s", s)
	}
}

func ReadPeek1() {
	reader := bufio.NewReaderSize(strings.NewReader("hello"), 16)
	// 读取
	for i := 0; i < 2; i++ {
		s, err := reader.Peek(16)
		fmt.Printf("结果: %s ", s)
		if err == io.EOF {
			fmt.Printf("%s", "已读完")
			break
		} else if err != nil {
			fmt.Printf("err:%s", err.Error())
			break
		}
	}
}

func ReadPeek2() {
	// 基于字符串创建读取器
	reader := bufio.NewReaderSize(strings.NewReader("hello"), 16)
	// 读取
	for i := 0; i < 2; i++ {
		// 大于缓冲区大小
		s, err := reader.Peek(17)
		fmt.Printf("结果: %s ", s)
		if err == io.EOF {
			fmt.Printf("%s", "已读完")
			break
		} else if err != nil {
			fmt.Printf("err:%s", err.Error())
			break
		}
	}
}

func Write1() {
	// 创建默认缓冲大小
	file, err := os.OpenFile("./public/b.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Printf("默认大小:%v \n", writer.Size())
}

func Write2() {
	// 创建默认缓冲大小
	file, err := os.OpenFile("./public/b.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriterSize(file, 40)
	fmt.Printf("大于16字节:%v \n", writer.Size())
	writer2 := bufio.NewReaderSize(file, 4)
	fmt.Printf("小于16字节:%v \n", writer2.Size())
}

func WriteContent1() {
	file, _ := os.OpenFile("./public/test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	reader := bufio.NewWriterSize(file, 20)
	// 当写入的内容字节大于缓冲区大小时,会直接写入文件
	write, err := reader.Write([]byte("hello,hello,hello 你好！"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("直接写入文件: %d \n", write)
}

func WriteContent2() {
	file, _ := os.OpenFile("./public/test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	reader := bufio.NewWriterSize(file, 20)
	// 当写入的内容字节小于缓冲区大小时,不会直接写入文件
	write, err := reader.Write([]byte("\nhello,Go!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 该函数会使缓冲区的内容，直接写入文件
	_ = reader.Flush()
	fmt.Printf("Flush写入文件: %d \n", write)
}
