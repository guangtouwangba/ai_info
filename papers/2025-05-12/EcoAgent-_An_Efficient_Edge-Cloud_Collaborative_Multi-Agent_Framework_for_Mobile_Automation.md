# EcoAgent: An Efficient Edge-Cloud Collaborative Multi-Agent Framework for Mobile Automation

**URL**: http://arxiv.org/abs/2505.05440v2

## 原始摘要

Cloud-based mobile agents powered by (multimodal) large language models
((M)LLMs) offer strong reasoning abilities but suffer from high latency and
cost. While fine-tuned (M)SLMs enable edge deployment, they often lose general
capabilities and struggle with complex tasks. To address this, we propose
\textbf{EcoAgent}, an \textbf{E}dge-\textbf{C}loud c\textbf{O}llaborative
multi-agent framework for mobile automation. EcoAgent features a closed-loop
collaboration among a cloud-based Planning Agent and two edge-based agents: the
Execution Agent for action execution and the Observation Agent for verifying
outcomes. The Observation Agent uses a Pre-Understanding Module to compress
screen images into concise text, reducing token usage and communication
overhead. In case of failure, the Planning Agent retrieves screen history
through a Memory Module and replans via a Reflection Module. Experiments on
AndroidWorld show that EcoAgent achieves task success rates comparable to
cloud-based mobile agents while significantly reducing MLLM token consumption,
enabling efficient and practical mobile automation.


## AI 摘要

EcoAgent是一个边缘-云协作的多智能体框架，旨在解决基于云的大语言模型(M)LLM移动代理的高延迟和高成本问题。它包含一个云端的规划智能体和两个边缘端的智能体：执行智能体负责动作执行，观察智能体通过预理解模块将屏幕图像压缩为简洁文本来验证结果。若任务失败，规划智能体通过记忆模块检索屏幕历史，并通过反思模块重新规划。在AndroidWorld上的实验表明，EcoAgent在保持与云端代理相当任务成功率的同时，显著减少了MLLM的token消耗，实现了高效实用的移动自动化。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T01:29:32Z
- **目录日期**: 2025-05-12
