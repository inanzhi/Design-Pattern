package bridge_pattern

// Color 是颜色接口
type Color interface {
	Paint() string
}

// Red 是 Color 的具体实现
type Red struct{}

func (r *Red) Paint() string {
	return "Red"
}

// Blue 是 Color 的具体实现
type Blue struct{}

func (b *Blue) Paint() string {
	return "Blue"
}
