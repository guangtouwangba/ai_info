# Keep Everyone Happy: Online Fair Division of Numerous Items with Few Copies

**URL**: http://arxiv.org/abs/2408.12845v2

## 原始摘要

This paper considers a novel variant of the online fair division problem
involving multiple agents in which a learner sequentially observes an
indivisible item that has to be irrevocably allocated to one of the agents
while satisfying a fairness and efficiency constraint. Existing algorithms
assume a small number of items with a sufficiently large number of copies,
which ensures a good utility estimation for all item-agent pairs from noisy
bandit feedback. However, this assumption may not hold in many real-life
applications, for example, an online platform that has a large number of users
(items) who use the platform's service providers (agents) only a few times (a
few copies of items), which makes it difficult to accurately estimate utilities
for all item-agent pairs. To address this, we assume utility is an unknown
function of item-agent features. We then propose algorithms that model online
fair division as a contextual bandit problem, with sub-linear regret
guarantees. Our experimental results further validate the effectiveness of the
proposed algorithms.


## AI 摘要

这篇论文提出了一种新的在线公平分配问题变体，研究如何在物品数量大但重复次数少的情况下，通过上下文老虎机方法解决效用估计难题。传统算法假设物品数量少且重复次数多，但现实中（如在线平台）往往用户多而使用次数少，导致效用估计不准。为此，作者假设效用是物品-代理特征的未知函数，将问题建模为上下文老虎机，并提出具有次线性遗憾保证的算法。实验验证了算法的有效性。该研究为稀疏交互场景下的公平资源分配提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-31T03:19:57Z
- **目录日期**: 2025-05-31
