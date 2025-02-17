<a href="./README_EN.md">English</a>

# 免责声明🧐

本工具仅面向 **合法授权** 的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。**请勿对非授权目标进行扫描**。

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任。

在安装并使用本工具前，**请您务必审慎阅读**、充分理解各条款内容，限制、免责条款或者其他涉及您重大权益的条款可能会以加粗、加下划线等形式提示您重点注意。
除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要安装并使用本工具。您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。

# 工具来源及其说明

```
（1）nps通信流量比较稳定，但特征抓的比较死，所以基于原版的nps进行二次开发
```

### 说明

为了工具的免杀性及其后期修改，本人不公开源码。 **本人承诺，工具无毒，只能简单进行二开**

# 魔改部分

```
（1）重写了nps的认证过程，通信过程均进行加密
（2）重些了npc的部分，预计后续分离config文件进行加载
（3）进行了nps未授权漏洞的修复，避免了默认配置未授权
（4）支持config文件从远端进行加载，在传输过程中均实现流量加密
```

# 编译

```shell
cd cmd/nps
sh windows.sh

cd cmd/npc
sh windows.sh


需要修改账号密码请修改nps中的good.conf
```

# 免杀情况

（这是魔改后的demo上去的，还请各位测试切莫进行☁️测试、沙箱测试、联网测试）

### 魔改后流量

魔改后的工具流量就不进行抓取了，需要的话，大家可以自行进行测试。

### 某社区☁️沙箱（demo版测试）

<img width="1244" alt="image" src="./images/image-20230521162635386.png">

### virustotal

<img width="1244" alt="image" src="./images/image-20230521141510009.png">

### windows defender（静态）

<img width="1154" alt="image" src="./images/image-20230521141351320.png">

### windows defender（动态）

<img width="1492" alt="image" src="./images/image-20230521141429304.png">

### 360、火绒等其他杀软未进行测试

# 项目使用

未进行测试nps服务端的注册，所以目前主要还是 ./nps的方式来运行

## 服务端使用

**服务端使用注意事项**

```
（1）在使用服务端的时候需要down下来 cmd/nps下的conf和web两个目录
（2）把conf和web两个目录和nps目录同级
（3）修改服务端密码（后期也许会支持md5、或者登陆后修改密码）最新版默认使用good.conf
```

<img width="713" alt="image" src="https://github.com/Q16G/npsmodify/assets/113832601/63c27034-e456-4ec7-a0f9-b08d0e7c28ca">

<img width="1425" alt="image" src="https://github.com/Q16G/npsmodify/assets/113832601/58a901da-6b92-45e8-8055-2f18e727cae0">


<img width="1511" alt="image" src="https://github.com/Q16G/npsmodify/assets/113832601/95d06b90-fd42-4858-9b3b-528e28d3315d">



<img width="802" alt="image" src="./images/image-20230521141541183.png">

## 客户端使用

### 配置文件启动

配置文件如下：

```
[common]
server_addr=127.0.0.1:8024
conn_type=tcp
vkey=123456
auto_reconnection=true
max_conn=1000
flow_limit=1000
rate_limit=1000
web_username=admin
web_password=123
crypt=true
compress=true
#pprof_addr=0.0.0.0:9999
disconnect_timeout=60

```

<img width="768" alt="image" src="./images/image-20230521141600000.png">

### 命令行启动

<img width="787" alt="image" src="./images/image-20230521141615884.png">

### 命令行以远端配置文件启动

```
npc.exe -rconfig ServerConfig地址
eg.
npc.exe -rconfig 127.0.0.1:23123
```

# 项目进度

本人由于个人原因，此项目停止维护。

✅ 2023.5.19 重新写了通信认证协议

✅ 2023.5.20 把连接流量进行混淆，仅仅支持客户端命令行启动，未支持conf文件启动

✅ 2023.5.21 支持本地config文件加载

✅ 2023.5.23 支持config文件**仅**从服务端拉取

✅ 2023.5.30 重新修改流量，进行了深度混淆

# 后续增加

（1）~~实现其他协议流量的魔改~~

# 参考

```
https://github.com/ehang-io/nps
```

