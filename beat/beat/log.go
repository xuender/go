package beat
import (
  "time"
)

// 日志
type Log struct {
  // 消息
  Msg   string
  // 创建时间
  Ca    time.Time
}
