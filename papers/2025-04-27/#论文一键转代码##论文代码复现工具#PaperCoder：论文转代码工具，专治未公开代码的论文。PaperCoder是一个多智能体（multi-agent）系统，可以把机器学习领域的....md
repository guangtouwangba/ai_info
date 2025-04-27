# #论文一键转代码##论文代码复现工具#PaperCoder：论文转代码工具，专治未公开代码的论文。PaperCoder是一个多智能体（multi-agent）系统，可以把机器学习领域的...

**URL**: https://weibo.com/6105753431/Ppdb89O3C

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%BA%E6%96%87%E4%B8%80%E9%94%AE%E8%BD%AC%E4%BB%A3%E7%A0%81%23&amp;extparam=%23%E8%AE%BA%E6%96%87%E4%B8%80%E9%94%AE%E8%BD%AC%E4%BB%A3%E7%A0%81%23" data-hide=""><span class="surl-text">#论文一键转代码#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%BA%E6%96%87%E4%BB%A3%E7%A0%81%E5%A4%8D%E7%8E%B0%E5%B7%A5%E5%85%B7%23&amp;extparam=%23%E8%AE%BA%E6%96%87%E4%BB%A3%E7%A0%81%E5%A4%8D%E7%8E%B0%E5%B7%A5%E5%85%B7%23" data-hide=""><span class="surl-text">#论文代码复现工具#</span></a><br><br>PaperCoder：论文转代码工具，专治未公开代码的论文。<br><br>PaperCoder是一个多智能体（multi-agent）系统，可以把机器学习领域的论文，自动拆解、理解，最后生成一个可以直接跑的完整代码库。<br><br>整个过程遵循三个阶段【图1】，保证生成的代码既准确又可执行：<br><br>- Planning阶段：首先，PaperCoder会制定一个详细的开发路线图，画出系统架构图（比如UML图），整理出各模块之间的依赖关系，同时生成一份标准配置文件（config.yaml），为后续工作做准备。<br><br>- Analyzing阶段：接着，每一个即将生成的文件都会被深入分析，比如它需要实现什么功能，输入输出是什么，和其他模块有什么耦合，确保还原实验细节不走样。<br><br>- Coding阶段：最后，系统根据之前的规划和分析，顺序生成各模块代码，严格遵循依赖关系，保证仓库一生成就能基本无障碍运行。<br><br>下面来到大家最关心的问题——效果如何？  <br><br>在Paper2Code和PaperBench这两个公开基准测试上，PaperCoder表现出色。【图2】<br><br>无论是模型自动评分，还是MS/PhD学生的人类打分，它都大幅领先ChatDev和MetaGPT。<br><br>具体来看，PaperCoder的代码改动率只有0.48%，也就是说，平均每1000行代码只需要改动不到5行就能直接跑通。同时，77%的原论文作者更偏好PaperCoder生成的代码，85%的受访者认为它有效加速了研究复现。<br><br>虽然PaperCoder目前主要聚焦在机器学习领域，但未来计划扩展到更广泛的科学领域，比如物理、生物等。<br><br>网友测试后表示，把论文直接变成可用代码，PaperCoder是目前最接近「一键复现」的工具。<br><br>感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fgoing-doer%2FPaper2Code" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>抱抱脸：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fpapers%2F2504.17192" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0vh3qe5hkj31940to1i9.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0vh3rmgrqj319k0fshd2.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

PaperCoder是一款多智能体系统，可将机器学习论文自动转换为可执行代码库。其工作流程分为三个阶段：规划（生成架构图和配置文件）、分析（明确模块功能和依赖关系）和编码（按依赖顺序生成代码）。在Paper2Code和PaperBench基准测试中，其代码改动率仅0.48%，77%的原作者更偏好其生成结果，85%用户认为能加速研究复现。相比ChatDev和MetaGPT表现更优，目前专注机器学习领域，未来计划扩展至物理、生物等学科。该项目已在GitHub开源，被评价为当前最接近"一键复现"的工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-27T15:04:09Z
- **目录日期**: 2025-04-27
