基于DAG(有向无环图)任务依赖&编排

###DAG
DAG（Directed Acyclic Graph）是的核心概念之一，DAG体现的是你的工作流，它由tasks.json定义，其中包含了你想要运行的一系列task，同时其中还定义了这些task的依赖关系
DAG由其中的task组成.
DAG本身并不关注A，B，C三个任务的具体内容，它关注的是三个任务的执行顺序以及依赖条件

###Task
Task是DAG中定义工作的基本单位，它的地位等同于工作流中的一个节点。Task和Operator是一枚硬币的正反两面，Task代表工作的抽象概念，Operator定义了对应Task要做的具体任务。 同一个DAG中的Task之间一般有先后顺序和依赖关系
Task的生命周期:
Running (task开始执行)
Success (task成功结束)
Task可以引用其他Task的输入输出，使用${taskxxx.output}的方式引用。引用语法为jsonPath，除最基础的${taskxxx.output}的值解析方式外，还支持其他复杂操作，如过滤等，具体见json-path语法

###Operators
DAG定义了一个工作流如何执行，而Operator定义了一个task执行的具体任务，是编写具体任务的类。Operator包括很多种类，BashOperator用来执行Bash命令，GoOperator可以执行Go函数，MySOperator可以操作MySQL数据库执行相关操作，当然你也可以从BaseOperator中继承并开发自己的Operator

###Executor
Executor是task的执行器，它有多种配置方式，如SeqExecutor串行运行task（适用于开发环境，是默认的配置），LocalExecutor可以在本地并发运行task 

##其他
失败重试策略retry backoff

###代码结构
泛化调用?
通过泛化调用客户端可以在服务方没有提供接口协议和不依赖服务方API的情况下对服务进行调用

平台化建设

参考资料：
airflow