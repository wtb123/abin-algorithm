## 一、典型并发任务：主程序等待另一个程序执行完一起结束

1. 传统的并发编程都是通过 锁 + 共享内存方式来实现并发下的消息传递、数据传递
2. go 主要是通过CSP，可以理解用 channel 来实现并发下的消息传递、数据传递
3. 数据传递 VS 消息传递（线程、协程控制）

```go
// sync.WaitGroup 也可以实现同样的功能

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

```

## 二、多路选择和超时















