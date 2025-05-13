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

IVE是一个受人类好奇心启发的机器人探索框架，通过视觉语言模型（VLM）将RGB-D观测抽象为语义场景图，想象新场景并预测其物理可行性，生成可执行动作序列。该框架在模拟和真实桌面环境中验证，相比强化学习基线，访问状态的多样性显著提升（熵值增加4.1-7.8倍），且收集的经验支持下游任务学习，其策略性能接近或超过基于人类示范训练的结果。IVE有效解决了VLM输出未落地的问题，实现了语义推理与物理执行的结合，推动了开放环境中无密集奖励的通用机器人学习。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T10:00:59Z
- **目录日期**: 2025-05-13
