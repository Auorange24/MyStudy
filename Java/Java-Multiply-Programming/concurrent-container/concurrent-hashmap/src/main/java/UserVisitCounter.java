import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * @Author HideOnLife
 * @Date 2025/5/30 15:37
 */

/*
* Use Concurrent Hashmap
* */
public class UserVisitCounter {
    // define static final concurrent container
    private final ConcurrentHashMap<String, Integer> visitCounter;
    private final Map<String, Integer> map;

    public UserVisitCounter() {
        this.visitCounter = new ConcurrentHashMap<>();
        this.map = new HashMap<>();
    }

    // when user request, call userVisited
    public void userVisit(String userId) {
        visitCounter.compute(userId, (key, value) -> value == null ? 1 : value + 1);
        map.compute(userId, (key, value) -> value == null ? 1 : value + 1);
    }

    // get user visited count
    public int getUserVisitCounter(String userId) {
        return visitCounter.getOrDefault(userId, 0);
    }

    public int getUserVisitCounterByMap(String userId) {
        return map.getOrDefault(userId, 0);
    }

}
