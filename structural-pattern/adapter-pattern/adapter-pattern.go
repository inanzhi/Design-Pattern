package adapter_pattern

//适配器模式（Adapter Pattern）是一种结构型设计模式，它允许将一个类的接口转换为客户希望的另一个接口，
//从而使原本由于接口不兼容而不能一起工作的类可以协同工作。
//在 Go 语言中，可以使用接口和组合来实现适配器模式。

//假设我们有一个现有的系统，它使用了一种接口，但我们需要集成一个新的库或服务，新的库接口不兼容于现有系统。
//通过适配器模式，我们可以创建一个适配器，将新库的接口转换为现有系统需要的接口。
//新库转化为现有系统需要的    新的来转化为我熟悉的   我出机制来适配别人

//适配器实现了现有系统需要的接口，并包含一个新库对象的引用。它在内部调用新库的方法，并将结果转换为现有系统需要的格式。
