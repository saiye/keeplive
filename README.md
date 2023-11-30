# keeplive
项目用于报告系统CPU ,内存 使用情况，通过钉钉推送给开发者 
[钉钉机器人文档](https://open.dingtalk.com/document/robots/custom-robot-access):

#### linux install
```text
build:
GOOS=linux GOARCH=amd64 go build  -ldflags="-s -w" -o keeplive

mkdir /usr/local/keeplive
mkdir /usr/local/keeplive/conf

mv ./keeplive  /usr/local/keeplive
cp ./conf/keeplive.ini.demo  /usr/local/keeplive/conf/keeplive.ini
cp ./conf/keeplive.service /etc/systemd/system/

chown -R www:www /usr/local/keeplive

systemctl daemon-reload
systemctl enable keeplive
systemctl start keeplive
systemctl status keeplive
```

#配置文件详解
```text
[app]
#环境标记,一般是机器名称
env="local"
#cpu,Memory警告百分比0-95
percent=10
#定时器时间单位秒
second=10

[dingtalk]
#钉钉token
access_token=
#密钥
secret=
#空的话@所有人,非空@手机号，多个则英文,号分割
phone_list=
#读取警报关键词，自己在钉钉机器人控制面板定义
keyword=danger

[check]
#url 数量为n则配置url1-url{n},如果链接返回错误或httpStatusCode!=200 report message 到钉钉
count=2
name1=phpServer
url1=http://p.dev.com/check

name2=gameServer
url2=http://abc.dev.com/check
```