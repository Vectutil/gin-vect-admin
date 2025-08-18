package robot

import (
	"bytes"
	"context"
	"encoding/json"
	"gin-vect-admin/internal/config"
	"net/http"
)

const (
	QWRobotMsgTypeText     = "text"
	QWRobotMsgTypeMarkdown = "markdown"
)

func SendQWRobot(ctx context.Context, data, QWRobotMsgType string) {
	// 默认是正式环境
	url := config.Cfg.WXRobot.ErrorRobot
	reqBody := WechatWebhookRequest{
		MsgType: QWRobotMsgType,
	}

	switch QWRobotMsgType {
	case QWRobotMsgTypeText:
		reqBody.Text.Content = data
	case QWRobotMsgTypeMarkdown:
		reqBody.Markdown.Content = data
	}

	jsonData, _ := json.Marshal(reqBody)
	// 创建 HTTP POST 请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	client.Do(req)
}
