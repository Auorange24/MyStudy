package org.example.sync.biasedLock;

/**
 * @Author hideoncode
 * @Date 2025/5/26 10:49
 */
public class BiasedLock implements Runnable {
    static int i = 0;

    @Override
    public void run() {
        for (int i = 0 ; i < 1000 ; i ++) {
            i ++;
        }
    }

    public static void main(String[] args) throws InterruptedException {
        BiasedLock biasedLock = new BiasedLock();
        Thread t1 = new Thread(biasedLock);
        Thread t2 = new Thread(biasedLock);
        t1.start();
        t2.start();
        t1.join();
        t2.join();
        System.out.println("i = " + i);
    }


}
