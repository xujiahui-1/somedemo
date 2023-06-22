gomock是golang官方开发维护的接口级别的mock方案，包含了GoMock包和mockgen工具两部分，其中GoMock包完成对桩对象生命周期的管理，mockgen工具用来生成interface对应的Mock类源文件。要使用gomock的一个前提是模块之间务必通过接口进行依赖，而不是依赖具体实现，否则mock会非常困难。这个工具目前业界用的并不多，主要是局限性太大，所以我们只需要简单了解一下如何使用就行。


安装gomoke
>go get github.com/golang/mock/mockgen@v1.6.0
> 
>go install github.com/golang/mock/mockgen

使用命令生成moke类

> mockgen  -destination mock_user_dao.go -package dao -source user_dao.go 
