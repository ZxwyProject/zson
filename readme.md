## ZxwyProject/zson
### 简介
+ 一个兼容 [Gin标准](https://github.com/gin-gonic/gin/tree/master/internal/json) 的 json库 切换工具

### 使用
+ `go build -tags "go_json" .`
- `"go_json"`: [goccy/go-json](github.com/goccy/go-json)
- `"jsoniter"`: [json-iterator/go](github.com/json-iterator/go)
- `"sonic avx"`: [bytedance/sonic](github.com/bytedance/sonic)

### 注意
+ 部分情况不会自动添加 `avx` tag, 导致无法启用 sonic, 建议手动添加
