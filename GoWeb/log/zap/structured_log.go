package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	url := "https://github.com/qixia1998"
	// 结构化日志打印
	logger.Sugar().Infow("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)

	// 非结构化日志
	logger.Sugar().Infof("failed to fetch URL: %s", url)
}
