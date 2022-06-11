# 红岩网校期末考核 -- 象棋Web应用



### 效果图

1.All func

![img]( https://s3.bmp.ovh/imgs/2022/06/11/dff88dd8074ad9ee.png )

2.移动棋子的返回数据

![img]( https://s3.bmp.ovh/imgs/2022/06/11/3bc3c6fe2c0d245c.png )

### API介绍

#### 一.User 操作

1.POST address:8080/user             注册

2.GET address:8080/user			  登录

#### 二.游戏界面 操作

1.POST address:8080/chess/create   创建房间

(可指定密码)

2.GET    address:8080/chess/join/:name 加入指定房间

3.POST  address:8080/chess/moving       操作棋子，数据通过WebSocket返回

4.POST  address:8080/chess/prepare	准备

