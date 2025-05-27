package org.Auorange.demo.jvm.classpath;

import java.io.File;
import java.io.IOException;

/**
 * @Author hideoncode
 * @Date 2025/5/20 22:46
 */

/*
* 每个类的抽象吗
* */
public interface Entry {

    // 读取类中的字节码
    byte[] readClass(String className) throws IOException;

    static Entry Create(String path) {
        return null;
    }
}
