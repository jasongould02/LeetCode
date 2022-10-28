class FooBar {
    private int n;
    private volatile int l = 0;
    public FooBar(int n) {
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        
        for (int i = 0; i < n; i++) {
            
        	// printFoo.run() outputs "foo". Do not change or remove this line.
        	printFoo.run();
            l = 1;
            while(l == 1);
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        
        for (int i = 0; i < n; i++) {
            while(l != 1);
            // printBar.run() outputs "bar". Do not change or remove this line.
        	printBar.run();
            l = 0;
        }
    }
}
