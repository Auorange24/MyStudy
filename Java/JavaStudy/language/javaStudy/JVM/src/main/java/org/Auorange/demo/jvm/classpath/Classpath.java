package org.Auorange.demo.jvm.classpath;

import java.nio.file.Files;
import java.nio.file.Paths;

/**
 * @Author hideoncode
 * @Date 2025/5/22 11:41
 */
public class Classpath {

    private Entry bootstrapClasspath; //启动类路径
    private Entry extensionClasspath; //扩展类路径
    private Entry userClassPath; //用户类路径

    public Classpath(String jreOption, String cpOption) {
        //启动类&扩展类 "C:\Program Files\Java\jdk1.8.0_161\jre"
        bootStrapAndExtensionClasspath(jreOption);
        //用户类 E:\..\org\itstack\demo\test\HelloWorld
        parseUserClasspath(cpOption);
    }

    private void bootStrapAndExtensionClasspath(String jreOption) {
        String jreDir = getJreDir(jreOption);

        //..jre/lib/*
        String jreLibPath = Paths.get(jreDir, "lib") + File.separator + "*";
        bootstrapClasspath = new WildcardEntry(jreLibPath);

        //..jre/lib/ext/*
        String jreExtPath = Paths.get(jreDir, "lib", "ext") + File.separator + "*";
        extensionClasspath = new WildcardEntry(jreExtPath);
    }

    private static String getJreDir(String jreOption) {
        if (jreOption != null && Files.exists(Paths.get(jreOption))) {
            return jreOption;
        }
        if (Files.exists(Paths.get("./jre"))) {
            return "./jre";
        }
        String jh = System.getenv("JAVA_HOME");
        if (jh != null) {
            return Paths.get(jh, "jre").toString();
        }
        throw new RuntimeException("Can not find JRE folder!");
    }

    private void parseUserClasspath(String cpOption) {
        if (cpOption == null) {
            cpOption = ".";
        }
        userClasspath = Entry.create(cpOption);
    }

    public byte[] readClass(String className) throws Exception {
        className = className + ".class";
        // [readClass] 启动类路径
        try {
            return bootstrapClasspath.readClass(className);
        } catch (Exception ignored) {}
        // [readClass] 扩展类路径
        try {
            return extensionClasspath.readClass(className);
        } catch (Exception ignored) {}
        // [readClass] 用户类路径
        return userClassPath.readClass(className);
    }
}
