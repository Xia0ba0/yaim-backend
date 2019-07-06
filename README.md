# yaim-backend

## INSTALL
```bash
# set GOPATH in .bashrc
export GOPATH=[PATH TO (yaim-backend)]
source ~/.bashrc

# install dependentcies
rm -rf ./src/github.com
rm -rf ./src/xorm.io

go get -u github.com/kataras/iris
go get github.com/go-xorm/xorm
go get -u github.com/go-sql-driver/mysql

# config database
db type mysql

db user root
db password 1005

create database test;
```
## STUDY
* [quick start](https://iris-go.com/start/#api-examples)
* [basic concept](https://studyiris.com/doc/)
* [detailed examples grouped by different needs](https://studyiris.com/example/index.html)

## 思路:
0. 继续把登录，注册，鉴权做了 然后启用Session
1. 把后端的各个模块封装成 Serivce 然后把Service 静态注入到MVC控制器
2. 在Handler函数中 调用Service 进行数据的 接收 处理 封装成API并返回
3. 现在需要的模块还有 数据库Service 消息推送Service 加解密Serivcve