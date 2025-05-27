package org.Auorange.demo.jvm.classpath.impl;

import org.Auorange.demo.jvm.classpath.Entry;

import java.io.IOException;
import java.nio.file.*;

/**
 * @Author hideoncode
 * @Date 2025/5/22 11:23
 */
public class ZipEntry implements Entry {

    private Path absolutePath;


    public ZipEntry(String path) {
        /*
        * get(path) 得到相对路径
        * toAbsolutePath 通过相对路径得到绝对路径
        * */
        this.absolutePath = Paths.get(path).toAbsolutePath();
    }

    @Override
    public byte[] readClass(String className) throws IOException {
        try (FileSystem fileSystem = FileSystems.newFileSystem(absolutePath, null)) {
            return Files.readAllBytes(fileSystem.getPath(className));
        }
    }

    @Override
    public String toString() {
        return this.absolutePath.toString();
    }
}
