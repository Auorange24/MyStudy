package org.example.javastudy.dao.player.impl;

import org.example.javastudy.dao.player.BasePlayer;
import org.example.javastudy.dao.player.PlayerType;

public class FootBallPlayer extends BasePlayer {

    public boolean isFootBallPlayer() {
        return getType() == PlayerType.FOOTBALL;
    }

    @Override
    public void play() {
        System.out.println("踢足球");
    }
}
