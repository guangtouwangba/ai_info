# #强化学习新范式##阿里高德重构强化学习#阿里高德提出了强化学习训练新范式：只优化原始目标，不用替代损失函数！团队发布了全新的组策略梯度优化方法GPG，从底...

**URL**: https://weibo.com/6105753431/Pp1VQuCNQ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E6%96%B0%E8%8C%83%E5%BC%8F%23&amp;extparam=%23%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E6%96%B0%E8%8C%83%E5%BC%8F%23" data-hide=""><span class="surl-text">#强化学习新范式#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%98%BF%E9%87%8C%E9%AB%98%E5%BE%B7%E9%87%8D%E6%9E%84%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%23&amp;extparam=%23%E9%98%BF%E9%87%8C%E9%AB%98%E5%BE%B7%E9%87%8D%E6%9E%84%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%23" data-hide=""><span class="surl-text">#阿里高德重构强化学习#</span></a><br><br>阿里高德提出了强化学习训练新范式：只优化原始目标，不用替代损失函数！<br><br>团队发布了全新的组策略梯度优化方法GPG，从底层重构强化学习框架，革新训练流程。核心亮点有三：<br><br>- 直接目标优化：摒弃传统替代损失函数设计，突破训练效率瓶颈；<br><br>- 极简训练架构：消除评论模型和参考模型，摆脱KL散度约束；<br><br>- 精准梯度估计（AGE）：首度揭示奖励偏差，提升策略稳定性。<br><br>GPG在单模态和多模态任务中表现优异，超越现有方法。尤其在数学推理和视觉理解等复杂任务中，展示了极高的通用性和鲁棒性。<br><br>方法细节上，GPG利用组内平均奖励归一化，降低方差，移除价值模型。特别是针对组内样本全对全错带来的梯度偏差，引入了动态梯度校正机制，有效提升训练稳定性。详情请见文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405159656308998242" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">从底层重构强化学习训练框架，阿里高德开源新方法：抛弃替代损失函数</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0u3g8vsafj30ls0c9myd.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

阿里高德提出强化学习新范式GPG，直接优化原始目标而无需替代损失函数。该方法突破传统框架，消除评论模型和KL约束，首创精准梯度估计(AGE)技术解决奖励偏差问题。通过组内奖励归一化和动态梯度校正，显著提升训练效率和稳定性。实验显示GPG在数学推理、视觉理解等复杂任务中表现优异，展现出强大的通用性和鲁棒性。这一创新为强化学习训练流程带来根本性变革。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-26T22:03:19Z
- **目录日期**: 2025-04-26
