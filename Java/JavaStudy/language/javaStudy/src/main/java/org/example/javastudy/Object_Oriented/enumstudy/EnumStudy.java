package org.example.javastudy.Object_Oriented.enumstudy;

import java.util.EnumSet;

public class EnumStudy {

    public enum PlayerType {
        TENNIS,
        FOOTBALL,
        BASKETBALL
    }

    public static void main(String[] args) {
        EnumSet<PlayerType> playerTypesAllOf = EnumSet.allOf(PlayerType.class);
        System.out.println(playerTypesAllOf);

        E
    }

}
