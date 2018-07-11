# go-t2sdkEx
恒生电子周边接口sdk二次封装,go版本



# 说明
此版本基于go实现了T2连接、收发包、解包、订阅等常规业务,但也仅仅是提供了一种用go调用c++版本t2的思路,有其他好方案的同学可以私信给我

# t2sdkEx动态库源码
因为go不能直接调用C++的动态库,需要通过C的方式导出才能调用,所以需要封装C++版本的t2sdk以c的方式导出,源码在
https://github.com/level2player/t2sdkEx


# 恒生t2sdk介绍
英文全称TCP 2 Software Development Kit。T2SDK是接入金融基础件2.0的客户端开发包，因为该开发包底层采用TCP协议，取名为“T2SDK”。T2SDK以dll（Windows下）或so（Linux下）的形式提供给开发者，通过使用T2SDK开发程序可以轻松实现和CRES2.0（金融基础件2.0）平台的对接，使用T2_SDK开发的程序在本文档中也称为“客户端”


