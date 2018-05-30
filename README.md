KCVoice
=======

为舰队Collection游戏实现的一个简单的爬虫。

## 安装

要安装本包，请在终端输入

    go get github.com/Nangcr/kcvoice

默认情况下，KCVoice会使用[萌娘百科](https://zh.moegirl.org)的数据来抓取信息，如果你想添加自定义数据来源，请参考代码中的常量字符串

## 使用

使用此包的任何功能前，请先设定数据来源，通常情况下，你需要在代码中使用以下函数

    s := NewDefaultSource()

然后通过数据来源对象的方法来获得数据，请注意方法的参数类型和返回值类型

## License

KCVoice is licensed under the [MIT License](LICENSE). You are encouraged
to embed KCVoice into your other projects, as long as the license
permits.

You are also encouraged to disclose your improvements to the public, so that
others may benefit from your modification, in the same way you receive benefits
from this project.
