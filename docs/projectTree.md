# douyin-dm 项目结构

## 注意的小点

### 每个go package 里面的go func name, name 大写字母开头表示，是可以在被引入的时候访问到的，小写字母开头表示只能package里面访问

### 找到局域网ip addr
#### win, cmd 输入 ipconfig /all 一般能找到
#### linux, terminal 输入 ip addr 一般能找到


```shell
├── common                             # 公共工具目录
│   └── jwt.go
├── controller                         # 类似spring那一套 controller 层
│   ├── comment.go
│   ├── common.go
│   ├── demo_data.go
│   ├── favorite.go
│   ├── feed.go
│   ├── message.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── dev                               # 打包的可执行文件
├── dev.sh                            # dev 脚本
├── docs                              # docs 目录
├── go.mod                            # go 依赖
├── go.sum
├── main.go                           # 入口
├── middleware                        # 中间件, 类似 spring那一套过滤器
│   └── AuthMiddleware.go             # auth, 并存放user信息到gin上下文
├── public                            
│   ├── bear.mp4
│   └── data
├── README.md
├── repository                        # 类似spring model + service层
│   ├── db_init.go
│   └── user.go
├── router.go                         # go 路由, 配置访问请求对应的处理方法
├── service                           # 其他服务, 可能搞 一个 socket来搞message
│   └── message.go
├── test                              # 测试, 
│   ├── base_api_test.go
│   ├── common.go
│   ├── interact_api_test.go
│   ├── message_server_test.go
│   └── social_api_test.go
```
