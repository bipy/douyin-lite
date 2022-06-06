# 抖音极简版

**TODO**

- 库表
  - [x] 用户表
  - [x] 视频表
  - [x] 关注表
  - [x] 点赞表
  - [x] 评论表
- 接口
  - [x] 视频流
  - [x] 用户注册
  - [x] 用户登录
  - [x] 用户信息
  - [ ] 投稿
  - [x] 发布列表
  - [x] 赞操作
  - [x] 点赞列表
  - [x] 评论操作
  - [x] 评论列表
  - [x] 关注操作
  - [x] 关注列表
  - [x] 粉丝列表
- 组件
  - [x] 持久化数据库
  - [ ] 缓存
  - [ ] 定时任务
  - [ ] 队列
  - [ ] 对象储存
- 安全
  - [ ] 日志
  - [ ] 单元测试
  - [x] JWT

**Build and Serve**

1. Set up MySQL Server

2. Run `resources/sql/table.sql` before deploy

```bash
git clone git@github.com:bipy/douyin-lite.git
cd douyin-lite

# Edit release.env (important)

docker build -t douyin-lite:v1.0 .
docker run -d --name douyin douyin-lite:v1.0
```
