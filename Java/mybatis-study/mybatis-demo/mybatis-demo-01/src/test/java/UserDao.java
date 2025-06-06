import org.example.mybatis.binding.MapperProxyFactory;
import org.junit.Test;

import java.util.HashMap;
import java.util.Map;
import java.util.logging.Logger;

public class UserDao {

    private Logger log = Logger.getLogger(UserDao.class.getName());

    @Test
    public void test_MapperProxyFactory() {
        MapperProxyFactory<IUserDao> factory = new MapperProxyFactory<>(IUserDao.class);
        Map<String, String> sqlSession = new HashMap<>();
        sqlSession.put("queryUserName", "模拟执行Mapper.xml中SQL语句查询用户名的操作");
        sqlSession.put("queryUserAge", "模拟执行Mapper.xml中SQL语句中查询用户年龄的操作");
        IUserDao userDao = factory.newInstance(sqlSession);
        String res = userDao.queryUserName("10001");
        log.info("res:" + res);
    }
}
