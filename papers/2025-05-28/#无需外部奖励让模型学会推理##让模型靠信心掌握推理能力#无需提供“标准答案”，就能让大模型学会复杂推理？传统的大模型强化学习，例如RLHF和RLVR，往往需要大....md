# #无需外部奖励让模型学会推理##让模型靠信心掌握推理能力#无需提供“标准答案”，就能让大模型学会复杂推理？传统的大模型强化学习，例如RLHF和RLVR，往往需要大...

**URL**: https://weibo.com/6105753431/PtTBzvFiW

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%97%A0%E9%9C%80%E5%A4%96%E9%83%A8%E5%A5%96%E5%8A%B1%E8%AE%A9%E6%A8%A1%E5%9E%8B%E5%AD%A6%E4%BC%9A%E6%8E%A8%E7%90%86%23&amp;extparam=%23%E6%97%A0%E9%9C%80%E5%A4%96%E9%83%A8%E5%A5%96%E5%8A%B1%E8%AE%A9%E6%A8%A1%E5%9E%8B%E5%AD%A6%E4%BC%9A%E6%8E%A8%E7%90%86%23" data-hide=""><span class="surl-text">#无需外部奖励让模型学会推理#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%A9%E6%A8%A1%E5%9E%8B%E9%9D%A0%E4%BF%A1%E5%BF%83%E6%8E%8C%E6%8F%A1%E6%8E%A8%E7%90%86%E8%83%BD%E5%8A%9B%23&amp;extparam=%23%E8%AE%A9%E6%A8%A1%E5%9E%8B%E9%9D%A0%E4%BF%A1%E5%BF%83%E6%8E%8C%E6%8F%A1%E6%8E%A8%E7%90%86%E8%83%BD%E5%8A%9B%23" data-hide=""><span class="surl-text">#让模型靠信心掌握推理能力#</span></a><br><br>无需提供“标准答案”，就能让大模型学会复杂推理？<br><br>传统的大模型强化学习，例如RLHF和RLVR，往往需要大量的外部奖励信号或标注数据来指导训练。<br><br>这些方法容易引入偏差，而且对特定领域验证器的依赖，限制了它们在更广泛场景中的应用。<br><br>更重要的是，这些依赖结果正确性的方法往往忽略了推理过程的质量，导致在不同领域间的泛化能力不足。<br><br>来自加州大学伯克利分校的研究团队提出一种基于内部反馈的强化学习框架（RLIF），让模型依赖内在的“信心”，学会复杂推理。【图1】<br><br>具体而言，他们的实现路径可以分为两个部分：<br><br>1. INTUITOR方法：利用模型自身的置信度作为唯一的内在奖励信号，通过不断优化这种置信度来提升模型的推理能力。这种方法不仅能提高模型在特定任务上的表现，还能显著增强模型的泛化能力。【图3】<br><br>2. 基于GRPO的实现：研究团队通过替换群组相对策略优化（GRPO）中的外部奖励机制，实现了完全无监督学习。这意味着模型不再需要依赖外部的判断或标准答案，而是依靠自身的内在评估来进行学习和提升。<br><br>RLIF框架的提出，带来了多重优势。它不仅减少了对监督基础设施的依赖，还提供了一种与任务无关的奖励信号，使得模型能够在外部验证不可用的领域中进行学习。<br><br>通过在数学推理、代码生成和指令遵循等任务上的实验，研究团队验证了INTUITOR方法的有效性。<br><br>实验结果表明，INTUITOR在数学推理任务上与GRPO表现相当，同时在代码生成等跨领域任务上展现出更好的泛化能力。【图4】<br><br>模型能够完全依赖内在反馈学会结构化推理，包括提前规划、分解问题，甚至遵循指令。<br><br>这项研究证实，模型内在信号可以驱动跨领域的高效学习，为缺乏可验证奖励的自主AI系统提供了一种可扩展的RLVR替代方案。<br><br>论文地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.19590" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>代码仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fsunblaze-ucb%2FIntuitor" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1v39k68wej30zk0to49n.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1v39ltroaj30qg0duad9.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1v39ozi1fj30zk08jdia.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1v39r1zwzj30i00dwabx.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

加州大学伯克利分校团队提出基于内部反馈的强化学习框架RLIF，通过INTUITOR方法让大模型利用自身置信度作为内在奖励信号，无需外部标准答案即可学习复杂推理。该框架采用GRPO实现无监督学习，减少对外部验证的依赖，在数学推理、代码生成等任务中展现出与监督方法相当的表现和更好的跨领域泛化能力。研究表明模型能通过内在信号自主掌握结构化推理能力，为缺乏可验证奖励的场景提供了可扩展的解决方案。论文和代码已公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T05:02:48Z
- **目录日期**: 2025-05-28
