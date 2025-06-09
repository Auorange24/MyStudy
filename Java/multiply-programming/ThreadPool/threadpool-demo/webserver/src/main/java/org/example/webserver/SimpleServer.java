package org.example.webserver;

import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

public class SimpleServer {
    // get cpu core
    private static final int CPU_PORT = Runtime.getRuntime().availableProcessors();
    // IO密集型的期望核心数
    private static final int CORE_POOL_SIZE = CPU_PORT * 2;
    private static final int MAX_POOL_SIZE = CPU_PORT * 2 + 1;
    private static final ThreadPoolExecutor exec = new ThreadPoolExecutor(
            CORE_POOL_SIZE,
            MAX_POOL_SIZE,
            60L,
            TimeUnit.SECONDS,
            new LinkedBlockingQueue<>(1000)
    );

    public static void main(String[] args) {
        while (true) {
            Runnable task = () -> System.out.println("Request Handled By" + Thread.)
        }
    }
}
