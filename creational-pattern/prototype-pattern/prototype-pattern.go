package prototype_pattern

//https://www.cnblogs.com/amunote/p/15256894.html
//原型模式

//介绍
//优雅地创建对象的拷贝
//将克隆某对象的职责交给了要被克隆的这个对象。被克隆的对象需要提供一个clone()方法。通过这个方法可以返回该对象的拷贝。

//使用场景
//创建新对象的操作比较耗资源（如数据库操作）或代价比较高时。比较起从头创建新对象，克隆对象明显更加可取
//要被克隆的对象创建起来比较复杂时：比如对象克隆的过程中存在深度拷贝或分层拷贝时；

//要被克隆的对象存在无法被直接访问到的私有成员时。

//tips
//深拷贝:既复制对象本身，又递归复制该对象所引用的其他对象，被复制的内部对象完全是新的
//浅拷贝（shallow copy):只复制对象的引用，基本数据类型复制值，引用类型复制引用
