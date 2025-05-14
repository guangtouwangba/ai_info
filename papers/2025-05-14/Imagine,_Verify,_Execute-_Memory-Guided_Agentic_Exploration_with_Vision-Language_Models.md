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

IVE是一个基于视觉语言模型（VLM）的机器人探索框架，受人类好奇心启发，通过“想象-验证-执行”机制驱动自主探索。它将RGB-D观测抽象为语义场景图，预测新场景的物理可行性，并生成可执行动作序列。实验表明，IVE在模拟和真实桌面环境中比强化学习基线探索更丰富（状态熵提升4.1-7.8倍），且收集的经验能有效支持下游任务学习，其训练的策略性能接近或超越基于人类示范数据训练的模型。该框架解决了VLM输出未接地的问题，实现了语义想象与物理执行的结合。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T02:29:49Z
- **目录日期**: 2025-05-14
