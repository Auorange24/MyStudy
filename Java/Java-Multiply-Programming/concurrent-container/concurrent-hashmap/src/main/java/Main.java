import java.util.HashMap;
import java.util.Map;

/**
 * @Author HideOnLife
 * @Date 2025/5/30 15:52
 */
public class Main {
    public static void main(String[] args) throws InterruptedException {
        UserVisitCounter userVisitCounter = new UserVisitCounter();
        Map<String, Integer> map = new HashMap<>();
        MyThread t1 = new MyThread(userVisitCounter, map);
        t1.setName("user1");
        MyThread t2 = new MyThread(userVisitCounter, map);
        t2.setName("user2");
        t1.start();
        t2.start();
        t1.join();
        t2.join();
        System.out.println("ConcurrentMap:");
        System.out.println("user1 : " + userVisitCounter.getUserVisitCounter("user1"));
        System.out.println("user2 : " + userVisitCounter.getUserVisitCounter("user2"));

        System.out.println("NonConcurrentHashMap:");
        System.out.println("user1 : " + userVisitCounter.getUserVisitCounterByMap("user1"));
        System.out.println("user2 : " + userVisitCounter.getUserVisitCounterByMap("user2"));
    }
}
