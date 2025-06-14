# AUTOCIRCUIT-RL: Reinforcement Learning-Driven LLM for Automated Circuit Topology Generation

**URL**: http://arxiv.org/abs/2506.03122v1

## 原始摘要

Analog circuit topology synthesis is integral to Electronic Design Automation
(EDA), enabling the automated creation of circuit structures tailored to
specific design requirements. However, the vast design search space and strict
constraint adherence make efficient synthesis challenging. Leveraging the
versatility of Large Language Models (LLMs), we propose AUTOCIRCUIT-RL,a novel
reinforcement learning (RL)-based framework for automated analog circuit
synthesis. The framework operates in two phases: instruction tuning, where an
LLM learns to generate circuit topologies from structured prompts encoding
design constraints, and RL refinement, which further improves the
instruction-tuned model using reward models that evaluate validity, efficiency,
and output voltage. The refined model is then used directly to generate
topologies that satisfy the design constraints. Empirical results show that
AUTOCIRCUIT-RL generates ~12% more valid circuits and improves efficiency by
~14% compared to the best baselines, while reducing duplicate generation rates
by ~38%. It achieves over 60% success in synthesizing valid circuits with
limited training data, demonstrating strong generalization. These findings
highlight the framework's effectiveness in scaling to complex circuits while
maintaining efficiency and constraint adherence, marking a significant
advancement in AI-driven circuit design.


## AI 摘要

AUTOCIRCUIT-RL是一种基于强化学习（RL）和大型语言模型（LLM）的模拟电路拓扑自动合成框架。该框架分为指令调优和RL优化两阶段：LLM先学习根据设计约束生成电路拓扑，再通过奖励模型优化生成结果。实验表明，AUTOCIRCUIT-RL比基线方法多生成12%的有效电路，效率提升14%，重复率降低38%，且在有限训练数据下成功率超60%，展现出强大的泛化能力。该研究显著推进了AI驱动的电路设计，兼顾效率与约束满足。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T23:01:26Z
- **目录日期**: 2025-06-04
