package test

import (
	"context"
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"testing"
	"time"
)

func TestUseTrace(t *testing.T) {
	//runtime.GOMAXPROCS()
	selectDB()
	writeFile()
	data := map[int]time.Time{}
	for i := 0; i < 100; i++ {
		data[i] = time.Now()
	}
	fmt.Println("len=", len(data))
}

func selectDB() {
	time.Sleep(time.Second * 1)
}

func writeFile() {
	time.Sleep(time.Second * 2)
}

func TestCreateCou(t *testing.T) {
	file, err := os.Create("./mytask.out")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer file.Close()
	err = trace.Start(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer trace.Stop()

	// 创建自定义任务
	ctx, task := trace.NewTask(context.Background(), "myTask")
	defer task.End()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		gn := i
		// 启动协程
		wg.Add(1)
		go func() {
			defer wg.Done()
			trace.WithRegion(ctx, fmt.Sprintf("goroutine-%d", gn), func() {
				sum := 0
				for n := 0; n < 1000000; n++ {
					sum = sum + n
				}
				fmt.Println("sum = ", sum)
			})
		}()
	}
	wg.Wait()
	fmt.Println("run ok!")
}
