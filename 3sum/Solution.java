import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Solution {

    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums); // Sort the array
        List<List<Integer>> result = new ArrayList<>();

        for (int i = 0; i < nums.length - 2; i++) {
            // Skip duplicates for array
            if (i > 0 && nums[i] == nums[i - 1])
                continue;

            int target = -nums[i];
            int left = i + 1;
            int right = nums.length - 1;

            while (left < right) {
                int sum = nums[left] + nums[right];
                if (sum == target) {
                    // Found triplet
                    result.add(Arrays.asList(nums[i], nums[left], nums[right]));

                    // Skip duplicates for left num and right num
                    while (left < right && nums[left] == nums[left + 1])
                        left++;
                    while (left < right && nums[right] == nums[right - 1])
                        right--;

                    // Move pointers
                    left++;
                    right--;
                } else if (sum < target) {
                    left++; // Move left pointer to the right (increment)
                } else {
                    right--; // Move right pointer to the left (decrement)
                }
            }
        }

        return result;
    }
}
