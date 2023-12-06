這題其實就是單純的複雜流程題，屬於比較有可能臨場想得出來那種，但是沒debugger的話有些地方不太容易抓到問題。

這題其實做起來滿簡單的：

- 兩個node相加放到一個node，把這個node鏈起來
  - 有in-place的做法（不額外建立新的node），不過得先把兩組list跑過一次，找出哪組比較長
- 全部加完以後把這個儲存結果的node從頭到尾再跑一次計算進位
- 要特別注意最後一個node(`node.Next == nil`)也要進位，要新添加一個node接在後面。

算是medium裡面滿簡單的一道題。

為了灌測資方便，建議做一個函數`GenerateData(data ...int)`，實作可參考這裡面的範例。

```text
> 2023/12/07 01:23:57	
Success:
	Runtime:2 ms, faster than 94.16% of Go online submissions.
	Memory Usage:4.4 MB, less than 62.22% of Go online submissions.
```