# scriptBrick
build function script to a final operation or value

一个轻量级脚本解析引擎，利用了ast语法树，同时结合函数工厂模式，支持嵌入到任何golang的项目
比如

    1、计算1000+sum(100.23,100.11,sum(1,2))+(100*2)的结果
    2、输出stringJoin(abc,sum(100.789,30.56),def100)的结果
    3、执行deleteTask(getTaskId(taskName),alltimeline)删除项目运行时的某个任何item

理论上go能做到的都能脚本化

自动识别函数名，支持带入一个参数为自定义context，实现脚本化配置项目流程

已经上了生产级，目前支持高并发项目1000req/S，没有问题

性能已经尽量优化，带测试用例

有优化建议的可以留言给我

