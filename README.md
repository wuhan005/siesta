# 😴 siesta ![Go](https://github.com/wuhan005/siesta/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/wuhan005/siesta)](https://goreportcard.com/report/github.com/wuhan005/siesta) [![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/wuhan005/siesta)

## 为什么会有这个项目？

最近国家出台了关于限制未成年人游戏时间的相关政策。但很可惜，我正在读小学二年级的弟弟只玩单机的 Minecraft，因此我爸在想能否写个简单的程序来约束他的游戏时间。我便糊了这么一个小玩意出来，原理很简单——超过时间了就杀死包含指定关键词的进程，提醒我弟去休息。

我知道这东西有千千万万种方式可以绕过，但是对于一个小学二年级的学生来说已经足够了。（如果我弟能为了 bypass 这东西开始钻研计算机，那就太好了！）

项目的名字取自英文 <ruby>午睡 <rp>(</rp><rt>siesta</rt><rp>)</rp>，寓意我弟要时刻注意休息，不要沉迷游戏。取这个名字才不是因为最近看了什么番。⁄(⁄ ⁄•⁄ω⁄•⁄ ⁄)⁄

## 开始使用

```bash
# otp 为基于时间的验证密钥
# process 为指定进程的关键词
./siesta.exe --otp <REDACTED> --process "minecraft,java"
```

你可以通过访问 `http://localhost:2830/` 来延长允许游玩的时间。

## License

MIT
