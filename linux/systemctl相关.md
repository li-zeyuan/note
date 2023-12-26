### systemctl status查看service不断重启
原因：服务发生Panic
解决：
- journalctl -xe | grep ${service}
- journalctl -u ${service} -r
- 文档：https://wangchujiang.com/linux-command/c/journalctl.html