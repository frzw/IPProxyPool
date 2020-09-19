# 代理池设计图

![](https://ftp.bmp.ovh/imgs/2020/09/53e9923f58db3611.jpg)

# go

- 优质的代理IP生成网站

- go爬取网站的代理IP
    - echo框架

- 代理IP存入mangodb

- 测试IP的有效性
    - 有效的标准
        - 访问百度等网站
    - 有效存入

- 存入PostgreSQL

- 定时测试PostgreSQL中IP是否有效、定时获取最新IP代理
    - 有效的标准
        - 访问百度等网站
    - 无效删除

# python

- 从POstgreSQL获取有效IP

- 从使用IP进行爬虫
    - 怎么使用代理IP进行爬虫
    - 怎么判断代理IP不可用需要使用后续的IP进行爬虫


# 技术栈
# htmlquery
 
# 优质代理网站
- https://ip.jiangxianli.com/?page=1

# GitHub参考网站
- https://github.com/jhao104/proxy_pool