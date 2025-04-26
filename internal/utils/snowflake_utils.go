package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

// GenerateID 生成雪花算法ID
func GenerateID() (int64, error) {
	once.Do(func() {
		var err error
		node, err = snowflake.NewNode(1)
		if err != nil {
			fmt.Printf("初始化雪花算法节点失败: %v", err)
		}
	})

	switch {
	case node != nil:
		return node.Generate().Int64(), nil

	default:
		// 雪花格式: 41位时间戳 + 10位节点ID + 12位序列号
		// 标准雪花纪元，节点ID 1，序列号使用当前纳秒的低12位
		ts := time.Now().UnixMilli() - 1288834974657
		nodeID := int64(1)
		seq := time.Now().UnixNano() & 0xFFF

		return (ts << 22) | (nodeID << 12) | seq, nil
	}
}
