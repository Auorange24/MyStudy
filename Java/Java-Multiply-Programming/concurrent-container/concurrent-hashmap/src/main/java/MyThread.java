import java.util.Map;

/**
 * @Author HideOnLife
 * @Date 2025/5/30 15:49
 */
public class MyThread extends Thread implements Runnable{

    UserVisitCounter userVisitCounter;

    Map<String, Integer> map;

    public MyThread(UserVisitCounter userVisitCounter, Map<String, Integer> map) {
        this.userVisitCounter = userVisitCounter;
        this.map = map;
    }


    @Override
    public void run() {
        for (int i = 0 ; i < 100000 ; i ++) {
            userVisitCounter.userVisit(Thread.currentThread().getName());
        }
    }
}
