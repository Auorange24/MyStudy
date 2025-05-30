package org.example.lockinterfacestudy;

import java.util.concurrent.locks.AbstractQueuedSynchronizer;

/**
 * @Author HideOnLife
 * @Date 2025/5/30 09:29
 */
public class Mutex {

    private static class Sync extends AbstractQueuedSynchronizer {
        @Override
        protected boolean tryAcquire(int arg) {
            // CAS 操作
            return compareAndSetState(0, 1);
        }

        @Override
        protected boolean tryRelease(int arg) {
            setState(0);
            return true;
        }

        @Override
        protected boolean isHeldExclusively() {
            return getState() == 1;
        }
    }

    private final Sync sync = new Sync();

    public void lock() {
        sync.acquire(1);
    }

    public void unlock() {
        sync.release(1);
    }

    public boolean isLocked() {
        return sync.isHeldExclusively();
    }

    public static void main(String[] args) {
        int arg = 1;
    }
}
