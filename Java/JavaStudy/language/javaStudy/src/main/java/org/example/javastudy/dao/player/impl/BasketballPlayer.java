package org.example.javastudy.dao.player.impl;

import org.example.javastudy.dao.player.BasePlayer;
import org.example.javastudy.dao.player.PlayerType;

public class BasketballPlayer extends BasePlayer {

    public boolean isBasketBallPlayer() {
        return getType() == PlayerType.BASKETBALL;
    }

    @Override
    public void play() {
        System.out.println("打篮球");
    }
}
