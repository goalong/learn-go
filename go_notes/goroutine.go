package main


// 入口函数main就以goroutine运行
// 默认情况进程启动仅允许一个系统线程服务于goroutine, 可使用环境变量或
// 标准库函数runtime.GOMAXPROCS修改，让调度器用多个线程实现多核并行，而不仅仅是并发

// runtime.Goexit将立即终止当前goroutine执行，调度器将确保已注册的defer延迟调用并执行
// runtime.Goshed让出底层线程，将当前goroutine暂停，放回队列等待下次被调度执行
