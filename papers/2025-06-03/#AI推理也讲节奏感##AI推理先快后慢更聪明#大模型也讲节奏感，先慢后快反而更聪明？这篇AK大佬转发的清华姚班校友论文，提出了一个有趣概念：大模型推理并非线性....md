# #AI推理也讲节奏感##AI推理先快后慢更聪明#大模型也讲节奏感，先慢后快反而更聪明？这篇AK大佬转发的清华姚班校友论文，提出了一个有趣概念：大模型推理并非线性...

**URL**: https://weibo.com/6105753431/PuOYZjffh

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%8E%A8%E7%90%86%E4%B9%9F%E8%AE%B2%E8%8A%82%E5%A5%8F%E6%84%9F%23&amp;extparam=%23AI%E6%8E%A8%E7%90%86%E4%B9%9F%E8%AE%B2%E8%8A%82%E5%A5%8F%E6%84%9F%23" data-hide=""><span class="surl-text">#AI推理也讲节奏感#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%8E%A8%E7%90%86%E5%85%88%E5%BF%AB%E5%90%8E%E6%85%A2%E6%9B%B4%E8%81%AA%E6%98%8E%23&amp;extparam=%23AI%E6%8E%A8%E7%90%86%E5%85%88%E5%BF%AB%E5%90%8E%E6%85%A2%E6%9B%B4%E8%81%AA%E6%98%8E%23" data-hide=""><span class="surl-text">#AI推理先快后慢更聪明#</span></a><br><br>大模型也讲节奏感，先慢后快反而更聪明？<br><br>这篇AK大佬转发的清华姚班校友论文，提出了一个有趣概念：大模型推理并非线性展开，而是一个节奏变化的动态过程。<br><br>基于此，研究团队提出了一种新方法——AlphaOne (α1)，能调控大模型推理时的“慢思考”和“快思考”。<br><br>这项方法的核心观点是：<br><br>AlphaOne方法引入一个参数α，通过它控制模型何时“慢下来”，何时“提速”。整个流程可以拆解为两段：<br><br>1. 慢速阶段（前期）：模型在生成答案时会插入“wait” token，相当于给自己一些时间进行更深层次的思考。这些token不是固定插入，而是根据一个动态策略（如线性退火）来逐步减少频率。<br><br>2. 快速阶段（后期）：到了设定的“α时刻”，系统自动切换到快节奏，停止等待并用特定token（比如“&lt;/think&gt;”）终结深度推理，输出答案。这相当于把“磨刀时间”集中在最开始，后半段迅速收尾。<br><br>这一方法带来多方面优势：<br><br>- 效率更高：与传统“全程快速”或“随机快慢”策略相比，α1在保持准确率的同时，显著减少了推理所需的token数量，平均节省14%；<br><br>- 准确率提升：在多个高强度推理任务中（如数学竞赛AIME24、AMC23，代码生成LiveCode，以及科学竞赛Olympiad），α1表现优于现有方法，如s1（线性加速策略）和CoD（递减节奏策略），准确率提升6%-13%；<br><br>- 泛化能力强：无论是1.5B还是32B规模的模型，α1都适用，不依赖具体模型架构；<br><br>- FREP指标领先：研究团队提出一个衡量“推理效率与表现平衡”的新指标FREP（Reasoning Efficiency-Performance），α1在各项测试中得分最高。<br><br>此外，研究还指出几个容易被忽视的细节：<br><br>- Post-α modulation不可或缺：如果不强制结束“wait”，模型会陷入“过度深思”状态，导致性能下降，实验中甚至损失高达20%的表现；<br><br>- 节奏调控不能太极端：插入等待token太少或太多都会适得其反，α不是越大越好，一个合适的α值能最大化准确率和token效率。<br><br>总的来看，α1用极简单的方式，将复杂推理流程划分为“慢-快”两个阶段，在不需重新训练、不依赖外部工具的前提下，就能显著增强模型智能。<br><br>顺带一提，论文一作是Runpei Dong博士，他目前就读于伊利诺伊大学香槟分校，在此之前，他于2024年获得了计算机科学硕士学位，期间在清华大学交叉信息研究院（IIIS）和西安交通大学联合培养，由姚期智院士指导。<br><br>完整论文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.24863" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i224kva0n2j314o16w7s4.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i224kvvw04j30xb0mona5.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

清华大学团队提出AlphaOne方法，通过动态调控AI推理节奏提升性能。该方法将推理过程分为两个阶段：前期插入"wait"标记进行深度思考（慢速阶段），后期快速输出结果（快速阶段）。实验显示，这种"先慢后快"的策略在数学竞赛、代码生成等任务中准确率提升6-13%，同时减少14%的token消耗。该方法适用于不同规模模型，无需重新训练。研究还发现过度思考会降低性能，需合理设置转换时机。论文由姚期智院士指导的Runpei Dong博士领衔，成果发表于arXiv。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T07:03:18Z
- **目录日期**: 2025-06-03
