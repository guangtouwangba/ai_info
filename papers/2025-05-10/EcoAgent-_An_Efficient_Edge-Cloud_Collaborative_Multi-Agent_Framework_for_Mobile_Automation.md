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

EcoAgent是一个边缘-云协作的多智能体框架，旨在解决基于云的大语言模型（LLM）高延迟、高成本的问题，同时弥补边缘端小模型（SLM）通用能力不足的缺陷。该框架包含云端的规划智能体和边缘端的执行、观察智能体，形成闭环协作：执行智能体操作设备，观察智能体通过图像压缩模块将屏幕信息转为简洁文本以节省token。若任务失败，规划智能体通过反射模块调取历史记录重新规划。在AndroidWorld测试中，EcoAgent在保持高成功率的同时显著降低了token消耗，实现了高效的移动自动化。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-10T17:01:43Z
- **目录日期**: 2025-05-10
