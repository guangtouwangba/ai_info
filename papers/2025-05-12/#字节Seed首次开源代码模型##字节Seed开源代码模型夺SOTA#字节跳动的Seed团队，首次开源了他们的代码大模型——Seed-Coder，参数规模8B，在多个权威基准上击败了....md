# #字节Seed首次开源代码模型##字节Seed开源代码模型夺SOTA#字节跳动的Seed团队，首次开源了他们的代码大模型——Seed-Coder，参数规模8B，在多个权威基准上击败了...

**URL**: https://weibo.com/6105753431/Prnj1uOSD

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%AD%97%E8%8A%82Seed%E9%A6%96%E6%AC%A1%E5%BC%80%E6%BA%90%E4%BB%A3%E7%A0%81%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%AD%97%E8%8A%82Seed%E9%A6%96%E6%AC%A1%E5%BC%80%E6%BA%90%E4%BB%A3%E7%A0%81%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#字节Seed首次开源代码模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%AD%97%E8%8A%82Seed%E5%BC%80%E6%BA%90%E4%BB%A3%E7%A0%81%E6%A8%A1%E5%9E%8B%E5%A4%BASOTA%23&amp;extparam=%23%E5%AD%97%E8%8A%82Seed%E5%BC%80%E6%BA%90%E4%BB%A3%E7%A0%81%E6%A8%A1%E5%9E%8B%E5%A4%BASOTA%23" data-hide=""><span class="surl-text">#字节Seed开源代码模型夺SOTA#</span></a><br><br>字节跳动的Seed团队，首次开源了他们的代码大模型——Seed-Coder，参数规模8B，在多个权威基准上击败了Qwen3、DeepSeek-R1，夺得SOTA成绩。<br><br>Seed-Coder共发布了三个版本：<br><br>- Base：基础模型<br>- Instruct：强化指令理解，通过监督微调+偏好优化，提升模型“听懂人话”的能力<br>- Reasoning：面向复杂推理任务，采用强化学习的方式，锻炼多步推理能力，在IOI 2024超过QwQ-32B<br><br>这个模型的特别之处在于，团队提出了“用小模型自管数据”的新范式，也就是模型自己策划训练数据，连生成和筛选都由模型完成，人工干预极少，具体方式包括：<br><br>1. 四类高质量数据源：<br>    <br>    - 文件级代码：从GitHub提取单个文件内容<br>        <br>    - 仓库级代码：保留项目结构关系<br>        <br>    - Git Commit：覆盖7400万次提交，格式化为代码变更预测任务<br>        <br>    - 网络代码相关内容：从网页提取结构化和非结构化代码信息<br>        <br>2. 数据去重与筛选：<br>    <br>    - SHA256+MinHash双重去重<br>        <br>    - 使用语法解析器排除错误代码<br>        <br>    - 通过LLM评分模型评估代码可读性、模块性等质量维度<br>        <br>3. 定制化评分机制：<br>    <br>    - 针对不同网站内容风格（博客/论坛）制定不同评分标准，避免误判有价值内容<br>        <br>目前完整代码和模型均已开源：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fbytedance-seed-coder.github.io%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1bx8lx14tj30w40tgh4x.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1bx8fbqduj30zk0fy0y2.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1bx8l0wvbj30zk0iz0ym.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1bx8mgdclj30zk0cf0zs.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

字节跳动Seed团队首次开源代码大模型Seed-Coder（8B参数），包含Base、Instruct和Reasoning三个版本，在多项基准测试中超越Qwen3等模型达到SOTA。其创新点在于采用"小模型自管数据"范式，通过四类高质量数据源（文件/仓库级代码、Git提交、网络代码内容）和自动化筛选机制（双重去重、语法检查、LLM质量评估），大幅减少人工干预。Reasoning版本通过强化学习实现复杂推理能力，在IOI 2024超越QwQ-32B。模型及代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T01:30:46Z
- **目录日期**: 2025-05-12
