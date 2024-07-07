package adapter_pattern

// Adapter 是适配器，包含一个 SpecificRequest 的引用
type Adapter struct {
	specificRequest SpecificRequest
}

func NewAdapter(request SpecificRequest) Adapter {
	return Adapter{request}
}

// 现有方法中间接调用新方法
// 实现 Request 接口的 Request 方法
func (a *Adapter) Request() string {
	return a.specificRequest.SpecificRequest()
}
