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

该研究提出了一种基于强化学习(RL)的自主机器学习(ML)代理框架，通过三个创新组件实现：1)探索增强微调，提升LLM代理生成多样化动作的能力；2)分步RL训练，加速经验收集；3)统一奖励模块，整合不同ML反馈信号。使用7B参数的Qwen-2.5模型训练的ML-Agent，仅通过9个ML任务训练，性能即超越671B参数的DeepSeek-R1代理，并展现出持续的改进能力和出色的跨任务泛化性能。这一学习型代理范式突破了传统依赖人工提示工程的局限，实现了通过交互实验自主优化的ML代理。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-01T05:02:21Z
- **目录日期**: 2025-06-01
