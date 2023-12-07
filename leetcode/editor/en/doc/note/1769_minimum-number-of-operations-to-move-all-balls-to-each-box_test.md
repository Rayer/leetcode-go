這題麻煩的地方在於，看起來有暴力破以外的解，但是到頭來我還是只發現暴力破的作法。

首先這個其實是個純數學題，我們先用一個`result := make([]int)` 來裝結果，我們對每個box做iteration。假設這個box裡面有球，那移動到其他盒子所需要的步數就為`Abs(i-j)`相當於他們的距離。對每個盒子都算完以後就可以得到答案了。

```text
> 2023/12/07 23:57:43	
Success:
	Runtime:24 ms, faster than 61.11% of Go online submissions.
	Memory Usage:5.1 MB, less than 50.00% of Go online submissions.
```