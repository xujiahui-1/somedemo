# COBAR
- cobra包提供以下功能：
- 轻松创建基于子命令的 CLI：如app server、app fetch等。
- 自动添加-h,--help等帮助性Flag
- 自动生成命令和Flag的帮助信息
- 创建完全符合 POSIX 的Flag(标志)（包括长、短版本）
- 支持嵌套子命令
- 支持全局、本地和级联Flag
- 智能建议（app srver... did you mean app server?）
- 为应用程序自动生成 shell 自动完成功能（bash、zsh、fish、powershell）
- 为应用程序自动生成man page
-  命令别名，可以在不破坏原有名称的情况下进行更改
- 支持灵活自定义help、usege等。
- 无缝集成viper构建12-factor应用

## 总的来说， 构建命令行程序

安装
>go get -u github.com/spf13/cobra@latest
> 
> go install github.com/spf13/cobra-cli@latest


初始化
>cobra-cli  init


向cobra添加其他命令
>cobra-cli add wget
> 
>cobra-cli add ping 