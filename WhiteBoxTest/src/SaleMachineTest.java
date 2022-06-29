package src;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SaleMachineTest {

    private static SaleMachine machine=new SaleMachine();

    @Test
    public void test1() {
        machine=new SaleMachine();
        String res=machine.operation("Beer", "5C");
        assertEquals("Input Information\n"+
                "Type:Beer;Money:5 Cents; Change:0\n"+
                "\n"+
                "Current State\n"+
                "Beer:                  5\n"+
                "Orange Juice:          6\n"+
                "5 Cents:               7\n"+
                "1 Dollar:              6", res);
    }

    @Test
    public void test2() {
        machine=new SaleMachine();
        String res=machine.operation("OrangeJuice", "5C");
        assertEquals("Input Information\n" +
                "Type: OrangeJuice; Money:5 Cents; Change:0\n" +
                "\n" +
                "Current State\n" +
                "Beer:                  6\n" +
                "Orange Juice:          5\n" +
                "5 Cents:               7\n" +
                "1 Dollar:              6", res);
    }

    @Test
    public void test3() {
        machine=new SaleMachine();
        String res=machine.operation("Beer", "1D");
        assertEquals("Input Information\n" +
                "Type:Beer; Money:1 Dollar; Change:5 Cents\n" +
                "\n" +
                "Current State\n" +
                "Beer:                  5\n" +
                "Orange Juice:          6\n" +
                "5 Cents:               5\n" +
                "1 Dollar:              7",res);
    }

    @Test
    public void test4() {
        machine=new SaleMachine();
        String res=machine.operation("OrangeJuice", "1D");
        System.out.println(res);
        assertEquals("Input Information\n" +
                "Type: OrangeJuice; Money: 1 Dollar; Change: 5 Cents\n" +
                "\n" +
                "Current State\n" +
                "Beer:                  6\n" +
                "Orange Juice:          5\n" +
                "5 Cents:               5\n" +
                "1 Dollar:              7",res);
    }

    @Test
    public void test5() {
        machine=new SaleMachine(0,0,6,6);
        String res=machine.operation("Beer", "1D");
        assertEquals("Failure Information\n" +
                "Change Shortage",res);
    }

    @Test
    public void test6() {
        machine=new SaleMachine(0,0,6,6);
        String res=machine.operation("OrangeJuice", "1D");
        assertEquals("Failure Information\n" +
                "Change Shortage",res);
    }

    @Test
    public void test7() {
        machine=new SaleMachine(6,6,0,6);
        String res=machine.operation("Beer", "1D");
        assertEquals("Failure Information\n" +
                "Beer Shortage",res);
    }

    @Test
    public void test8() {
        machine=new SaleMachine(6,6,6,0);
        String res=machine.operation("OrangeJuice", "5C");
        assertEquals("Failure Information\n" +
                "OrangeJuice Shortage",res);
    }

    @Test
    public void test9() {
        machine=new SaleMachine(6,6,6,6);
        String res=machine.operation("Beer", "5G");
        assertEquals("Failure Information\n" +
                "Money Error",res);
    }

    @Test
    public void test10() {
        machine=new SaleMachine(6,6,6,6);
        String res=machine.operation("Coca-Cola", "1D");
        assertEquals("Failure Information\n" +
                "Type Error",res);
    }

    @Test
    public void test11() {
        machine=new SaleMachine(6,6,0,6);
        String res=machine.operation("Beer", "5C");
        assertEquals("Failure Information\n" +
                "Beer Shortage",res);
    }

    @Test
    public void test12() {
        machine=new SaleMachine(6,6,6,6);
        String res=machine.operation("Coca-Cola", "5C");
        assertEquals("Failure Information\n" +
                "Type Error",res);
    }


    @Test
    public void test13() {
        machine=new SaleMachine(6,6,6,0);
        String res=machine.operation("OrangeJuice", "1D");
        assertEquals("Failure Information\n" +
                "OrangeJuice Shortage",res);
    }

    @Test
    public void test14() {
        machine=new SaleMachine(6,6,-1,6);
        String res=machine.operation("Beer", "1D");
        assertEquals("Failure Information\n" +
                "Beer Shortage",res);
    }


    @Test
    public void test15() {
        machine=new SaleMachine(6,6,6,-1);
        String res=machine.operation("OrangeJuice", "1D");
        assertEquals("Failure Information\n" +
                "OrangeJuice Shortage",res);
    }

}
