// 哈希表查找法，构建哈希表的同时进行哈希查找
// 时间复杂度：O(n)  空间复杂度：O(n)

import java.util.*;

public class Solution_3 {

    public int[] twoSum(int[] nums, int target) {
        int[] res = new int[2];

        HashMap<Integer, Integer> map = new HashMap<>();
        for(int i=0; i<nums.length; i++){
            int temp = target - nums[i];
            if(map.containsKey(temp)){
                res[0] = i;
                res[1] = map.get(temp);
                return res;
            }

            map.put(nums[i], i);
        }

        return res;
    }
}