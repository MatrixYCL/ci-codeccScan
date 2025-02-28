## 工具介绍

**CppLint**是谷歌开源的C++代码风格检查工具，可确保C++代码符合谷歌编码规范，并能检查语法错误。它是由Python语言编写的。

目前支持的规则已经在checkers.json中进行了声明。

## 开发新规则的步骤

1.从master拉规则分支，命名建议为：story_xxx_MMdd （xxx为规则名，MMdd表示月日）

2.将规则分支代码clone到本地

3.CppLint工具分为原生规则和自定义规则。
- 所有自定义规则都存放在third_rules目录，新建规则请以规则名命名，规则开发可参考tosa_fn_name_length.py；
- 原生规则都在/tool/cpplint.py中，可通过检索规则名例如build/namespaces相关逻辑的修改。

4.新建规则后，更新sdk/config/third_config.txt配置文件，

5.在checker.json中添加新增规则的描述。在描述时需要说明该规则对应到哪一条规范，并附上链接。规则描述示例如下：

- 单测函数行数限制也是普通函数的2倍，即为160行。即单测函数的行数，包含代码行、注释行、空白行，不得超过160行，否则需要重新评估函数的功能是否过于复杂需进行分拆。  [tencent standards/cpp 2.6](https://{github.com/xxxxx}/standards/cpp#26-%E5%BF%85%E9%A1%BB%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95)

6.在test目录下添加规则测试代码文件

7.MR代码到test分支

8.执行流水线部署http://{devops.public.url}/console/pipeline/codecc-tool-auto/p-65ca5c51adf641e0be1e68d042d6f422

9.测试完成提合并请求到master分支，由工具负责人jimxzcai审核评估后正式发布到生产
