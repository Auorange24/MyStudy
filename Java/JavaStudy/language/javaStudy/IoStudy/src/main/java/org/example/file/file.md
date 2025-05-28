# File Stream

## File Class

* `Java File Class`位于`java.io.File`下，只能对文件本身进行操作，不能对文件内容进行操作，想要对文件内容进行操作，必须要借助输入输出流。

> 应该不是操作系统中的`file`指针，`file`指针可以对文件进行操作。

## File Constructor

```java
public File(String pathname);
public File(String parent, String child);
public File(File parent, String child);
```

**关于文件分隔符**

+ `MacOS`的分隔符为`/`。
+ `Windows`的分隔符为`\`。
+ 使用`Java.separator`跨平台获取分隔符，这个属性会根据操作系统自动返回正确的路径分隔符。

**`Java`类的注意点**

+ 一个`File`对象代表硬盘中实际存在的一个文件或者目录。
+ `File`类的构造方法不会检验这个文件或者目录是否真是存在，无论文件、目录是否存在，都会创建`File`对象。

**创建类实例**

```Java
public class UseFile {

    private final static String path = File.separator + "Users" +
            File.separator + "hideoncode" +
            File.separator + "Desktop" +
            File.separator + "code" +
            File.separator + "MyStudy" +
            File.separator + "Java" +
            File.separator + "JavaStudy" +
            File.separator + "language" +
            File.separator + "javaStudy" +
            File.separator + "IoStudy" +
            File.separator + "src" +
            File.separator + "main" +
            File.separator + "java" +
            File.separator + "org" +
            File.separator + "example" +
            File.separator + "file" +
            File.separator + "123.txt";

    public static void main(String[] args) {
        File file = new File(path);
        System.out.println(file.getAbsolutePath());
        System.out.println(file.isFile());
        System.out.println(file.length());
    }
}
```

## File Method

**`File`对象的方法**

+ `getAbsolutePath`，返回此`File`的绝对路径。
+ `getPath`，返回此`File`的绝对路径。
+ `getName`，返回文件名、目录名。
+ `length`，返回文件长度，以字节为单位。`File`为目录时该返回值无意义。

**判断功能的方法**

+ `exists`，判断文件、目录是否存在。
+ `isDirectory`，判断是否问目录。
+ `isFile`，判断是否为文件。

**修改文件**
