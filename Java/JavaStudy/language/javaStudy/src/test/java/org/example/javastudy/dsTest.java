package org.example.javastudy;

import lombok.Data;
import lombok.Getter;
import org.example.javastudy.dao.User;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.Scanner;

public class dsTest {

    @Test
    public void test() {
        Integer[] arr = new Integer[10];
        System.out.println(arr.getClass().getName());
    }

    @Test
    public void copy_Array_test_2() {
        int arr1[] = {1, 2, 3, 4, 5};
        int arr2[] = {6, 7, 8, 9, 10};
        int merge_arr[] = new int[arr1.length + arr2.length];
        System.arraycopy(arr1, 0, merge_arr, 0, arr1.length);
        System.arraycopy(arr2, 0, merge_arr, arr1.length, arr2.length);
        for (int element : merge_arr) {
            System.out.println(element);
        }
    }

    @Test
    public void copy_Array_test_1() {
        int arr1[] = {1, 2, 3, 4, 5};
        int arr2[] = {6, 7, 8, 9, 10};
        int merge_arr[] = new int[arr1.length + arr2.length];
        int index = 0;
        for (int element : arr1) {
            merge_arr[index++] = element;
        }
        for (int element : arr2) {
            merge_arr[index++] = element;
        }
        for (int element : merge_arr) {
            System.out.println(element);
        }
    }

    @Test
    public void sort_array_test_1() {
        int arr[] = {1, 4, 5, 2, 3};
        Arrays.sort(arr);
        for (int element : arr) {
            System.out.println(element);
        }
    }

    @Test
    public void two_dimensional_array_test_1() {
        int[][] arr = {{1, 2}, {3, 4}};

    }

    @Test
    public void two_dimensional_array_test_2() {
        Scanner scanner = new Scanner(System.in);
        int level = scanner.nextInt();
        int[][] arr = new int[level][level];

        arr[0][0] = 1;
        for (int i = 1 ; i < level ; i ++) {
            for (int j = 0 ; j <= i ; j ++) {
                arr[i][j] = arr[i - 1][j] + arr[i - 1][j - 1];
            }
        }
    }

    @Test
    public void string_test_1() {
        String s1 = "str1";
        String s2 = "str2";
        System.out.println(s1);
        System.out.println(s2);
        System.out.println(s1.equals(s2));
    }




}
