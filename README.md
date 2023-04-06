## 一、项目搭建记录
```bash
# 项目结构搭建
mkdir ldap-debug
cd ldap-debug/
go mod init ldap-debug
# Go包 ldap 安装
go get github.com/go-ldap/ldap/v3
touch main.go
```

## 二、自行尝试
>  国内加速克隆:
> `git clone https://ghproxy.com/https://github.com/mindoc-org/ldap-debug.git`

### (一)、设置Go包源加速和安装Go包
```bash
# 启用 Go Mod
set GO111MODULE=on
# 使用国内源
set GOPROXY=https://proxy.golang.com.cn,direct
# 安装Go包
go mod tidy -v
```
### (二)、Windows 编译&运行
```bash
# 编译
go build -o ldap-debug.exe main.go
# 运行
.\ldap-debug.exe
```

### (三)、Linux 编译&运行
```bash
# 编译
go build -o ldap-debug main.go
# 运行
./ldap-debug
```
