package org.example.javastudy;

import org.junit.jupiter.api.Test;

public class ErrorTest {

    @Test
    public void throwTest() {
        try {
            method(4);
        } catch (RuntimeException e) {
            e.printStackTrace();
        }
    }

    void method(int i) {
        if (i % 2 == 0) {
            throw new ArithmeticException("error_test");
        }
    }
}
