# #大模型原来是这样思考的##AI推理就是梯度下降#大模型“思考”过程首次被详细拆解，原来LLM在推理时，并不是简单输出答案，而是在“模拟”优化——就像在做梯度...

**URL**: https://weibo.com/6105753431/PvTNFqzgp

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%8E%9F%E6%9D%A5%E6%98%AF%E8%BF%99%E6%A0%B7%E6%80%9D%E8%80%83%E7%9A%84%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%8E%9F%E6%9D%A5%E6%98%AF%E8%BF%99%E6%A0%B7%E6%80%9D%E8%80%83%E7%9A%84%23" data-hide=""><span class="surl-text">#大模型原来是这样思考的#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%8E%A8%E7%90%86%E5%B0%B1%E6%98%AF%E6%A2%AF%E5%BA%A6%E4%B8%8B%E9%99%8D%23&amp;extparam=%23AI%E6%8E%A8%E7%90%86%E5%B0%B1%E6%98%AF%E6%A2%AF%E5%BA%A6%E4%B8%8B%E9%99%8D%23" data-hide=""><span class="surl-text">#AI推理就是梯度下降#</span></a><br><br>大模型“思考”过程首次被详细拆解，原来LLM在推理时，并不是简单输出答案，而是在“模拟”优化——就像在做梯度下降。<br><br>RaML团队用QwQ-32B模型实测，发现随着推理令牌增加，模型对正确答案的置信度逐步上升。这不仅证明了推理轨迹的“优化”性质，也说明模型确实在“思考”。<br><br>他们进一步用元学习的视角解释训练过程：每个问题是一个任务，推理轨迹是内循环优化，而整体训练过程是外循环优化。SFT和RL的差异，也可以理解为优化路径的差别——SFT借助外部“最佳梯度”引导更快收敛，尤其适合中小模型；而RL探索空间更广，理论潜力更高。<br><br>“反思”令牌像是在做优化重启，有助跳出局部最优。相反，那些强行终止“思考”的提示词，虽然节省算力，但容易让模型陷入鞍点，影响最终效果。<br><br>有趣的是，团队发现：只用数学推理训练的大模型，在科学和代码推理任务上同样表现出色。这说明，LLM确实学到了“通用推理逻辑”。 <a href="https://weibo.com/ttarticle/p/show?id=2309405176026673578363" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">揭秘LLM“思考”之谜：推理即“梯度下降”，元学习框架解构训练过程</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2a8ui6rboj30rs0fm0v1.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

RaML团队研究发现，大语言模型(LLM)的推理过程类似梯度下降优化：随着推理令牌增加，模型对正确答案的置信度逐步提升。研究用元学习框架解释训练过程 - 将每个问题视为任务，推理轨迹为内循环优化，整体训练为外循环优化。SFT和RL的区别在于优化路径不同，其中RL探索空间更广。研究还发现"反思"令牌有助于跳出局部最优，而数学推理训练可提升其他领域的推理能力，表明LLM掌握了通用推理逻辑。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T09:04:01Z
- **目录日期**: 2025-06-10
