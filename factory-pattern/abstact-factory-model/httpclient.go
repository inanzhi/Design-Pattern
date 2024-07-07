package abstact_factory_model

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Doer 操作者
type Doer interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

//Client结构体实现了Doer接口中的唯一方法，于是可以将Client结构体返回给Doer接口
//type Client struct {
//}
//func (c *Client) Do(req *http.Request) (*http.Response, error) {
//	return c.do(req)
//}

// NewHTTPClient 返回一个net/http包提供的HTTPClient
// 注意 我们这里返回示例而不是指针，因为只是利用对象调用方法，而不是修改实例
// 但是如果想要修改原实例，同样不穿指针，而使用SetXXX,，防止被意外修改
func NewHTTPClient() Doer {
	return &http.Client{}
}

// 定义新的结构体，来使用同一接口，以便于可以相互替换，避免真实接口失败。
type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (resp *http.Response, err error) {
	res := httptest.NewRecorder()
	return res.Result(), nil
}

func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}

// QueryUser 正常请求方法 传一个操作者进去
func QueryUser(doer Doer) error {
	//一个整体丢进Do的参数之中   res Do(req)
	req, err := http.NewRequest(http.MethodGet, "http://iam.api.marmotedu.com:8080/v1/secrets", nil)
	if err != nil {
		return err
	}

	res, err := doer.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("res:", res)
	return nil
}

// TestQueryUser 用测试用例来测试正常的查询方法
func TestQueryUser(t *testing.T) {
	doer := NewMockHTTPClient()
	if err := QueryUser(doer); err != nil {
		t.Error(err)
	}
}

// 正常使用该方法
func use() error {
	doer := NewHTTPClient()
	if err := QueryUser(doer); err != nil {
		return err
	}
	return nil
}
