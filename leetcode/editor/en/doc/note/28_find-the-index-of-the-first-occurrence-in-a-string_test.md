用go的話有非常直觀簡單且內建的方法

```go
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
```

要土炮的話可以參考一下`strings.Index(string, string) int`的實作