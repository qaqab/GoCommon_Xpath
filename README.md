# GoCommon_Xpath

`GoCommon_Xpath` 是一个基于 `antchfx/htmlquery` 的 Go 包，用于在 HTML 文档中使用 XPath 查询提取信息。这个包旨在简化在 Go 语言中使用 XPath 提取 HTML 内容的过程。

## 背景

HTML 是一种用于构建网页的标记语言。在处理网页数据时，常常需要从 HTML 文档中提取特定的信息。XPath 是一种用于在 XML 文档中查找信息的语言，并且同样适用于 HTML 文档。`GoCommon_Xpath` 包封装了 `antchfx/htmlquery` 提供的功能，提供了一些便捷的方法来执行 XPath 查询和提取信息。

## 特性

- 使用 XPath 查询提取 HTML 节点列表。
- 使用 XPath 查询提取单个 HTML 节点。
- 提取节点的文本内容。

## 安装

使用 `go get` 下载并安装 `GoCommon_Xpath` 包：

```sh
go get github.com/yourusername/GoCommon_Xpath
