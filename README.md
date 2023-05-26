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
