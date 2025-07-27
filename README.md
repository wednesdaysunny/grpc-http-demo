# gRPC-Gateway Demo

基于 gRPC-Gateway 的 Go 项目，同时支持 gRPC 和 HTTP 接口。

## 项目结构

```
├── cmd/
│   └── main.go          # 主程序入口
├── server/
│   ├── user_service.go  # gRPC 服务实现
│   └── gateway.go       # HTTP 网关服务
├── proto/
│   └── user/
│       ├── user.proto   # Protocol Buffer 定义
│       ├── user.pb.go   # 生成的 Go 代码
│       ├── user_grpc.pb.go
│       ├── user.pb.gw.go
│       └── user.swagger.json
├── pkg/
│   └── config/          # 配置相关
├── buf.yaml             # Buf 配置文件
├── buf.gen.yaml         # 代码生成配置
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 功能特性

- ✅ 同时支持 gRPC 和 HTTP 接口
- ✅ 基于 Protocol Buffers 的 API 定义
- ✅ 自动生成 HTTP RESTful API
- ✅ 自动生成 Swagger 文档
- ✅ 健康检查端点
- ✅ 优雅关闭
- ✅ 反射支持（便于调试）

## API 接口

### gRPC 端口: 9091
### HTTP 端口: 9090

### 用户管理 API

| 方法 | gRPC | HTTP | 描述 |
|------|------|------|------|
| CreateUser | CreateUser | POST /api/v1/users | 创建用户 |
| GetUser | GetUser | GET /api/v1/users/{id} | 获取用户 |
| ListUsers | ListUsers | GET /api/v1/users | 用户列表 |
| UpdateUser | UpdateUser | PUT /api/v1/users/{id} | 更新用户 |
| DeleteUser | DeleteUser | DELETE /api/v1/users/{id} | 删除用户 |

### 系统接口

- `GET /health` - 健康检查

## 快速开始

### 1. 安装依赖

```bash
make deps
```

### 2. 生成 Proto 文件（可选）

```bash
make proto
```

### 3. 运行服务

```bash
# 开发模式
make dev

# 或者编译后运行
make run
```

### 4. 测试接口

#### HTTP 接口测试

```bash
# 创建用户
curl -X POST http://localhost:9090/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "张三", "email": "zhangsan@example.com"}'

# 获取用户列表
curl http://localhost:9090/api/v1/users

# 获取特定用户
curl http://localhost:9090/api/v1/users/{user_id}

# 更新用户
curl -X PUT http://localhost:9090/api/v1/users/{user_id} \
  -H "Content-Type: application/json" \
  -d '{"name": "李四", "email": "lisi@example.com"}'

# 删除用户
curl -X DELETE http://localhost:9090/api/v1/users/{user_id}
```

#### gRPC 接口测试

使用 grpcurl 工具：

```bash
# 列出服务
grpcurl -plaintext localhost:9091 list

# 创建用户
grpcurl -plaintext -d '{"name": "张三", "email": "zhangsan@example.com"}' \
  localhost:9091 user.UserService/CreateUser

# 获取用户列表
grpcurl -plaintext -d '{}' localhost:9091 user.UserService/ListUsers
```

## 开发工具

### Make 命令

```bash
make build        # 编译项目
make run          # 编译并运行
make dev          # 开发模式运行
make clean        # 清理构建文件
make proto        # 生成 proto 文件
make test         # 运行测试
make fmt          # 格式化代码
make lint         # 代码检查
make deps         # 安装依赖
make install-tools # 安装开发工具
```

### 所需工具

- Go 1.19+
- buf (Protocol Buffer 工具)
- grpcurl (gRPC 测试工具，可选)
- golangci-lint (代码检查工具，可选)

## 技术栈

- **gRPC**: 高性能 RPC 框架
- **gRPC-Gateway**: gRPC 到 HTTP 的反向代理
- **Protocol Buffers**: 接口定义语言
- **Buf**: 现代 Protocol Buffer 工具链

## 扩展功能

### 添加新的服务

1. 在 `proto/` 目录下创建新的 `.proto` 文件
2. 运行 `make proto` 生成代码
3. 在 `server/` 目录下实现服务逻辑
4. 在 `cmd/main.go` 中注册新服务

### 添加中间件

可以在 `server/gateway.go` 中添加 HTTP 中间件，或在 gRPC 服务中添加拦截器。

### 配置管理

在 `pkg/config/` 目录下添加配置管理功能。

## 注意事项

- 确保 gRPC 服务先启动，HTTP 网关才能正常连接
- 生产环境建议使用 TLS 加密
- 可以根据需要添加认证和授权中间件
- 建议添加日志和监控功能