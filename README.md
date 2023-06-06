# istep

刷步数使用的程序（可选基于网站一起步数、网站刷步或官方API三种方式）

## 零、实现原理

基于zepp life接口修改，入口使用[17bushu](https://www.17bushu.com/)

核心前置步骤：`下载zepp life` -> `注册/登录` -> `绑定支付宝/微信`

## 一、下载

在[releases](https://github.com/ns-cn/istep/releases)中下载对应操作系统的最新版本

## 二、配置方法

按照默认的`istep.json`配置文件，修改`username`和`password`为自己的账号密码，及需要修改的步数的最大最小值

```json
{
  "showClient": false,
  "users": [
    {
      "user": "18716524101",
      "password": "cEBzc3cwcmQ=",
      "encoded": true,
      "ignore": false,
      "steps": {
        "min": 30000,
        "max": 30000
      }
    }
  ]
}
```

| 字段         | 说明                                    |
|------------|---------------------------------------|
| showClient | 是否显示客户端                               |
| users      | 用户列表                                  |
| user       | 用户名，可选zepplife的手机号或邮箱                 |
| password   | 密码 （可选配置base64编码）                     |
| encoded    | 密码是否编码(采用base64简单编码)，false则代表密码采用明文方式 |
| ignore     | 是否忽略该用户 （忽略则不执行）                      |
| steps      | 步数范围                                  |
| min        | 最小步数                                  |
| max        | 最大步数                                  |

## 三、运行

```shell
istep                # 使用默认的配置文件
istep -l other.json  # 读取其他配置文件
```

## 四、帮助文档

```shell
> istep -h
基于zepp life修改步数的程序

Usage:
  istep [flags]
  istep [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     打印当前版本号

Flags:
  -h, --help          help for istep
  -l, --load string   config file (default "istep.json")

Use "istep [command] --help" for more information about a command.
```