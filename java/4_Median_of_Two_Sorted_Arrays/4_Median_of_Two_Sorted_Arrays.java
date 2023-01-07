class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int[] combined = new int[nums1.length + nums2.length];
        
        System.arraycopy(nums1, 0, combined, 0, nums1.length);
        System.arraycopy(nums2, 0, combined, nums1.length, nums2.length);
        
        Arrays.sort(combined);
        int mid = combined.length / 2;
        
        if (combined.length % 2 == 1) {
            return (double) combined[mid];
        } else {
            return (double) (combined[mid-1] + combined[mid]) / 2.0;
        }
        
    }
}
