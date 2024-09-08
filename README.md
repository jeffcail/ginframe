# ginframe
> 基于Go语言gin框架搭建的可快速开发的微服务脚手架


<a href="https://github.com/jeffcail/ginframe/releases">
    <img src="https://img.shields.io/github/release/ginframe/releases.svg" alt="GitHub release">
  </a>
   <a href="https://github.com/jeffcail/ginframe/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="license">
  </a>

[TOC]- [关于我](#关于我)
- [关于我](#关于我)
- [ginframe](#ginframe)
    - [简介](#简介)
    - [目录结构](#目录结构)
    - [安装](#安装)
    - [目的及优势](#目的及优势)
    - [职责](#职责)
      - [1. 文件配置](#文件配置)
      - [2. api路由](#api路由)
      - [3. rpc](#rpc)
      - [4. websocket路由](#websocket路由)
      - [5. api返回统一格式](#api返回统一格式)
      - [6. gorm](#gorm)
      - [7. redis](#redis)
      - [8. Mongo](#Mongo)
      - [9. uber.zap.log](#uber.zap.log)
      - [10. GOMAXPROCS](#GOMAXPROCS)
      - [11. httprequest](#httprequest)
      - [12. leveldb](#leveldb)
      - [13. ElasticSearch](#ElasticSearch)
      - [14. AES](#AES)
      - [15. 加密](#加密)
      - [16. 动态搜索+分页](#动态搜索+分页)
      - [17. map合并和并发安全map](#map合并和并发安全map)
      - [18. 时间处理工具类](#时间处理工具类)
      - [19. 敏感词识别](#敏感词识别)
      - [20. 邮件类工具](#邮件类工具)
      - [21. kafka生产者、消费者](#kafka生产者、消费者)
      - [22. etcd客户端连接初始化 (写入、读取、修改、删除)](#etcd客户端连接初始化 (写入、读取、修改、删除))
      - [23. jwt](#jwt)
      - [24. 登录签发token](#登录签发token)
      - [25. token校验中间件](#token校验中间件)
      - [26. 账号登录状态是否被禁用校验中间件](#账号登录状态是否被禁用校验中间件)
      - [27. 常用正则表达式](#常用正则表达式)
      - [28. 数组切片去重](#数组切片去重)
# 关于我
21实际拾荒人

# ginframe
基于Go语言gin框架搭建的可快速开发的微服务脚手架


## 简介
基于gin框架，搭建一个快速开发的脚手架。

## 目录结构
```markdown
server-common                         -- 服务公共模块
    config                            -- 解析配置方法
    const                             -- 全局常量
    driver                            -- 全局驱动
    nacosRF                           -- 全局nacos配置
    pkg                               -- 全局公共包
        gorm
        httprequest
        jwt
        leveldb
        mongo
        redis
        uber
        viper
        wetcd
        wkafka
    process                           -- cpu核心
    servers                           -- gprc 服务发现
    utils                             -- 全局工具类
        email                         -- 发送邮件
        encry                         -- 加密
        enum                          -- api统一分装返回
        ip                            -- ip工具
        maps                          -- map工具
        orm                           -- gorm动态搜索、分页
        regmatch                      -- 正则匹配
        slice                         -- 切片操作
        wordsfilter                   -- 敏感词过滤
        wtime                         -- 时间处理
server-user                           -- 用户服务
    boot                              -- 启动目录
        db.go
        grpc.go
        http.go
        init.go
        log.go
        ws.go
    cachedb                           -- 缓存操作
    core                              -- 核心目录
        db.go
    daos                              -- daos
    global                            -- user服务全局配置
    grpcservices                      -- grpc services
    handler                           -- 控制器
    input                             -- 入参
    middlewares                       -- 中间件
    models                            -- 模型
    out                               -- 出参映射
    pb                                -- protobuf生成的文件目录
    proto                             -- protobuf文件目录
    router                            -- 路由
        api.go
        ws.go
    rpc                               -- rpc
    scripts                           -- 脚本
    service                           -- 服务层
    ulogger                           -- 服务日志
    ws                                -- websocket
    main.go                           -- 入口文件
server-test                           -- 测试服务
......                                -- 其他服务
.gitignore
go.mod                                -- mod包管理文件
LICENSE
README.md
```

## 安装
安装完之后名字ginframe可改，可根据自己的需求精简或者添加架子结构。也可直接使用
```shell
git clone https://github.com/jeffcail/ginframe.git

cd ginframe

make env

make mod

```

## 目的及优势

* 快速上手、快速开发、快速交付
* 高性能、高扩展，避免重复造轮子


## 职责
快速开发、避免重复造轮子

### 文件配置
application.yml为主配置文件.ConfigRemote觉得是否启用远程配置，支持Nacos.
config.yml 为应用配置

### api路由
http路由 

### rpc

### websocket路由
websocket路由

### api返回统一格式
成功、失败、分页

### gorm

### redis

### Mongo

### uber.zap.log

### GOMAXPROCS

### httprequest
http请求. GET、POST带header头和参数

### leveldb

### ElasticSearch

### AES
可用于api接口参数加密

### 加密
md5加密 、sha256加密 、sha512加密 、文件md5加密 、 （密码+盐）hash加密(可以用于加密登录密码).

### 动态搜索+分页

### map合并和并发安全map

### 时间处理工具类

### 邮件类工具

### kafka生产者、消费者

### etcd客户端连接初始化 (写入、读取、修改、删除)

### jwt

### 登录签发token

### token校验中间件

### 账号登录状态是否被禁用校验中间件

### 常用正则表达式
1. 手机号
2. 座机号
3. 18位身份证号
4. 护照编号
5. 港澳通行证
6. IP地址(ipv4)
7. IPV6
8. MAC地址
9. 电子邮箱
10. 统一社会信用代码
11. 密码
12. 网址URL 带端口号
13. 网址URL 不带端口号
14. 金额

### 数组切片去重
