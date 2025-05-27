package org.example.javastudy.Object_Oriented.immutable_class;

import lombok.Getter;

@Getter
public final class Writer {

    private final String name;
    private final int age;
    private final Book book;

    public Writer(String name, int age, Book book) {
        this.name = name;
        this.age = age;
        this.book = book;
    }
}
