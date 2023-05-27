症状：
1、容器端口不能正常访问

排查步骤：
1、检查容器是否正常启动，进入容器是否能正常telnet通端口
2、检查端口是否绑定0.0.0.0
3、查看input规则：iptables  -L  -n；增加input规则：iptables -A INPUT -p tcp --dport 3306 -j ACCEPT
4、检查安全组是否放行端口