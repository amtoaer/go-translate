## 这是什么？

这是一个简易的命令行翻译工具，使用 Golang 编写。

## 如何运行？

1. 下载 release 中编译好的二进制文件。
2. 在有 Golang 的开发环境的情况下：
   ```bash
   git clone https://github.com/amtoaer/go-translate
   cd go-translate
   go build
   ```

## 使用方法？

```bash
./go-translate '你要翻译的内容' from '语言' to '语言' via '翻译方法'
# 或采用默认翻译方法
./go-translate '你要翻译的内容' from '语言' to '语言'
```

在当前版本中:

- 支持的语言：`en`/`zh`
- 支持的翻译方法：`niu`/`youdao`

## 覆盖 token？

程序默认使用代码中保存的各类`token`/`apikey`/`secret`运行，然而，这些都是开发者以个人名义申请的，在使用人数增多时很可能会受到限制。推荐用户自行申请后覆盖默认配置。可配置项的列表见`config/config.go`中的`Config`结构体（以当前版本为例）：

```go
type Config struct {
	BaiduToken   string
	NiuToken     string
	YoudaoID     string
	YoudaoSecret string
}
```

用户配置时，请将配置文件编辑为`json`格式，并保存为`~/.gotrans`。配置支持按需填写，`json`字段需与变量名严格一致。

如对于偏好小牛翻译的用户，他们仅仅需要配置`NiuToken`一项即可，配置文件可以写为：

```json
{
  "NiuToken": "申请到的token"
}
```

每次翻译时记得在最后指定翻译方法（`via niu`）即可。

## 修改默认？

当前默认翻译方法为`youdao`，暂不支持修改（不过支持所需工作量也不大，应该会很快支持）。

## 参与贡献？

该程序设计十分简单，核心思路是：

1. 每个翻译来源都需要实现`translator/interface.go`中的`Translator`接口
2. 使用`Mapper`做映射，统一该程序接受的`lang`格式

对于翻译来源的贡献者：

1. 在`translator`下新建对应文件，实现接口
2. 在`config/config.go`内的 Config 结构体中加入默认值和可配置字段
3. 在`translator/interface`添加翻译方法和方法的构造
4. 在`main.go`中加入翻译方法入口
