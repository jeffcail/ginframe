[TOC]- [关于我](#关于我)
- [关于我](#关于我)
- [echoframe](#ginframe)
    - [简介](#简介)
    - [安装](#安装)
    - [目的及优势](#目的及优势)
    - [技术点对应文档](#技术点对应文档)
    - [职责](#职责)
      - [1. 文件配置](#文件配置)
      - [2. api路由](#api路由)
      - [3. websocket路由](#websocket路由)
      - [4. api返回统一格式](#api返回统一格式)
      - [5. gorm](#gorm)
      - [6. redis](#redis)
      - [7. Mongo](#Mongo)
      - [8. uber.zap.log](#uber.zap.log)
      - [9. GOMAXPROCS](#GOMAXPROCS)
      - [10. httprequest](#httprequest)
      - [11. leveldb](#leveldb)
      - [12. ElasticSearch](#ElasticSearch)
      - [13. AES](#AES)
      - [14. 加密](#加密)
      - [15. 动态搜索+分页](#动态搜索+分页)
      - [16. map合并和并发安全map](#map合并和并发安全map)
      - [17. 时间处理工具类](#时间处理工具类)
      - [18. 敏感词识别](#敏感词识别)
      - [20. 邮件类工具](#邮件类工具)
      - [21. kafka生产者、消费者](#kafka生产者、消费者)
      - [22. etcd客户端连接初始化 (写入、读取、修改、删除)](#etcd客户端连接初始化 (写入、读取、修改、删除))
      - [22. jwt](#jwt (jwt))
# 关于我
执着于理想，纯粹与当下...

# ginframe
基于Go语言gin框架搭建的可快速开发的脚手架


## 简介
基于echo框架，搭建一个快速开发的脚手架。

## 安装
安装完之后名字ginframe可改，可根据自己的需求精简或者添加架子结构。也可直接使用
```shell
git clone https://github.com/jeffcail/ginframe.git

cd ginframe

go mod tidy

go run main.go
```

## 目的及优势

* 快速上手、快速开发、快速交付
* 高性能、高扩展，避免重复造轮子

## 技术栈和对应的包
* viper: https://github.com/spf13/viper
* Nacos: https://nacos.io/zh-cn/
* Echo: https://github.com/gin-gonic/gin

## 职责
快速开发、避免重复造轮子

### 文件配置
application.yml为主配置文件.ConfigRemote觉得是否启用远程配置，支持Nacos.
config.yml 为应用配置

### api路由
http路由 

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