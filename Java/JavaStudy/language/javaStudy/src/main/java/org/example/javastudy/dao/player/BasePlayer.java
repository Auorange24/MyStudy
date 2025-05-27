package org.example.javastudy.dao.player;

import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public abstract class BasePlayer {

    public PlayerType type;

    public abstract void play();

    public void sleep() {
        System.out.println("player need to sleep");
    }
}
