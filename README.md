# stomp（开发阶段）
## 0.概述
STOMP是一个基于帧的协议，它的帧是在HTTP上建模的。一帧由一个命令、一组可选的标题和一个可选的正文组成。STOMP是基于文本的，但也允许传输二进制消息。STOMP的默认编码是UTF-8，但是它支持消息正文的替代编码规范。  

该项目基于WebSocket的子协议stomp golang实现; 协议文档：
目前还处于开发阶段，主要完成stomp服务端的逻辑   

ð欢迎创建issue ！！！

|模块规划|进度|
|----|----|
|stomp服务端|正在进行|
|client客户端||


## 参考资料
[1] [GitHub开源项目go-stomp/stomp](https://github.com/go-stomp/stomp) frame使用该库并做了一些优化  
[2] [spring/websocket-stomp开发文档](https://docs.spring.io/spring-framework/docs/current/reference/html/web.html#websocket-stomp)  
[3] [协议文档stomp_v1.2](http://stomp.github.io/stomp-specification-1.2.html)  