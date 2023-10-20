### 运行
```go
go run .
//or
go run main.go
```
- 配置读取文件优先级: 项目目录文件>命令行文件 > bash环境变量文件 > gin默认文件
- 项目目录文件:main.go文件里的InitiallizeViper方法中配置,比如:InitiallizeViper("config.yaml") //使用本地配置文件
- 命令行文件:在终端执行go run main.go -c 文件名,比如:go run main.go -c config.yaml
- base环境变量文件:在自己电脑设置环境变量,比如:mac在终端执行:export EWA_CONFIG="config.yaml",成功后在自己项目运行:go run main.go
- gin默认文件:在main.go文件里的AppMode变量里赋值,比如:const AppMode = "debug" //运行环境，主要有三种：debug、test、release,对应core/internal/constants.go文件里的常量
