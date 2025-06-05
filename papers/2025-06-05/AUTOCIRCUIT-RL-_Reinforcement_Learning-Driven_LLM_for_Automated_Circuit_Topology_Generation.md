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

AUTOCIRCUIT-RL是一种基于强化学习（RL）和大型语言模型（LLM）的模拟电路拓扑自动合成框架。该框架分两阶段：指令调优阶段，LLM学习根据设计约束生成电路拓扑；RL优化阶段，通过评估电路有效性、效率和输出电压的奖励模型进一步优化模型。实验表明，AUTOCIRCUIT-RL比基线方法多生成12%的有效电路，效率提升14%，重复生成率降低38%，并在有限训练数据下实现60%以上的有效电路合成成功率，展现了强大的泛化能力，推动了AI驱动的电路设计发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T01:29:08Z
- **目录日期**: 2025-06-05
