package org.example.sync.atomic;

import io.netty.channel.Channel;
import sun.misc.Unsafe;

import java.util.concurrent.atomic.AtomicInteger;

/**
 * @Author hideoncode
 * @Date 2025/5/26 15:14
 */
public class Atomic implements Runnable{

    static AtomicInteger i = new AtomicInteger(0);
    static int non_count = 0;

    public void non_increase() {
        non_count ++;
    }

    public void increase() {
        i.getAndAdd(1);
    }

    @Override
    public void run() {
        for (int i = 0 ; i < 1000 ; i ++) {
           increase();
           non_increase();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Atomic atomic = new Atomic();
        Thread t1 = new Thread(atomic);
        Thread t2 = new Thread(atomic);
        t1.start();
        t2.start();
        t1.join();
        t2.join();
        System.out.println("i = " + i);
        System.out.println("non_count = " + non_count);
    }

}
