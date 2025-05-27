import org.Auorange.demo.jvm.Cmd;
import org.junit.jupiter.api.Test;

import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * @Author hideoncode
 * @Date 2025/5/20 22:23
 */
public class ApiTest {
    // 测试command
    @Test
    public void test() {
        String path = "/classpath";
        Path absolutePath = Paths.get(path).toAbsolutePath();
        System.out.println(absolutePath);
    }
}
