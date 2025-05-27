package org.example.javastudy;

import org.example.javastudy.service.QueueService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class JavaStudyApplicationTests {

    @Autowired
    private QueueService queueService;

    @Test
    public void testQueue() throws InterruptedException {
        new Thread(() -> {
            for (int i = 0; i < 100; i++) {
                queueService.sendMessage("message" + i);
            }
        }).start();
        new Thread(() -> queueService.onMessage()).start();
        Thread.currentThread().join();
    }


}
