import java.util.*;
public class Solution {

    public int threeSumClosest(int[] nums, int target) {
        // Sort array
        Arrays.sort(nums);

        int closestSum = nums[0] + nums[1] + nums[2]; // sum structure

        for (int i = 0; i < nums.length; i++) {
            int left = i + 1;
            int right = nums.length - 1;

            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];

                // Update closestSum if closer to the target
                if (Math.abs(target - sum) < Math.abs(target - closestSum)) {
                    closestSum = sum;
                }

                // Move pointers with comparison
                if (sum < target) {
                    left++;
                } else if (sum > target) {
                    right--;
                } else {
                    // return sum if the result exactly equals to target
                    return sum;
                }
            }
        }

        return closestSum;
    }
}