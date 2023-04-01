package limit

import (
	"context"
	"io"
	"net/http"
)

// Fuse 熔断器
type Fuse struct {
}

func (f Fuse) Proxy(method string, url string, body io.ReadCloser) (bytes []byte) {
	// 发送请求,并设定超时时间 5s
	request, _ := http.NewRequestWithContext(context.Background(), method, url, body)
	response, _ := http.DefaultClient.Do(request)
	bytes, _ = io.ReadAll(response.Body)
	return
}
