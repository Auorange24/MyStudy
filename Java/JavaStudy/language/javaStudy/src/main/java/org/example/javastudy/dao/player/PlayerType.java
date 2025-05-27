package org.example.javastudy.dao.player;

public enum PlayerType {

    TENNIS,
    FOOTBALL,
    BASKETBALL;

    private String name;

    public PlayerType(String name) {
        this.name = name;
    }
}
