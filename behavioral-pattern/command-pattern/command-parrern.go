package command_pattern

//命令模式

//它建议将请求封装为一个独立的对象。在这个对象里包含请求相关的全部信息，因此可以将其独立执行。
//在命令模式中有如下基础组件：

//命令接口（Command Interface）：声明执行命令的方法。
//具体命令（Concrete Command）：实现命令接口，定义具体的操作。
//接收者（Receiver）：实际执行命令的对象。
//调用者（Invoker）：(请求执行命令)的对象。
//客户端（Client）：创建具体命令对象并设置其接收者。

//优点
//解耦：将请求发送者和请求处理者解耦。
//可扩展：可以很容易地添加新命令。
//支持撤销和重做：通过存储命令的历史记录，可以实现操作的撤销和重做。
//组合命令：可以将多个命令组合成一个复合命令。

//缺点
//可能会导致类的膨胀：每个具体命令都需要一个类，可能会导致类的数量急剧增加。
//命令对象的生命周期管理：需要考虑命令对象的创建和销毁。
