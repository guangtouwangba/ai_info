# #AI学会选择性思考##让AI在该思考时思考#人类想太多会容易秃顶，AI想太多会怎么样呢？答案显然易见：用复杂的推理来解决所有问题，计算效率会变得极其低下，同时...

**URL**: https://weibo.com/6105753431/PsPpoADdH

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%AD%A6%E4%BC%9A%E9%80%89%E6%8B%A9%E6%80%A7%E6%80%9D%E8%80%83%23&amp;extparam=%23AI%E5%AD%A6%E4%BC%9A%E9%80%89%E6%8B%A9%E6%80%A7%E6%80%9D%E8%80%83%23" data-hide=""><span class="surl-text">#AI学会选择性思考#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%A9AI%E5%9C%A8%E8%AF%A5%E6%80%9D%E8%80%83%E6%97%B6%E6%80%9D%E8%80%83%23&amp;extparam=%23%E8%AE%A9AI%E5%9C%A8%E8%AF%A5%E6%80%9D%E8%80%83%E6%97%B6%E6%80%9D%E8%80%83%23" data-hide=""><span class="surl-text">#让AI在该思考时思考#</span></a><br><br>人类想太多会容易秃顶，AI想太多会怎么样呢？<br><br>答案显然易见：用复杂的推理来解决所有问题，计算效率会变得极其低下，同时成本不停upupup。<br><br>那么，有没有什么方法让AI学会在恰当的时候进行思考呢？有的，朋友，有的。<br><br>来自新加坡国立大学的研究团队提出了一种名为Thinkless的框架，让AI学会了根据任务复杂度和模型能力，自动选择简短回答或长链推理。【图2】<br><br>具体怎么做到的呢？Thinkless可以分为两个阶段：蒸馏预热和利用DeGRPO进行强化学习。<br><br>蒸馏预热阶段通过两个专家模型进行知识蒸馏，一个用于生成详细的推理链，另一个用于生成简洁答案，使目标模型掌握两种回答风格。<br><br>利用DeGRPO进行强化学习的第二个阶段是这个框架的核心阶段。它将混合推理的学习目标分解为两个部分：<br><br>- 控制标记损失函数：用于推理模式的选择<br>- 回答损失函数：用于提高生成答案的准确性<br><br>通过解耦的方式平衡两个目标的贡献，稳定训练过程，有效防止了在普通GRPO中观察到的崩溃。<br><br>Thinkless框架在Minerva Algebra、MATH-500、GSM8K等基准测试中都表现出色。结果表明该框架能够显著减少长链推理的使用，提高推理语言模型的效率。【图3】<br><br>论文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.13379" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>开源仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2FVainF%2FThinkless" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1mz0x1cz4j30q40zk7hz.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1mz0yfiyaj30zk0e6wm1.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1mz12ijc7j30zk0ddtfn.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

新加坡国立大学团队提出Thinkless框架，使AI能根据任务复杂度自动选择简短回答或长链推理。该框架分两阶段：1)蒸馏预热阶段通过两个专家模型（详细推理/简洁回答）进行知识蒸馏；2)强化学习阶段使用DeGRPO算法，解耦控制标记损失（选择推理模式）和回答损失（提高准确性），避免训练崩溃。实验显示，Thinkless在Minerva Algebra等测试中显著减少不必要的长链推理，提升效率。论文和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T07:03:20Z
- **目录日期**: 2025-05-21
