這題算喇賽題。如果不是靈光一閃，或者剛好以前看過這題目，要臨場想到最佳解是非常困難的。不過這題仍然是有跡可循

### 暴力破

這應該是絕大多數的人的第一印象，直接用一個雙層for把答案兜出來，複雜度是 `O(n^2)`，這裡就不放示範code了，這個都寫不出來的大概玩leetcode也會很痛苦。

### 算出目標再找

這做法滿間接的，同樣是一個迴圈，我們要做兩件事
1. 先算出目標值，以這邊例子來講，若為`for index, value := nums` 來講，我們可以先用`target - value`去找出目標值，再找找看目標值有沒有在nums裡面。
2. 「尋找目標值」這件事情不需要迴圈，我們可以用map做到`O(1)`來找

所以我們可以做一個map，key為值(因為我們要搜尋key)，而value為該key值的index。然後由於每個index僅准許被使用一次，所以我們要排除自己加自己的情形。整組code看起來會像這樣：

```go
func twoSum(nums []int, target int) []int {
	//make it map for sorting
	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i
	}

	for i, v := range nums {

		t := target - v
		if index, exists := numsMap[t]; exists {
			if index == i {
				continue
			}
			return []int{i, index}
		}
	}
	return nil

}
```

應該是有更快的做法，畢竟...

```text
> 2023/12/05 20:55:06	
Success:
	Runtime:9 ms, faster than 40.64% of Go online submissions.
	Memory Usage:5.6 MB, less than 5.64% of Go online submissions.
```