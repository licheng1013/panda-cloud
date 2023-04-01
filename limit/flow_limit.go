package limit

import (
	"github.com/gin-gonic/gin"
)

// FlowLimit 限流器
type FlowLimit struct {
}

func (f FlowLimit) Control(info RequestInfo) (trigger bool, data any) {
	//log.Println("限流处理 -> ", info)
	return false, gin.H{"message": "限流处理"}
}
