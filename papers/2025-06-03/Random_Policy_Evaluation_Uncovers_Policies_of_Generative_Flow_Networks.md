# Random Policy Evaluation Uncovers Policies of Generative Flow Networks

**URL**: http://arxiv.org/abs/2406.02213v3

## 原始摘要

The Generative Flow Network (GFlowNet) is a probabilistic framework in which
an agent learns a stochastic policy and flow functions to sample objects
proportionally to an unnormalized reward function. A number of recent works
explored connections between GFlowNets and maximum entropy (MaxEnt) RL, which
modifies the standard objective of RL agents by learning an entropy-regularized
objective. However, the relationship between GFlowNets and standard RL remains
largely unexplored, despite the inherent similarities in their sequential
decision-making nature. While GFlowNets can discover diverse solutions through
specialized flow-matching objectives, connecting them can simplify their
implementation through established RL principles and improve RL's diverse
solution discovery capabilities. In this paper, we bridge this gap by revealing
a fundamental connection between GFlowNets and one RL's most basic components
-- policy evaluation. Surprisingly, we find that the value function obtained
from evaluating a uniform policy is closely associated with the flow functions
in GFlowNets through the lens of flow iteration under certain structural
conditions. Building upon these insights, we introduce a rectified random
policy evaluation (RPE) algorithm, which achieves the same reward-matching
effect as GFlowNets based on simply evaluating a fixed random policy in these
cases, offering a new perspective. Empirical results across extensive
benchmarks demonstrate that RPE achieves competitive results compared to
previous approaches, shedding light on the previously overlooked connection
between (non-MaxEnt) RL and GFlowNets.


## AI 摘要

这篇论文揭示了生成流网络(GFlowNet)与强化学习(RL)中策略评估的基本联系。研究发现，在特定结构条件下，通过评估均匀策略获得的价值函数与GFlowNets中的流函数密切相关。基于这一发现，作者提出了修正随机策略评估(RPE)算法，该算法通过评估固定随机策略即可实现与GFlowNets相同的奖励匹配效果。实验结果表明，RPE在多个基准测试中取得了与现有方法相当的性能。这一发现为理解(非最大熵)RL与GFlowNets之间被忽视的联系提供了新视角。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T07:02:43Z
- **目录日期**: 2025-06-03
