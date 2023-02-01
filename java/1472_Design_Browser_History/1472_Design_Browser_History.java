class BrowserHistory {

    LinkedList<String> history;
    int currentPage = 0;

    public BrowserHistory(String homepage) {
        history = new LinkedList<String>();
        history.addLast(homepage);
    }

    public void visit(String url) {
        if (currentPage < history.size()) {
            while (currentPage+1 < history.size()) {
                history.removeLast();
            }
        }
        currentPage += 1;
        history.addLast(url);
    }
    
    public String back(int steps) {
        currentPage -= steps;
        if (currentPage < 0) {
            currentPage = 0;
        }
        return history.get(currentPage);
    }
    
    public String forward(int steps) {
        currentPage += steps;
        if (currentPage > history.size() - 1) {
            currentPage = history.size() - 1;
        }
        return history.get(currentPage);
    }
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * BrowserHistory obj = new BrowserHistory(homepage);
 * obj.visit(url);
 * String param_2 = obj.back(steps);
 * String param_3 = obj.forward(steps);
 */
