package org.example.volatilestudy;

/**
 * @Author hideoncode
 * @Date 2025/5/26 22:49
 */
public class VolatileStudy implements Runnable{

    private boolean running = false;

    @Override
    public void run() {
        if (!running) {
            System.out.println("hello");
        }
        running = true;
    }



    public static void main(String[] args) throws InterruptedException {
        VolatileStudy volatileStudy = new VolatileStudy();
        Thread t1 = new Thread(volatileStudy);
        Thread t2 = new Thread(volatileStudy);
        t1.start();
        t2.start();
        t1.join();
        t2.join();

    }
}
