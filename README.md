# douyin team: beautifyDm

运行前安装一下依赖
```shell
go mod tidy
```
运行
```shell 
rm dev || go build -o dev && ./dev
```
![官方文档](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof)

## 接口分配

| principal | interface                     | state       |
|-----------|-------------------------------|-------------|
| null      | /douyin/feed                  | wait charge |
| yly       | /douyin/user/regiser          | done        |
| yly       | /douyin/user/login            | done        |
| null      | /douyin/user/                 | wait charge |
| null      | /douyin/publish/action/       | wait charge |
| null      | /douyin/publish/list/         | wait charge |
| null      | /douyin/favorite/action/      | wait charge |
| null      | /douyin/favorite/list/        | wait charge |
| null      | /douyin/comment/action/       | wait charge |
| null      | /douyin/comment/list/         | wait charge |
| null      | /douyin/relation/action/      | wait charge |
| null      | /douyin/relation/follow/list/ | wait charge |
| null      | /douyin/friend/list           | wait charge |
| null      | /douyin/message/chat          | wait charge |
| null      | /douyin/message/list          | wait charge |

### 详细说明请移动docs

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试
