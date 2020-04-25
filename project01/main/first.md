# go

## 执行流程

```
     go build            运行
.go  --------->  .exe   -------> 结果

     go run
.go  --------->  结果
```

### 区别

1. 编译后生成的可执行文件，可以直接拷贝其他机器上运行（未安装go sdk）。
2. 直接使用go run运行，要在另一台机器上运行需要安装go sdk。
3. 编译时，编译器会将程序运行依赖的库文件包含在可执行文件中。


## 规范

1. 优先使用行注释
2. 缩进
    gofmt -w 文件名
3. 行长不要超过80字符

## 官方编程指南

https://golang.org/

## api文档

api：application program interface

golang中文网： https://studygolang.com/pkgdoc
