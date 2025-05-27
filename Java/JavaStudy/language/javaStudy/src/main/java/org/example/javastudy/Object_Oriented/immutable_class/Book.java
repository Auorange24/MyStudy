package org.example.javastudy.Object_Oriented.immutable_class;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class Book {

    private String name;
    private int price;

    // Object的默认方法包含toString
    @Override
    public String toString() {
        return "Book{" +
                "name='" + name + '\'' +
                ", price=" + price +
                '}';
    }
}
