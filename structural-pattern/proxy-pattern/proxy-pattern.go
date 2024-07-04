package proxy_pattern

import "fmt"

// 关注对象之间的通信
// 可以为另一个对象提供一个替身或占位符，以控制对这个对象的访问
// 控制对真实对象的访问权限 外面对象缩小了真实对象的范畴
// 可以记录日志
// 在实际需要的时候才会创建对象，节省资源
// 可以缓存真实对象的结果，避免重复计算
// 不改变真实对象前提下，增加或减少对其的访问

// Image 是接口，定义了 Display 方法
type Image interface {
	Display()
}

// 真实对象
type RealImage struct {
	fileName string
}

// 真实对象的构造函数
func NewRealImage(fileName string) *RealImage {
	fmt.Println("Loading image", fileName)
	return &RealImage{fileName: fileName}
}

func (img *RealImage) Display() {
	fmt.Println("Displaying image:", img.fileName)
}

// 代理对象中包含了真实对象  是过渡了一层
type ProxyImage struct {
	realImage *RealImage
	fileName  string
}

func NewProxyImage(fileName string) *ProxyImage {
	return &ProxyImage{fileName: fileName}
}

// 代理对象去替代真实对象做一件事，其实只是代理对象去调用真实对象的方法，偷梁换柱
func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.fileName)
	}
	p.realImage.Display()
}

//使用代理对象

func main() {
	image := NewProxyImage("test_image.jpg")

	//代理对象
	image.Display()
	fmt.Println("")

	image.Display()
}
