# Agenda for Golang

------

我们此次作业使用Cobra实现了终端的Agenda系统，主要包括以下功能：

- **登录**
- **登出**
- **注册**
- **删除账户**
- **查询用户**
- **创建会议**
- **增加会议参与者**
- **会议查询**
- **删除会议参与者**
- **取消会议**
- **退出会议**
- **清空会议**

[可参考AgendaC++源码](https://github.com/MoRunChang2015/Agenda)
## 任务目标

 - 熟悉 go 命令行工具管理项目
 - 综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
 - 使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令不会影响其他命令的代码
 - 项目部署在 Github 上，合适多人协作，特别是代码归并
 - 支持日志（原则上不使用debug调试程序）


## 具体运行

从github上下载到本地，在linux下的agenda-go-cli-master下运行go build然后在终端进入当前文件夹，使用./agenda-go-cli加上你所需要的操作命令便可以运行。（当然你需要完成Go语言的相关环境配置）
        
#运行指令和运行截图

具体运行过程与实验要求截图可以参见我们的实验博客
[实验要求截图可以参见我们的实验博客]( http://blog.csdn.net/qq_33689717/article/details/78403706/)



### 相关子命令: 
- 登录
	- Command: `login -u=username -p=password`
	- Parse:
		- -u: 用户名
		- -p: 密码
- 登出
	- Command: `logout`
	- No Parse
	- Extra:
		- 若有用户登录则打印相应登出信息，若无用户登录则提示出错
- 注册
	- Command: `register -u=username -p=password -m=mail -c=cellphone`
	- Parse:
		- -u:用户名
		- -p:密码
		- -m:邮箱地址
		- -c: 手机号码
	- Extra
		- 用户名具有唯一性，不允许与现有用户名相同

- 删除账户
	- Command: `deleteaccout`
	- No Parse
	- Extra: 
		- 如无用户登录则返回出错信息
		- 如有用户登录则删除用户作为发起者的会议，同时将用户从作为参与者的会议中删除
		- 删除登录信息
- 查询用户
	- Command: `queryuser`
	- Option: `-u=username`
	- Parse:
		- -u: 用户名

- 创建会议
	- Command: `createmeeting -t=title -p=participator -s=starttime -e=endtime`
	- Parse:
		- -t : 会议标题
		- -p: 参与者(以数组形式呈现)
		- -s: 开始时间
		- -e: 结束时间
	- Extra:
		- 只允许已登录用户进行此操作
		- 会议标题具有唯一性
		- 参与者必须为注册用户
		- 参与者数量 >= 1
		- 开始时间与结束时间应符合逻辑，采用24小时制
		- 不允许发起者或参与者在此时间段内有其他会议
- 增加会议参与者
	- Command: `addparticipator -t=title -p=participator`
	- Parse:
		- -t: 会议标题
		- -p: 参与者列表
	- Extra:
		- 仅允许已登录用户对自己发起的仍存在的会议进行增加操作, 否则返回出错信息
		- 参与者必须为注册用户
		- 参与者同样需要检查会议时间段内是否空闲
- 删除会议参与者
	- Command: `removeparticipator -t=title -p=participator`
	- Parse:
		- -t: 会议标题
		- -p: 参与者列表
	- Extra:
		- 仅允许已登录用户对自己发起的仍存在的会议进行删除操作, 否则返回出错信息
		- 参与者必须为注册用户并且存在于会议中
		- 参与者数量 == 0 时， 会议删除
- 会议查询
	- Command: `querymeeting -s=starttime -e=endtime`
	- Parse:
		- -s : 开始时间
		- -e : 结束时间
	- Extra:
		- 已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
		- 在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
		- 注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。
- 取消会议
	- Command: `deletemeeting -t=title`
	- Parse:
		- -t : 会议标题
	- Extra:
		- 仅允许已登录用户删除自己发起的会议
- 退出会议
	- Command: `quitmeeting -t=title`
	- Parse:
		- -t: 会议标题
	- Extra:
		- 仅允许已登录用户退出自己参加的会议
		- 若因此导致会议参与人数为0则删除会议
- 清空会议
	- Command: `clearmeeting`
	- No Parse
	- Extra:
		- 删除已登录用户发起的所有会议
