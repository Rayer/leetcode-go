# LeetCodeProjects

 這個repository主要用來存放一些leetcode的解題過程跟心得。IDE使用JetBrains的GoLand，使用的plugin為[LeetCode Editor](https://plugins.jetbrains.com/plugin/12132-leetcode-editor)，感謝**shuzijun**開發這個那麼方便的東西。
 
## 設定

這邊的題目都以[leetcode.com](https://leetcode.com)為主。

### 下載Plugin

請從`preferences->plugins`點選`marketplace`頁籤後尋找`leetcode editor`安裝。同作者有一個`leetcode editor pro`的plugin，功能似乎完全相同，只是是付費。如果有條件的朋友可以安裝這個贊助一下作者。

### 設定Plugin

先從下拉式選單選擇**leetcode.com**(預設應該是leetcode.cn)，填入帳號密碼，語言選擇go，然後建議以下設定

- TempFilePath: 選擇你leetcode的專案根目錄，如果你是從git clone這個專案的話，就是clone近來的目錄。
- 點選Custom Template
- 我個人使用的template會貼在最下面，建議使用這個template。
  - 這個template有提供可以直接執行的test method，通常在ide上該函數左邊都有個▶️可以直接按下去。
  - 有regular test跟table test
  - 可以參考`1903_largest-odd-number-in-string_test.go`看看如何使用table test

### 無法使用testify?

某些版本的GoLand預設新專案沒有開啟go mod支援（真的嚇到我了），所以即使`go mod init leetcode`產出go.mod，然後`go get github.com/stretchr/testify` ，GoLand仍然無法吃到。

- 打開Preference
- 打開Go
- 點入Go Module
- 打開**Enable Go Module Integration**

按照正常流程`go get github.com/stretchr/testify` 應該就能用了。

### Custom Template參考

```text
package leetcode
import "testing"
${question.content}

${question.code}

func Test_$!velocityTool.camelCaseName(${question.titleSlug})(t *testing.T) {

}

func Test_$!velocityTool.camelCaseName(${question.titleSlug})_table(t *testing.T) {
    type args struct {
    }
    
    tests := []struct {
        name string
        args args
        expected interface{}
    }{
        {
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
        
        })
    }
    
}
```