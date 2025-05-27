package org.example.javastudy.Object_Oriented.InnerClass;


import lombok.Data;

@Data
public class OuterClass {

    public int id;

    public class InnerClass {
        public String name;

        public void print_id() {
            System.out.println("id = " + id);
        }
    }

    public InnerClass getInnerClass() {
        return new InnerClass();
    }
}
