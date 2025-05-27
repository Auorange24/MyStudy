package org.example.sync;

import java.util.concurrent.locks.Lock;

/**
 * @Author hideoncode
 * @Date 2025/5/26 10:01
 */
public class AccountingSyncClass implements Runnable{

    static int count_1 = 0;
    static int count_2 = 0;

    public synchronized void increase_1() {
        count_1++;
    }

    public static synchronized void increase_2() {
        count_2++;
    }

    @Override
    public void run() {
        for (int i = 0; i < 100000 ; i ++) {
            increase_1();
            increase_2();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        AccountingSyncClass accountingSyncClass_1 = new AccountingSyncClass();
        AccountingSyncClass accountingSyncClass_2 = new AccountingSyncClass();
        Thread t1 = new Thread(accountingSyncClass_1);
        Thread t2 = new Thread(accountingSyncClass_2);
        t1.start();
        t2.start();
        t1.join();
        t2.join();
        System.out.println("count_1:" + count_1);
        System.out.println("count_2:" + count_2);
    }
}
