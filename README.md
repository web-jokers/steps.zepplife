# ibushu

刷步数使用的程序（基于一起步数网页）

## 一、下载



## 一、配置方法

按照默认的`ibushu.json`配置文件，修改`username`和`password`为自己的账号密码，及需要修改的步数的最大最小值

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

| 字段         | 说明                     |
|------------|------------------------|
| showClient | 是否显示客户端                |
| users      | 用户列表                   |
| user       | 用户名，可选zeeplife的手机号或邮箱  |
| password   | 密码 （可选配置base64编码）      |
| encoded    | 密码是否已经编码(采用base64简单编码) |
| ignore     | 是否忽略该用户 （忽略则不执行）       |
| steps      | 步数范围                   |
| min        | 最小步数                   |
| max        | 最大步数                   |

## 二、运行

```shell
ibushu                # 使用默认的配置文件
ibushu -l other.json  # 读取其他配置文件
```

## 三、帮助文档

```shell
基于zeep life修改步数的程序

Usage:
  ibushu [flags]
  ibushu [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     打印当前版本号

Flags:
  -h, --help          help for ibushu
  -l, --load string   config file (default "ibushu.json")

Use "ibushu [command] --help" for more information about a command.
```