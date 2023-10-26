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