class Solution {
    public int[] relativeSortArray(int[] arr1, int[] arr2) {
        HashMap<Integer, Integer> a1Count = new HashMap<Integer, Integer>();
        for(int i = 0; i < arr1.length; i++) {
            a1Count.merge(arr1[i], 1, Integer::sum);
        }
        int position = 0;
        for(int i = 0; i < arr2.length; i++) {
            int dig = arr2[i];
            int amount = a1Count.get(dig);
            
            while(amount > 0) {
                arr1[position] = dig;
                position += 1;
                amount -= 1;
            }
            a1Count.remove(dig);
        }
        
        List<Integer> sortedKeys = new ArrayList<Integer>(a1Count.keySet());
        Collections.sort(sortedKeys);
        
        for(Integer a : sortedKeys) {
            int dig = a;
            int amount = a1Count.get(dig);
            
            while(amount > 0) {
                arr1[position] = dig;
                position += 1;
                amount -= 1;
            }
            a1Count.remove(dig);
        }
        return arr1;
    }
}