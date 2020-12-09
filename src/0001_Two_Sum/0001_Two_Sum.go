/*
解题思路：使用 map 缓存 slice，然后利用 map 查找对应数据
*/

package _001_Two_Sum

func twoSum(nums []int, target int) []int {
	// map[num]index
	tmp := make(map[int]int)

	for index := range nums {
		if i, ok := tmp[target-nums[index]]; ok {
			return []int{i, index} // 返回index时调整顺序
		}

		// 将存入 map 的动作放在后面，
		// 防止 nums 中有相同数字，覆盖了之前的 index
		tmp[nums[index]] = index
	}
	return nil
}
