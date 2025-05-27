package org.example.sync;

/**
 * @Author hideoncode
 * @Date 2025/5/26 09:49
 */
public class AccountingSync implements Runnable{

    static int count = 0;

    public synchronized void increase() {
        count ++;
    }

    @Override
    public void run() {
        for (int i = 0 ; i < 100000 ; i ++) {
            increase();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        AccountingSync accountingSync = new AccountingSync();
        Thread t1 = new Thread(accountingSync);
        Thread t2 = new Thread(accountingSync);
        t1.start();
        t2.start();
        t1.join();
        t2.join();
        System.out.println(count);
    }
}
