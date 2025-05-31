# ML-Agent: Reinforcing LLM Agents for Autonomous Machine Learning Engineering

**URL**: http://arxiv.org/abs/2505.23723v1

## 原始摘要

The emergence of large language model (LLM)-based agents has significantly
advanced the development of autonomous machine learning (ML) engineering.
However, most existing approaches rely heavily on manual prompt engineering,
failing to adapt and optimize based on diverse experimental experiences.
Focusing on this, for the first time, we explore the paradigm of learning-based
agentic ML, where an LLM agent learns through interactive experimentation on ML
tasks using online reinforcement learning (RL). To realize this, we propose a
novel agentic ML training framework with three key components: (1)
exploration-enriched fine-tuning, which enables LLM agents to generate diverse
actions for enhanced RL exploration; (2) step-wise RL, which enables training
on a single action step, accelerating experience collection and improving
training efficiency; (3) an agentic ML-specific reward module, which unifies
varied ML feedback signals into consistent rewards for RL optimization.
Leveraging this framework, we train ML-Agent, driven by a 7B-sized Qwen-2.5 LLM
for autonomous ML. Remarkably, despite being trained on merely 9 ML tasks, our
7B-sized ML-Agent outperforms the 671B-sized DeepSeek-R1 agent. Furthermore, it
achieves continuous performance improvements and demonstrates exceptional
cross-task generalization capabilities.


## AI 摘要

这篇论文提出了一种基于强化学习(LLM)的自主机器学习(ML)智能体训练框架。该框架包含三个关键创新：(1)探索增强微调，提升智能体动作多样性；(2)分步强化学习，加速经验收集；(3)统一的ML任务奖励模块。研究人员使用7B参数的Qwen-2.5模型训练出ML-Agent，尽管仅在9个ML任务上训练，其性能却超过了671B参数的DeepSeek-R1智能体，并展现出持续的改进能力和出色的跨任务泛化性能。这一成果首次实现了基于学习的ML智能体范式，突破了传统依赖人工提示工程的局限。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-31T15:01:51Z
- **目录日期**: 2025-05-31
