# douyin team: beautifyDm

运行前安装一下依赖
```shell
go mod tidy
```
运行
```shell 
rm dev >> /dev/null
go build -o dev && ./dev
```
### [官方文档](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof)
### 课程资料大全

- [gorm 详解](https://bytedance.feishu.cn/file/boxcngmUNHi2joONiiEOgSpJt8d)
- [go 三件套详解](https://bytedance.feishu.cn/file/boxcnKHOoYmud2SuUGmhFaGbjVb)
- [高性能 go](https://bytedance.feishu.cn/file/boxcngF8NWGNFuXUkdyQViZq6vd)
- [调优 go](https://bytedance.feishu.cn/file/boxcn7AkvSWnRkHEttsuYHqW24g)
- [实战项目-go 笔记服务](https://bytedance.feishu.cn/docx/Wwa4dfwScogfjLxclXKcStGEncd)
- [gorm 官网](https://gorm.io/)



### 参考
- [官方demo](github.com/RaymondCode/simple-demo/repository) 参考了很多很多
- [Eliaukle-go-blog](https://github.com/Eliaukle/Simple-Blog-Community.git) 参考登陆认证
- [详细教程](https://blog.csdn.net/qq_50737715/article/details/127437065)

## 接口分配

| principal   | interface                     | state       |
|-------------|-------------------------------|-------------|
| war3dgjddhh | /douyin/feed                  | doing       |
| war3dgjddhh | /douyin/user/regiser          | done        |
| war3dgjddhh | /douyin/user/login            | done        |
| war3dgjddhh | /douyin/user/                 | done        |
| war3dgjddhh | /douyin/publish/action/       | doing       |
| war3dgjddhh | /douyin/publish/list/         | doing       |
| null        | /douyin/favorite/action/      | wait charge |
| null        | /douyin/favorite/list/        | wait charge |
| null        | /douyin/comment/action/       | wait charge |
| null        | /douyin/comment/list/         | wait charge |
| null        | /douyin/relation/action/      | wait charge |
| null        | /douyin/relation/follow/list/ | wait charge |
| null        | /douyin/friend/list           | wait charge |
| null        | /douyin/message/chat          | wait charge |
| null        | /douyin/message/list          | wait charge |

### 详细说明请移动docs

### 功能说明

接口功能不完善

* 视频上传后会保存到本地 video 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试


# 快速启动

## 安装mysql [ x ]

### 创建一个database douyin
```sql
create database if not exists douyin
```
### 修改 repository/db_init.go 里面的init方法，配置一下相关信息, e.g user, password, port 改成你本机的

## 安装抖声app,[下载地址](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 安装好之后快速点击我的两次就会出现配置页面填入本机局域网ip

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

