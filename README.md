# workerpool
goroutine pool

是一个golang的协程池，尽管golang的goroutine机制很优秀，但是内存的资源是有限制的，有时候我们需要控制同时启动的协程数量等。

~~启发于~~(魔改)[fasthttp](https://github.com/valyala/fasthttp)

- 可指定协程池的容量，控制同时能够运行的协程数量

- 按需创建，根据任务的数量创建工作协程

- 自动协程回收，空闲超过一定时间的协程会自动回收，释放内存

- 排队等待，当协程池的worker都再忙的时候，新来的任务会排队等待

- 可以随时启停pool，停止的时候，正在工作的会等待它工作完成，再排队的会直接结束

- CPU热缓存。。。

  

仅供交流学习使用

PS：在512M的内存的服务器下，已经稳定运行了一年左右