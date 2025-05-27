package org.example.javastudy.service;

import lombok.extern.slf4j.Slf4j;
import org.redisson.api.RBlockingDeque;
import org.redisson.api.RedissonClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Slf4j
@Service
public class QueueService {
    @Autowired
    private RedissonClient redissonClient;
    private static final String REDIS_MQ = "redisMQ";

    public void sendMessage(String message) {
        RBlockingDeque<String> blockingDeque = redissonClient.getBlockingDeque(REDIS_MQ);
        try {
            blockingDeque.putFirst(message);
//            log.info("将消息: {} 插入到队列。", message);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    public void onMessage() {
        RBlockingDeque<String> queue = redissonClient.getBlockingDeque(REDIS_MQ);
        while (true) {
            try {
                String message = queue.takeLast();
//                log.info("从队列中读取到消息:{}.", REDIS_MQ, message);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

}
