# 总结几种socket粘包的解包方式

## 什么是粘包
在socket缓存中，由于网络阻塞或者服务器等各种原因，tcp数据包在缓存中头尾相接，这种现象就是粘包。

![Alt 粘包图片](https://imgs.developpaper.com/imgs/1606616-20200408144442234-1542519195.jpg "粘包图片")

## 解决方式
### **fixed length**
***
**解释：**
客户端和服务端发送到缓冲区的数据长度是固定的，如果消息的长度小于固定长度，那消息与消息之间有一定的空白区域，会浪费一定资源。如果消息的长度大于固定长度，则消息需要拆包发送，会出现半包现象。\
**应用：**
Netty io.netty.handler.codec.FixedLengthFrameDecoder
***
### **delimiter based**
**解释：**
每则消息通过分割符断开，例如每则消息的末尾约定为`'\n'`，这样可以通过扫描字节流，如果遇到了分割符则将消息两次分割符之间的字节流合并解包，作为一次请求处理。但是，这种方式由于需要逐字节扫描字节流，所以对于长字节流来说，性能消耗严重，所以比较适合短字节流的请求。\
**应用：**
Netty io.netty.handler.codec.DelimiterBasedFrameDecoder
***
### **length field based**
**解释：**
在协议头里添加一个字符流长度的字段，这样每次接收到字节流请求，先解析字节流头部的协议，确定整个字节流的长度，根据具体的长度来读取完整的字节流来确定一次完整的消息请求。\
**应用：**
Netty io.netty.handler.codec.LengthFieldBasedFrameDecoder
