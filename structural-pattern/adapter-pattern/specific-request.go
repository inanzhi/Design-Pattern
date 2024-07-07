package adapter_pattern

// SpecificRequest 是新库中的类，接口不兼容现有系统
type SpecificRequest struct{}

func (s SpecificRequest) SpecificRequest() string {
	return "Specific request"
}
