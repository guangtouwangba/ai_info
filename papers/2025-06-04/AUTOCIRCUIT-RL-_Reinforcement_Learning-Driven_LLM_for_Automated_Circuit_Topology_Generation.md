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

本文提出了一种基于强化学习(LLM)的模拟电路自动综合框架AUTOCIRCUIT-RL。该框架分两阶段：首先通过指令调优使LLM学习根据设计约束生成电路拓扑，然后利用评估有效性、效率和输出电压的奖励模型进行强化学习优化。实验表明，与基线相比，该框架能多生成12%的有效电路，效率提升14%，重复生成率降低38%。仅用有限训练数据就能实现60%以上的有效电路综合成功率，展现出优异的泛化能力。这一成果标志着AI驱动电路设计的重要进展，能高效处理复杂电路同时满足设计约束。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T19:01:32Z
- **目录日期**: 2025-06-04
