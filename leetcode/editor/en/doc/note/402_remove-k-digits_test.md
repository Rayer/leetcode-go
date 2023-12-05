很直覺的就會用DP來解決這問題，然後會iterate所有的digit去找出最小的。 但是測資很雞賊，他會用一個極大的int，大到連int64都裝不下那種 ，所以直接用atoi or itoa是無效的，要使用純文字的方法來解決。

這個比較偏腦筋急轉彎的地方在於，移除數字是有規則的，比方說1236310

規則是找到第一個數字遞減的點，移除該數字前面的數字。假設k=4，那接下來移除順序會是 `1236310 => 126310 => 12310 => 1210 => 110`

### TLE answer

```go
func removeKdigits(num string, k int) string {
	if k == 0 {
		// Remove head 0s
		num = strings.TrimLeft(num, "0")
		if num == "" {
			return "0"
		}
		return num
	}

	if len(num) == 1 {
		return "0"
	}

	r := []rune(num)
	removeOne := false
	for index, _ := range num {
		if index == 0 {
			continue
		}
		if r[index-1] > r[index] {
			r = append(r[0:index-1], r[index:]...)
			removeOne = true
			break
		}
	}
	if removeOne == false {
		r = append(r[0 : len(r)-1])
	}
	return removeKdigits(string(r), k-1)
}
```

仔細想想，似乎沒必要遞迴增加一堆成本啊.... 改用for迴圈就輕易過關了 -_-a 只是過關的滿難看的 

```text
> 2023/12/05 20:18:21	
Success:
	Runtime:2709 ms, faster than 5.48% of Go online submissions.
	Memory Usage:5.1 MB, less than 34.25% of Go online submissions.
```