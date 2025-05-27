package org.Auorange.demo.jvm;

import com.beust.jcommander.JCommander;
import com.beust.jcommander.Parameter;
import java.util.List;

public class Cmd {

    @Parameter(names = {"-?", "-help"}, description = "print hello message", order = 3, help = true)
    boolean helpFlag = false;

    @Parameter(names = "-version", description =  "print version and exit", order = 2)
    boolean versionFlag = false;

    @Parameter(names = {"-cp", "-classpath"}, description = "classpath", order = 1)
    String classpath;

    // 接收参数
    @Parameter(description = "main class and args")
    List<String> mainClassAndArgs;

    boolean ok;

    String getMainClass() {
        return mainClassAndArgs != null && !mainClassAndArgs.isEmpty()
                ? mainClassAndArgs.get(0) : null;
    }

    List<String> getAppArgs() {
        // 将jvm命令删掉？
        return mainClassAndArgs != null && mainClassAndArgs.size() > 1
                ? mainClassAndArgs.subList(1, mainClassAndArgs.size()) : null;
    }

    static Cmd parse(String[] args) {
        // 通过JCommander命令将终端的命令解析到cmd中
        Cmd cmd = new Cmd();
        JCommander jCommander = JCommander.newBuilder()
                .addObject(cmd)
                .build();
        jCommander.parse(args);
        cmd.ok = true;
        return cmd;
    }
    
}
