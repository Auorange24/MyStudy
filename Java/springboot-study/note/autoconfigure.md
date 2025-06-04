# 自动配置

## 自动配置的相关注解

## `@springBootApplication` resolve

```java
@Target({ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Inherited
@SpringBootConfiguration
@EnableAutoConfiguration
@ComponentScan(
    excludeFilters = {@Filter(
    type = FilterType.CUSTOM,
    classes = {TypeExcludeFilter.class} // 排除类型过滤器
), @Filter(
    type = FilterType.CUSTOM,
    classes = {AutoConfigurationExcludeFilter.class} // 排除自动配置过滤器
)}
)
```

+ `@Target({ElementType.TYPE})` 表示该注解可以用于类上。
+ `@Retention(RetentionPolicy.RUNTIME)` 表示该注解在运行时生效。
+ `@Documented` 表示该注解会被记录到 Javadoc 中，当鼠标悬停在被该注解修饰的注解之上，能看到注解的具体信息。
+ `@Inherited` 表示该注解可以被继承。
+ `@SpringBootConfiguration` 表示该类是 Spring Boot 的配置类。和普通的配置注解没有区别，只是为了标识这是springboot项目。
+ `@EnableAutoConfiguration` 表示该类是 Spring Boot 的自动配置类。
+ `@ComponentScan` 表示该类是 Spring Boot 的组件扫描类。


## `@EnableAutoConfiguration` resolve

```java
@Target({ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@Documented
@Inherited
@AutoConfigurationPackage
@Import({AutoConfigurationImportSelector.class})
public @interface EnableAutoConfiguration {
    String ENABLED_OVERRIDE_PROPERTY = "spring.boot.enableautoconfiguration";

    Class<?>[] exclude() default {};

    String[] excludeName() default {};
}
```

+ `@AutoConfigurationPackage` 表示该类是 Spring Boot 的自动配置包。将项目src中main包下的所有组件注册到容器中，
例如标识了`@Component`的类，会自动注册到容器中。
+ `@Import({AutoConfigurationImportSelector.class})` 表示该类是 Spring Boot 的自动配置导入类。将这个类引入到容器中，
方便后续使用。

## `AutoConfigurationImportSelector` resolve

+ 扫描类路径文件
+ 根据条件加载自动配置类 


## `@EnableConfigurationProperties`&`@ConfigurationProperties`

二者搭配始终实现了配置文件中的自定义字段到类中的属性中的连接。

**用法**
```java
@ConfigurationProperties(prefix = "person") // 指定配置属性前缀
public class Person {
    private String name;
    private Integer age;
}

@EnableConfigurationProperties(Person.class) // 开启配置属性支持
public class Class {
    @Autowired
    private Person person;
}
```

## 如何实现一个类可以被扫描到

+ 把配置类的路径写到到`META-INF/spring.factories`中
+ 在要自动配置的类上加上`@configuration`注解

> 类路径下载了固定文件夹但是没有使用配置类修饰，也会被加载，但是不会被当成一个配置类。
> 不会解析属性和@Bean方法。

## 自动配置对配置类做了什么

+ 处理配置类。扫描类中带 @Bean 注解的方法——这些方法会被调用，返回的对象被注册为 Spring Bean。
解析类上的 @Import、@ComponentScan，进一步导入其他配置或扫描更多 Bean。
判断各种条件注解（如 @ConditionalOnClass、@ConditionalOnMissingBean 等），决定该配置类/某些 Bean 是否应生效。
+ 只会执行 @Bean 方法
只有加了 @Bean 注解的方法会被实际执行（即：被调用，用以创建 Bean 实例并注册到 Spring 容器）。
其他普通方法不会自动被执行，只有在被其他 Bean 显式调用时才会运行。
+ 配置类里的生命周期方法。
如果配置类实现了如 InitializingBean 等接口、或带有 @PostConstruct，也会在常规 Spring 生命周期流程中执行。

## 如何创建一个starter