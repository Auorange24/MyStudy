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

**线程池刚创建时是什么样子的**

初始化了核心线程数、最大线程数、阻塞队列。

此时并没有初始化线程，如果想在执行之前创建线程，可以调用`prestartAllCoreThreads`方法。该
方法会启动所有核心线程，使它们空闲等待任务，会覆盖默认的策略，只有在执行新任务时才会启动核心线程。

```java
public int prestartAllCoreThreads() {
    int n = 0;
    while (addWorker(null, null)) {
        ++ n;
    }
    return n;
}
```

**当执行`execute`方法提交任务会发生什么**

1. 判断当前线程数是否小于核心线程数`corePoolSize`：
    + 如果小于，通过`ThreadFactory`创建一个线程执行任务。
2. 如果当前线程数大于等于核心线程数：
    + 存在空闲线程直接执行。
    + 不存在空闲线程加入阻塞队列中。
3. 如果空闲队列满了怎么办：
    + 如果当前线程数小于最大线程数，则创建非核心线程进行处理。
    + 如果当前线程数等于最大线程数，则使用`RejectedExecurionHandler`对象来处理这个任务。

> 当提交任务之后，如果线程数小于核心线程数且存在空闲线程，还是会创建新线程而不是复用旧线程。


**`execute`的执行流程**

```java
public void execute(Runnable command) {
    // task is null, throw exception
    if (command == null) {
        throw new NullPointerException();
    }
    // get threadpool status and working thread number
    // A atomic integer, hign bit low bit 
    int c = ctl.get();
    // 1. 检查当前运行的工作线程数是否少于核心线程数（corePoolSize）
    if (workerCountOf(c) < corePoolSize) {
        // 如果少于核心线程数，尝试添加一个新的工作线程来执行提交的任务
        // addWorker方法会检查线程池状态和工作线程数，并决定是否真的添加新线程
        if (addWorker(command, true))
            return;
        // 重新获取线程池的状态，因为在尝试添加线程的过程中线程池的状态可能已经发生变化
        c = ctl.get();
    }
    // 2. 尝试将任务添加到任务队列中
    if (isRunning(c) && workQueue.offer(command)) {
        int recheck = ctl.get();
        // 双重检查线程池的状态
        if (! isRunning(recheck) && remove(command))  // 如果线程池已经停止，从队列中移除任务
            reject(command);
        // 如果线程池正在运行，但是工作线程数为0，尝试添加一个新的工作线程
        else if (workerCountOf(recheck) == 0)
            addWorker(null, false);
    }
    // 3. 如果任务队列满了，尝试添加一个新的非核心工作线程来执行任务
    else if (!addWorker(command, false))
        // 如果无法添加新的工作线程（可能因为线程池已经停止或者达到最大线程数限制），则拒绝任务
        reject(command);
}
```

# 四、线程池复用的原理

**线程在线程池内部如何表示**

线程在线程池内部表示成为`worker`类。

```java
/*
实现了Runnable接口，是一个线程
即成了AQS，拥有锁的特性
*/
private final class Worker 
    extends AbstractQueuedSynchronizer
    implements Runnable {}
```

创建线程使用`addWorker`方法。
调用线程执行任务使用`runWorker`方法。

```java
final void runWorker(Worker w) {
    // 获取当前工作线程
    Thread wt = Thread.currentThread();
    // 从线程中提取第一个任务
    Runnable task = w.firstTask;
    w.firstTask = null;
    // 解锁Worker
    w.unlock();
    boolean completeAbruptly = true;
    try {
        while (task != null || (task = getTask() != null)) {
            // 锁定Worker，确保在执行任务期间不会被其他线程干扰
            w.lock();
            // 如果线程正在停止，并确保线程已经中断
            // 如果线程没有中断并且线程已经达到停止状态，终止线程
            if ((runStateAtLeast(ctl.get(), STOP) || (Thread.interrupted() && ))) {
                wt.interrupt();
            }
            try {
                // 在执行任务之前，可以插入一些自定义操作
                beforeExecute(wt, task);
                Throwable thrown = null;
                try {
                    // 执行任务
                    task.run();
                } catch (RuntimeException x) {
                    thrown = x;
                    throw x;
                }
            }
        }
    }
}
```



# 五、线程池如何获取任务

> 线程池如何获取任务并且实现超时

线程在执行完任务之后，会继续执行`getTask`方法获取任务，获取不到就退出。

```java
private Runnable getTask() {
    // 标志，表示最后一个poll()操作是否超时
    boolean timeOut = false;
    // 无限循环，直到获取到任务或决定工作线程应该退出
    for (;;) {
        // 获取线程状态和当前线程数目
        int c = ctl.get();
        // 获取线程状态
        int rs = runStateOf(c);
        if (rs >= SHUTDOWN && (rs >= STOP || workQueue.isEmpty())) {
            decrementWorkerCount();
            return null;
        } 
    }
}
```
