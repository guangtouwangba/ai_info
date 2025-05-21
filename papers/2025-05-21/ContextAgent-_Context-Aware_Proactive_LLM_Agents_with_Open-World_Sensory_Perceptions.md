# ContextAgent: Context-Aware Proactive LLM Agents with Open-World Sensory Perceptions

**URL**: http://arxiv.org/abs/2505.14668v1

## 原始摘要

Recent advances in Large Language Models (LLMs) have propelled intelligent
agents from reactive responses to proactive support. While promising, existing
proactive agents either rely exclusively on observations from enclosed
environments (e.g., desktop UIs) with direct LLM inference or employ rule-based
proactive notifications, leading to suboptimal user intent understanding and
limited functionality for proactive service. In this paper, we introduce
ContextAgent, the first context-aware proactive agent that incorporates
extensive sensory contexts to enhance the proactive capabilities of LLM agents.
ContextAgent first extracts multi-dimensional contexts from massive sensory
perceptions on wearables (e.g., video and audio) to understand user intentions.
ContextAgent then leverages the sensory contexts and the persona contexts from
historical data to predict the necessity for proactive services. When proactive
assistance is needed, ContextAgent further automatically calls the necessary
tools to assist users unobtrusively. To evaluate this new task, we curate
ContextAgentBench, the first benchmark for evaluating context-aware proactive
LLM agents, covering 1,000 samples across nine daily scenarios and twenty
tools. Experiments on ContextAgentBench show that ContextAgent outperforms
baselines by achieving up to 8.5% and 6.0% higher accuracy in proactive
predictions and tool calling, respectively. We hope our research can inspire
the development of more advanced, human-centric, proactive AI assistants.


## AI 摘要

本文提出ContextAgent，首个结合多维度感知数据的上下文感知主动AI助手。该系统通过可穿戴设备获取视频、音频等多模态感知数据，结合用户历史行为特征，预测主动服务需求并自动调用工具。研究者构建了首个上下文感知主动代理基准ContextAgentBench（覆盖9个日常场景和20种工具的1000个样本）。实验表明，ContextAgent在主动预测和工具调用准确率上分别比基线高8.5%和6.0%。该研究为开发更先进、以人为中心的主动AI助手提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T03:19:11Z
- **目录日期**: 2025-05-21
