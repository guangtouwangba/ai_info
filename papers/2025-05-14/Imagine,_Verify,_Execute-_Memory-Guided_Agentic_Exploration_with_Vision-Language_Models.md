# Imagine, Verify, Execute: Memory-Guided Agentic Exploration with Vision-Language Models

**URL**: http://arxiv.org/abs/2505.07815v1

## 原始摘要

Exploration is essential for general-purpose robotic learning, especially in
open-ended environments where dense rewards, explicit goals, or task-specific
supervision are scarce. Vision-language models (VLMs), with their semantic
reasoning over objects, spatial relations, and potential outcomes, present a
compelling foundation for generating high-level exploratory behaviors. However,
their outputs are often ungrounded, making it difficult to determine whether
imagined transitions are physically feasible or informative. To bridge the gap
between imagination and execution, we present IVE (Imagine, Verify, Execute),
an agentic exploration framework inspired by human curiosity. Human exploration
is often driven by the desire to discover novel scene configurations and to
deepen understanding of the environment. Similarly, IVE leverages VLMs to
abstract RGB-D observations into semantic scene graphs, imagine novel scenes,
predict their physical plausibility, and generate executable skill sequences
through action tools. We evaluate IVE in both simulated and real-world tabletop
environments. The results show that IVE enables more diverse and meaningful
exploration than RL baselines, as evidenced by a 4.1 to 7.8x increase in the
entropy of visited states. Moreover, the collected experience supports
downstream learning, producing policies that closely match or exceed the
performance of those trained on human-collected demonstrations.


## AI 摘要

IVE是一个受人类好奇心启发的机器人探索框架，利用视觉语言模型（VLM）将RGB-D观测转化为语义场景图，想象新场景并预测其物理可行性，通过动作工具生成可执行技能序列。在模拟和真实桌面环境中，IVE展现出比强化学习基线更丰富、更有意义的探索能力，访问状态的熵值提高了4.1至7.8倍。此外，IVE收集的经验支持下游学习，其训练的策略性能接近或超过基于人类示范数据训练的模型。该框架有效弥合了想象与执行之间的差距，适用于开放环境中稀疏奖励或无明确目标的任务。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T00:01:06Z
- **目录日期**: 2025-05-14
