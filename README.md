### 需求

- 发送/接收
- 实现群聊
- 高并发= 单机最好+分布式+弹性扩容

#### 实现资源标准化编码
- 资源信息采集并标准化，转化成content/url
- 中级编码，终极目标都是拼接一个消息体（json/xml）

#### 结构

- url：文字，表情包，图片语音-->上传服务器
- url 发送到server

#### 保证消息的可扩展性

- 兼容基础媒介如图片文字语音，（url/pic/content/num）
- 能承载大量新的业务，扩展不能对现有业务产生影响
- 红包/打卡/签到等本质上是消息内容不一样

#### 消息结构

```go
type Message struct{
    Id      int64   `json:`"id,omitempty"       form: "id"      //消息id                  
    Userid  int64   `json:`"userid,omitempty"   from: "userid"  //发送用户id
    Cmd     int     `json:`"cmd,omitempty"      from: "cmd"     //群聊还是私聊
    Dstid   int64   `json:`"dstid,omitempty"    from: "dstid"   //对端（用户）id/群id 
    Medis   int     `json:`"media,omitempty"    from: "media"   //消息样式
    Content string  `json:`"content,omitempty"  from: "content" //消息内容
    Pic     string  `json:`"pic,omitempty"      from: "pic"     //预览图片
    Url     string  `json:`"url,omitempty"      from: "url"     //服务url
    Meno    string  `json:`"memo,omitempty"     from: "memo"    //简单描述
    Amount  int     `json:`"amount,omitempty"   from: "amount"  //和数字相关
}
```

#### 服务器负载分析

- A发图512k
- 100人在线群同时接收到512kb * 100 = 1024kb * 50 =50mb
- 1024个群50*1024=50g

#### 解决方案:
- 缩略图，渲染速度和下载速度 
- 提高资源服务并发能力使用云服务（qos/alioss）,100ms以内
- 压缩消息体，发送文件路径而不是整个文件

#### 高并发

- 单机并发性能优化
- 海量用户分布式部署
- 应对突发事件弹性扩容
