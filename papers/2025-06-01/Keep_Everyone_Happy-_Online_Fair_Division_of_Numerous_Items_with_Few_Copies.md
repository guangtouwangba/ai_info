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

本文提出了一种新型在线公平分配问题的变体，针对多智能体场景下不可分割物品的序贯分配问题。传统算法假设物品数量少且副本充足，但在实际应用中（如在线平台），用户（物品）多而服务使用（副本）少，导致效用估计困难。为此，作者将效用建模为物品-智能体特征的未知函数，将问题转化为上下文赌博机问题，并提出具有次线性遗憾保证的算法。实验结果验证了算法的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-01T12:02:20Z
- **目录日期**: 2025-06-01
