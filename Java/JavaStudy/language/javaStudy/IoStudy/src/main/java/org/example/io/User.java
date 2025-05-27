package org.example.io;

import lombok.Data;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.*;

/**
 * @Author HideOnLife
 * @Date 2025/5/27 16:04
 */

@Data
public class User implements Serializable {
    private static final Logger log = LoggerFactory.getLogger(User.class);
    private String name;
    private int age;


    public static void main(String[] args) {
        User user = new User();
        user.setName("test");
        user.setAge(18);
        System.out.println(user);

        // write object into file
        try (ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("io_test"))) {
            oos.writeObject(user);
        } catch (IOException e) {
            log.info(e.getMessage());
        }

        // Read Object From File
        try (ObjectInputStream ois = new ObjectInputStream(new FileInputStream("io_test"))) {
            User user1 = (User) ois.readObject();
            System.out.println(user1);
        } catch (IOException | ClassNotFoundException e) {
            log.info(e.getMessage());
        }
    }
}
