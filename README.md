# [GoCodeGenerator](https://github.com/polariseye/GoCodeGenerator)
go的代码生成

## 有待改进的东西一览
1. 不支持命令缩写形式
2. 现在规划的命令太多，需要简化一下

## 使用流程
1. 加载目标语言配置 loadconfig
2. 打开数据库连接 open 
3. 选择要使用的数据库 selectdb
4. 选择代码生成要使用的模板 selecttemplate 
5. 进行代码生成 build

## 代码使用注意事项:
1. flagly项目请使用：[flagly](https://github.com/polariseye/flagly)