// 哈希表查找法，构建哈希表与哈希表查找分开
// 时间复杂度：O(n)  空间复杂度：O(n)

import java.util.*;

public class Solution_2 {

    public int[] twoSum(int[] nums, int target) {
        int[] res = new int[2]; 

        // 1. 构建哈希表
        HashMap<Integer, Integer> map = new HashMap<>();
        for(int i=0; i<nums.length; i++){
            map.put(nums[i], i);
        }

        // 2. 哈希查找
        //    对于每一个元素'nums[i]',查找是否存在'target-nums[i]',注意不可重复查找同一元素
        for(int i=0; i<nums.length; i++){
            int temp = target - nums[i];
            if(map.containsKey(temp) && map.get(temp)!=i){
                res[0] = i;
                res[1] = map.get(temp);
                return res;
            }
        }

        return res;
    }
}