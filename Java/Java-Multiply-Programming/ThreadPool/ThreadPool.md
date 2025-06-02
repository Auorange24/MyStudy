# ThreadPol

> 线程池

# 一、什么是线程池

线程池是一种池化技术。可以实现资源的复用，避免资源重复创建和销毁带来的额外开销。

**线程池的好处**

+ 降低资源消耗
+ 提高响应速度
+ 提高线程的可管理性

# 二、线程池的构造

Java使用`ThreadPoolExecutor`类创建线程池。一个该类实例即是一个线程池。

**ThreadPoolExecutor Constructor Method**

```java
public ThreadPoolExecutor(
    int corePoolSize,
    int maximumPoolSize,
    long keepAliveTime,
    TimeUnit unit,
    BlockingQueue<Runnable> workQueue,
    ThreadFactory ThreadFactory,
    RejectedExecitionHandler handler
) {
    if (corePoolSize < 0)
}
```

# 三、线程池的运行原理

