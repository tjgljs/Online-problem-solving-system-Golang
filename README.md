# Online-problem-solving-system-Golang

基于Gin、Gorm、Vue 实现的在线练习系统

后台语言：Golang、框架：Gin、GORM

前台框架：Vue、ElementUI


整合 Swagger
参考文档： https://github.com/swaggo/gin-swagger 接口访问地址：http://localhost:8080/swagger/index.html

// GetProblemList

// @Tags 公共方法

// @Summary 问题列表

// @Param page query int false "page"

// @Param size query int false "size"

// @Success 200 {string} json "{"code":"200","msg","","data":""}"

// @Router /problem-list [get]

安装 jwt

go get github.com/dgrijalva/jwt-go




系统模块：
 用户模块：
 密码登录
 邮箱注册
 用户详情
 题目管理模块
 题目列表、题目详情
 题目创建、题目修改
 
 分类管理模块：
 分类列表
 分类创建、分类修改、分类删除
 判题模块
 提交记录列表
 代码的提交及判断
 
 排名模块：
 排名的列表情况
 竞赛模块
 竞赛列表
 竞赛管理
 竞赛报名
