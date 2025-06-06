package org.example.mybatis.binding;

import java.lang.reflect.Proxy;
import java.util.Map;


public class MapperProxyFactory<T> {
    // 定义的数据库查询接口
    private final Class<T> mapperInterface;
    public MapperProxyFactory(Class<T> mapperInterface) {
        this.mapperInterface = mapperInterface;
    }
    // 利用数据库查询接口类创建代理类
    public T newInstance(Map<String, String> sqlSession) {
        final MapperProxy<T> mapperProxy = new MapperProxy<>(mapperInterface, sqlSession);
        return (T)Proxy.newProxyInstance(mapperInterface.getClassLoader(), new Class[]{mapperInterface}, mapperProxy);
    }
}
