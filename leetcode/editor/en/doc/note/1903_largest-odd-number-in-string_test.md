思路滿簡單的，只要是尾數是奇數的就是奇數，所以只要for迴圈到奇數後取頭到該數字的距離即可加入candidate清單。

1. 取candidate, for迴圈直到偶數
2. 看看目前candidate的位數是否比這個數還大，是的話跳過
3. 如果這個數的位數大於candidate位數，直接把candidate清空後把這個數放入
4. 如果兩者位數相等，append這個數進入candidate

然後在candidate環節，從最大位數一路比大小到最小即可。理論上應該可以放入上個for迴圈，不過code會變得非常複雜。anyway, 成績並非很好。

```text
> 2023/12/07 16:26:05	
Success:
	Runtime:55 ms, faster than 5.19% of Go online submissions.
	Memory Usage:7.6 MB, less than 5.19% of Go online submissions.
```