## ZxwyProject/zson
### 简介 Introduction:
+ 一个兼容 [Gin 标准](https://github.com/gin-gonic/gin/tree/master/internal/json) 的 json 库 切换工具
+ A json library switching tool compatible with [Gin](https://github.com/gin-gonic/gin/tree/master/internal/json)

### 使用 Usage:
- 下载 Installation: `go get -u github.com/ZxwyProject/zson`
+ 构建 Build: `go build -tags "go_json" .`
- 标签 Tags:
- `"go_json"`: [`github.com/goccy/go-json`](https://github.com/goccy/go-json)
- `"jsoniter"`: [`github.com/json-iterator/go`](https://github.com/json-iterator/go)
- `"sonic avx"`: [`github.com/bytedance/sonic`](https://github.com/bytedance/sonic)

### 注意 Note:
+ 使用 sonic 时记得额外添加 `avx` tag
+ Use sonic Remember to add the `avx` tag
