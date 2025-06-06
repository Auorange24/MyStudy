# sync.Pool

## sync.Pool Method

```go
package sync
type pool struct {
    New func() any
    ...
}
func (p *sync.pool) Get() any {}
func (p *sync.pool) Put(v any)
```

+ **`pool.New`**

用于创建对象的函数，当对象池中没有足够的对象复用时，使用该方法创建新的对象。

> sync.pool是并发安全的。
> 但是New函数内部的执行逻辑要由注册方维持其并发安全。

+ **`Get`**

从对象池中并发安全的获取对象。

如果对象池中还有对象，则直接返回。

如果对象池中不存在对象，则使用New函数创建新对象返回。

> 对象池返回的是any类型的变量，需要进行类型断言。


+ **`Put`**

将对象并发安全地归还到对象池中。可以结合defer命令确保归还操作一定会执行。

Put操作仅仅是将实例返回池子，对象何时被释放由GC控制。

> 当实例返回到对象池中，对象池不会对其作任何清理操作。
>
> 若立即执行Get操作，可能会得到脏数据的实例，也可能是New新构造的实例。
>
> 所以Put前或者Get时，应该统一对实例进行数据清理。

## sync.pool 实现原理

> 基于GMP架构中的processor

**核心要点**

+ 为每个处理器P创建一个`private`私有实例和`shared`共享实例
+ 对`P`的`private`私有实例，`P`可以无锁化访问。Get、Put操作都会优先操作私有实例。
+ 对`P`的`shared`共享对象实例列表，`P`访问需要通过`CAS`操作保证并发安全。

## sync.pool source code

**sync.pool struct**

```go
type Pool struct {
    noCopy noCopy
    local unsafe.Pointer
    localSize uintptr
    victim unsafe.Pointer
    victimSize unintptr
    New func() any
}

type poolLocal struct {
    poolLocalInternal 
    private  any
    shared poolChain
}

type poolChain struct {
    head, tail *poolChainElt
}

type 
```


### 获取对象实例

**Get**

```go
func (p *Pool) Get() any {
    /*
    将当前Goroutinues绑定在当前的processor中
    防止将该goroutinues调度到其他P上
    */ 
    l, pid := p.pin() // 返回P对应的local容器、P id
    // 尝试获取loca容器中的当前P的私有private实例
    x := l.private
    l.private = nil // 获取到对象之后该对象不能再次被复用，设置为nil
    if x == nil {
        // 尝试从当前P shared中获取对象实例
        x, _ = l.shared.popHead()
        if x == nil {
            // 从victim中获取
            x = p.getSlow(pid)
        }
    }
    // 将P与当前goroutinues解绑操作
    runtime_procUnpin()
    // 还未获得对象实例且New不为空，则使用New函数创建
    if x == nil && p.New != nil {
        x = p.New()
    }
    return x
}
```

**`popHead()`**

从当前P对应的local的shared共享队列中获取实例

```go
func (c *poolChain) popHead() (any, bool) {
    d := c.head 
    // 遍历环形链表
    for d != nil {
        if val, ok := d.popHead() ; ok {
            return val, ok
        }
        d = loadPoolChainElt(&d.prev)
    }
    return nil, false
}
```

**`getSlow`**

从当前P的local中的获取实例失败之后，会执行getSlow流程，从其他P的loca中获取对象实例。

```go
func (p *Pool) getSlow(pid int) any {
    // 读取全局 locals 数组的长度
    size := runtime_LoadAcquintptr(&p.locasize)
    // 获取全局 locals 数组
    locals := p.local
    // 沿着pid + 1d的p对于的locals
}
```



### 清理对象实例