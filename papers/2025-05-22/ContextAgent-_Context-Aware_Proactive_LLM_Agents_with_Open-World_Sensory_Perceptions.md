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

本文提出ContextAgent，首个结合多维度感知情境的主动式LLM智能体。它通过可穿戴设备（如视频/音频）提取环境感知数据理解用户意图，结合历史行为画像预测主动服务需求，并自动调用工具提供无干扰协助。研究者构建了首个情境感知主动代理基准ContextAgentBench（含9类日常场景的1000个样本和20种工具）。实验表明，该模型在主动预测和工具调用准确率上分别提升8.5%和6.0%，为开发更先进、以人为中心的主动AI助手提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T01:29:01Z
- **目录日期**: 2025-05-22
