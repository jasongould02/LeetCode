class Foo {
    private CountDownLatch l;
    public Foo() {
        l = new CountDownLatch(2);
    }

    public void first(Runnable printFirst) throws InterruptedException {
        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();
        l.countDown();
    }

    public void second(Runnable printSecond) throws InterruptedException {
        while(l.getCount() != 1);
        // printSecond.run() outputs "second". Do not change or remove this line.
        printSecond.run();
        l.countDown();
    }

    public void third(Runnable printThird) throws InterruptedException {
        l.await();
        // printThird.run() outputs "third". Do not change or remove this line.
        printThird.run();
    }
}