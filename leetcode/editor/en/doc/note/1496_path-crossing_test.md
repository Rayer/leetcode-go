其實這題看題目的圖解就滿好想的了。簡單說，就拿一個set，裝裡面x/y座標，有重複的座標就回true。golang沒有set，所以就用map來代替即可。

唯一要注意的點是，一開始的原點也要算進去(0/0)，題目很好心地在測資就有提示你這一點。

```go
func isPathCrossing(path string) bool {
	hTable := make(map[string]bool)

	curX := 0
	curY := 0
	hTable["0/0"] = true
	for _, p := range path {
		switch p {
		case 'N':
			curX += 1
		case 'S':
			curX -= 1
		case 'E':
			curY += 1
		case 'W':
			curY -= 1
		}
		pos := fmt.Sprintf("%d/%d", curX, curY)
		if _, exists := hTable[pos]; exists {
			return true
		}
		hTable[pos] = true
	}
	return false
}
```

```text
> 2023/12/24 07:53:54	
Success:
	Runtime:2 ms, faster than 50.91% of Go online submissions.
	Memory Usage:2.3 MB, less than 81.82% of Go online submissions.
```