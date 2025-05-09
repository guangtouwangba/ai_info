# EcoAgent: An Efficient Edge-Cloud Collaborative Multi-Agent Framework for Mobile Automation

**URL**: http://arxiv.org/abs/2505.05440v1

## 原始摘要

Cloud-based mobile agents powered by (multimodal) large language models
((M)LLMs) offer strong reasoning abilities but suffer from high latency and
cost. While fine-tuned (M)SLMs enable edge deployment, they often lose general
capabilities and struggle with complex tasks. To address this, we propose
EcoAgent, an Edge-Cloud cOllaborative multi-agent framework for mobile
automation. EcoAgent features a closed-loop collaboration among a cloud-based
Planning Agent and two edge-based agents: the Execution Agent for action
execution and the Observation Agent for verifying outcomes. The Observation
Agent uses a Pre-Understanding Module to compress screen images into concise
text, reducing token usage. In case of failure, the Planning Agent retrieves
screen history and replans via a Reflection Module. Experiments on AndroidWorld
show that EcoAgent maintains high task success rates while significantly
reducing MLLM token consumption, enabling efficient and practical mobile
automation.


## AI 摘要

EcoAgent是一个边缘-云协同的多智能体框架，旨在解决云端大型语言模型（(M)LLMs）的高延迟和高成本问题，同时克服边缘端轻量模型（(M)SLMs）通用能力不足的缺陷。该框架包含云端规划智能体和两个边缘智能体：执行智能体负责操作，观察智能体通过预理解模块压缩屏幕图像为简洁文本以减少token消耗。若任务失败，规划智能体通过反射模块重新规划。实验表明，EcoAgent在保持高任务成功率的同时显著降低了token消耗，实现了高效实用的移动自动化。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T08:01:51Z
- **目录日期**: 2025-05-09
