class Solution {
    public List<String> topKFrequent(String[] words, int k) {
        HashMap<String, Integer> listing = new HashMap<String, Integer>();
        for (String s : words) {
            listing.merge(s, 1, Integer::sum);
        }

        ArrayList<String> result = new ArrayList(listing.keySet());
        Comparator<String> comparator = new Comparator<String>() {
            @Override
            public int compare(String k1, String k2) {
                if (listing.get(k1).equals(listing.get(k2))) {
                    return k1.compareTo(k2);
                } else {
                    return listing.get(k2) - listing.get(k1);
                }
            }
        };
        
        Collections.sort(result, comparator);
        return result.subList(0, k);
    }
}
