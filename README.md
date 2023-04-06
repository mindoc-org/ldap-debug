```bash
# 项目结构搭建
mkdir ldap-debug
cd ldap-debug/
go mod init ldap-debug
touch main.go

# 环境变量设置
set GO111MODULE=on
set GOPROXY=https://proxy.golang.com.cn,direct
# go包安装
go get github.com/go-ldap/ldap/v3
# go mod tidy -v

# 编译
go build -o ldap-debug.exe main.go
# 运行
.\ldap-debug.exe
```