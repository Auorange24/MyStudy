package org.example.file;

import io.reactivex.rxjava3.core.Single;

import java.io.File;
import java.lang.reflect.Field;

/**
 * @Author HideOnLife
 * @Date 2025/5/27 23:45
 */
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
