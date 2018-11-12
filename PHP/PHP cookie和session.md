## PHP cookie和session

setcookie函数  
描述：设置cookie  
语法：bool setcookie(string $name)
### cookie
cookie的作用  
1. 会话状态管理
2. 个性化设置
3. 浏览器行为跟踪


### session
为什么使用session  
1. HTTP是一种无状态的协议，
2. SESSION提供在PHP脚本中定义全局变量的方法，使得这个全局变量在同一个SESSION中对于所有的PHP脚本文件内都有效。所以，SESSION是基于http服务器的用于保持状态的方法。
3. SESSION允许通过将数据存储在http服务器中，以在整个用户会话过程中保持该数据，所以SESSION不仅是一个时间的概念，还包括了特定的用户和服务器


#### 与session有关的函数  
```
session_start  
描述：启动新会话或者重用现有会话  
语法：bool session_start([array $options=[]])

$options参数是一个关联数组，如果提供的话，则会用其中的项目覆盖“会话配置”中的配置选项  
如果通过get或者post方式，或者使用cookie提交了会话id，则会重用现有会话  


session_id
获取/设置当前会话id
session_name
读取/设置会话名称

session_destroy
销毁一个会话中的全部数据

```

