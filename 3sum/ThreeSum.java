import java.util.*;

public class ThreeSum {

    public static void main(String[] args) {
        Solution solution = new Solution();
        // Example input array
        int[] nums = {-1, 0, 1, 2, -1, -4};

        // Call the threeSum function
        List<List<Integer>> result = solution.threeSum(nums);

        // Print the result
        for (List<Integer> triplet : result) {
            System.out.println(triplet);
        }
    }


}

