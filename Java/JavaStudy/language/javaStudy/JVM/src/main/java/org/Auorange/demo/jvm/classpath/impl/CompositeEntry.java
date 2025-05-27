package org.Auorange.demo.jvm.classpath.impl;

import org.Auorange.demo.jvm.classpath.Entry;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * @Author hideoncode
 * @Date 2025/5/22 09:28
 */

/*
*
* */
public class CompositeEntry implements Entry {

    private final List<Entry> entryList = new ArrayList<>();

    /*
    * File.pathSeparator 路径分隔符
    * Windows的路径分隔符 \
    * Linux的路径分隔符 /
    * */

    public CompositeEntry(String pathList) {
        String []paths = pathList.split(File.pathSeparator);
        for (String path : paths) {
            entryList.add(Entry.Create(path));
        }
    }

    @Override
    public byte[] readClass(String className) throws IOException {
        for (Entry entry : entryList) {
            try {
                return entry.readClass(className);
            } catch (Exception ignored) {

            }
        }
        throw new IOException("class not found + " + className);
    }

    @Override
    public String toString() {
        String[] strs = new String[entryList.size()];
        for (int i = 0; i < entryList.size(); i++) {
            strs[i] = entryList.get(i).toString();
        }
        return String.join(File.pathSeparator, strs);
    }
}
