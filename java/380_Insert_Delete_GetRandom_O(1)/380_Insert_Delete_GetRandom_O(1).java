class RandomizedSet {
    HashSet<Integer> set;
    public RandomizedSet() {
        this.set = new HashSet<Integer>();
    }
    
    public boolean insert(int val) {
        return set.add(val);
    }
    
    public boolean remove(int val) {
        return set.remove(val);
    }
    
    public int getRandom() {
        int random = (int) (Math.random() * set.size() + 1);
        List<Integer> list = new ArrayList<Integer>(set);
        return list.get(random-1);
    }
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet obj = new RandomizedSet();
 * boolean param_1 = obj.insert(val);
 * boolean param_2 = obj.remove(val);
 * int param_3 = obj.getRandom();
 */
