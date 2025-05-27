package org.example.javastudy;

import org.example.javastudy.Object_Oriented.InnerClass.OuterClass;
import org.example.javastudy.Object_Oriented.immutable_class.Book;
import org.example.javastudy.dao.player.impl.BasketballPlayer;
import org.example.javastudy.dao.player.impl.FootBallPlayer;
import org.junit.jupiter.api.Test;

public class ObjectTest {
    @Test
    public void object_test_1() {
        BasketballPlayer basketballPlayer = new BasketballPlayer();
        basketballPlayer.sleep();
        basketballPlayer.play();
        FootBallPlayer footBallPlayer = new FootBallPlayer();
        footBallPlayer.sleep();
        footBallPlayer.play();
    }

    @Test
    public void object_test_2() {
        OuterClass outerClass = new OuterClass();
        OuterClass.InnerClass innerClass = new OuterClass().new InnerClass();
    }

    @Test
    public void object_test_3() {
        Book book = new Book();
        book.setName();
    }
}
