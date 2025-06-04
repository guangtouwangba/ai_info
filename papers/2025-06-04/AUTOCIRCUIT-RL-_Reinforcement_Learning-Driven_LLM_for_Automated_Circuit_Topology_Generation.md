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

AUTOCIRCUIT-RL是一种基于强化学习（RL）的自动化模拟电路拓扑合成框架，利用大语言模型（LLM）分两阶段工作：指令调优学习根据设计约束生成电路拓扑，RL细化通过评估有效性、效率和输出电压的奖励模型进一步优化。相比最佳基线，该框架生成的有效电路多12%，效率提升14%，重复生成率降低38%，仅用有限训练数据即可实现超60%的有效电路合成成功率，展现出强大的泛化能力。这一成果标志着AI驱动电路设计在复杂电路扩展、效率与约束遵循方面的重大进展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T08:01:49Z
- **目录日期**: 2025-06-04
