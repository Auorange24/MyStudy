# lock interface



## AQS

###  AbstractOwnableSynchronizer

**`serialVersionUID`**

用于Java对象序列化、反序列化过程的一致，可以手动设置，也可以自动生成。


**`exclusiveOwnerThread`** 
* 记录独占线程，记录获取当前锁的线程。记录独占线程的原因为了实现可重入功能。 
* 设置了独占线程的设置和获得方法。
* 当下一次获取这个锁的线程与当前持有锁的线程相同时，就可以获取到锁，同时`AQS`的`state`值会加`1`。

```java
public abstract class AbstractOwnableSynchronizer implements java.io.Serializable {
    private static final long serialVersionUID = 3737899427754241961L;
    protected AbstractOwnableSynchronizer() { }
    private transient Thread exclusiveOwnerThread;
    protected final void setExclusiveOwnerThread(Thread thread) {
        exclusiveOwnerThread = thread;
    }
    protected final Thread getExclusiveOwnerThread() {
        return exclusiveOwnerThread;
    }
}
```

### AbstractQueuedSynchronizer

**`Node`静态不可变内部类**

* 共享锁模式下的`Node`
* 独占锁模式下的`Node`
* 标识线程状态的值
  * 