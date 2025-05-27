package org.example.javastudy;

import org.junit.jupiter.api.Test;

import java.util.*;

public class MapTest {

    @Test
    public void list_test_1() {
        ArrayList<String> list = new ArrayList<>();

    }

    @Test
    public void linked_list_test_1() {
        List<String> list = new LinkedList<>();
        List<String> list1 = new LinkedList<>();
        List<Integer> list2 = new LinkedList<>();
        System.out.println();
    }

    @Test
    public void map_test_1() {
        Map<Integer, Integer> map = new HashMap<>();
        map.put(1, 1);
        map.put(2, 2);
        map.put(3, 4);
        map.get(1);
        map.remove(1);
        for (Integer i : map.values()) {
            System.out.println(i);
        }

        for (Integer i : map.keySet()) {
            System.out.println(i);
        }
    }
}
