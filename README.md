# PCC性能挑战赛2017

比赛提供3台机器(8核16GB),完成类似facebook的like功能

## Api:
* /pcc?action=like&oid=xxx&uid=xxx
返回:
```
{
"oid":1, "uid":1, "like_list":[1]
}
```

* /pcc?action=is_like&oid=xxx&uid=xxx
返回:
```
{
"oid":1, "uid":1, "is_like":1
}
```
* /pcc?action=count&oid=xxx
返回:
```
{
"oid":1, "count":1024
}
```
* /pcc?action=list&oid=xxx&uid=xxx&page_size=xxx&is_friend=x
返回:
```
{
"oid":1, "like_list":[1, 2]
}
```
## 数据规模:
- oid: 10M
- uid: 10M
- 好友数: 1000正太分布,假设平均200
- 假设平均每用户like 100文章
- 假设每篇文章缓存最近20个like用户
- 假设存每个id使用10字节
## 存储设计:
- 以uid为key,用set存like的全部文章
  - 空间: 10M * 100 * 10B = 10GB
- 以uid为key,用list存用户的好友列表
  - 空间: 10M * 200 * 10B = 20GB
- 以oid为key,存list最近like用户
  - 空间: 10M * 20 * 10 B= 2GB
- 以oid为key,用list存全部like用户数
  - 空间:
- 以oid为key,存like计数
  - 空间: 10M * 10B = 100MB

将用户like文章,文章最近like用户,文章like计数存在内存中,其它信息存在磁盘中
like, is_like和count可以实现较高的qps.

list api:

跟好友无关的查询可以一次查磁盘得出
跟好友相关的查询可以一次查磁盘得到好友列表,再一次查内存得到like列表

## 技术选型
- 语言: golang
- framework: gin
- 存储: redis 2.x, ssdb, twemproxy



