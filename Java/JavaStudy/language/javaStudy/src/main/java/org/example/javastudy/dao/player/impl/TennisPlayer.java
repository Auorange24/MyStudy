package org.example.javastudy.dao.player.impl;

import org.example.javastudy.dao.player.BasePlayer;
import org.example.javastudy.dao.player.PlayerType;

public class TennisPlayer extends BasePlayer {

    public boolean isTennisPlayer() {
        return getType() == PlayerType.TENNIS;
    }

    @Override
    public void play() {
        System.out.println("play tennis");
    }
}
