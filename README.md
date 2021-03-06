# 交付工具
## helmChart快速dryrun的标准输出
### 需求
+ 读取当前或指定目录的所有chart包
  - ✓ 目前暂时支持tgz包(目前工作当中的目标文件都是这个格式)
  - ✓ 忽略文件夹
  - ✓ 忽略非.tgz的文件
+ 依次解压到当前或指定目录
  - ✓ 默认会在当前目录创建文件夹，并将解压路径指向该目录
  - ✓ 默认创建的文件夹叫ouput
  - ✓ 检查目标目录是否存在，如果不存在则创建
  - ✓ 依次解压tgz文件列表到指定的目标路径下
+ 通过helm命令渲染chart
  - 先检查helm是否存在，不存在添加helm
    - 判断系统 mac or centos
    - 暂时不做，需要用户准备k8s接口以及helm，默认用户已有环境
  - dry run 渲染
    -  ✓ 获取解压出的所有目录，并识别有效的目录
    -  ✓ 获取对应的文件夹内的Chart.yaml中的name的值，并以垓值为服务名称
    -  ✓ 输出注入到一个对应的yaml文件
    -  ✓ 可以指定ns 如果不指定 默认是default
    - 支持修改所有image的名称,替换仓库以及路径名称，不指定则不替换
    - 识别配置中的多个不同类型文件
    - 检查所有deployment中的container配置，是否包含limt和request，填充缺失,默认2c2G；如果存在limit，则以limit为request填充；
    - 替换storageclass

