# heapdump-watcher

Go 语言开发的堆转储文件自动化监控管理工具

## 功能特性

- 🔍 自动监控目录中新生成的堆转储文件（*.prof）
- 📦 自动压缩堆转储文件为 ZIP 格式
- ☁️ 自动上传压缩后的文件至对象存储服务（OSS）
- 🔔 支持多种通知渠道：
  - 钉钉机器人通知
  - 企业微信机器人通知
  - 邮件通知

## 工作原理

1. 持续监控指定目录下的堆转储文件
2. 当检测到新的 `.prof` 文件时：
   - 等待文件写入完成
   - 将文件压缩为 ZIP 格式
   - 上传压缩文件至 OSS
   - 通过配置的渠道发送通知
   - 清理临时文件

## 配置说明

本应用支持以下配置项：
- 堆转储文件监控路径
- OSS 认证和配置信息
- 通知渠道配置：
  - 钉钉机器人 Webhook 设置
  - 企业微信机器人配置
  - 邮件服务器设置

## 环境要求

- Go 1.x 或更高版本
- OSS 服务访问权限
- 至少配置一个通知渠道（钉钉、企业微信或邮件）

## 安装部署

```bash
git clone https://github.com/yourusername/heapdump-watcher
cd heapdump-watcher
go mod download
```

## 使用方法

1. 配置相关设置
2. 运行应用：
```bash
go run main.go
```

## 参考文档

### 钉钉机器人
- [自定义机器人接入文档](https://open.dingtalk.com/document/orgapp/custom-robot-access)

### 企业微信机器人
- [群机器人配置说明](https://developer.work.weixin.qq.com/document/path/91770)

### 企业微信应用
- [集成示例](https://www.nowcoder.com/discuss/534745989103575040)

### 邮件集成
- [Go 邮件实现参考](https://learnku.com/go/t/70932)

## 贡献指南

欢迎提交 Pull Request 来改进这个项目！

## 开源协议

[请添加您的开源协议]