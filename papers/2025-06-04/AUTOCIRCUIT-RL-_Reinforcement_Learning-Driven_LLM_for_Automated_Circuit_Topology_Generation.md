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

AUTOCIRCUIT-RL是一种基于强化学习（RL）和大型语言模型（LLM）的模拟电路拓扑自动合成框架。该框架分为指令调优和RL优化两阶段：LLM首先学习根据设计约束生成电路拓扑，再通过奖励模型评估有效性、效率和输出电压以进一步优化模型。实验表明，相比基线方法，AUTOCIRCUIT-RL生成的电路有效性提高12%，效率提升14%，重复率降低38%，并在有限训练数据下实现60%以上的有效电路合成成功率。该框架在保持高效和满足约束的同时，可扩展至复杂电路设计，推动了AI驱动电路设计的进步。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T14:01:51Z
- **目录日期**: 2025-06-04
