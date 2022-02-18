# 交付工具
## helmChart快速dryrun的标准输出
### 需求
+ 读取当前或指定目录的所有chart包
  - <font color=#28FF28 >✓</font> 目前暂时支持tgz包(目前工作当中的目标文件都是这个格式)
  - ✓ 忽略文件夹
  - ✓ 忽略非.tgz的文件
+ 依次解压到当前或指定目录
  - 默认会在当前目录创建文件夹，并将解压路径指向该目录
+ 通过helm命令渲染chart
  - 先检查helm是否存在，不存在添加helm
  - dry run 渲染
  - 可以指定ns 如果不指定 默认是default
  - 输出都统一注入到一个yaml文件
  
