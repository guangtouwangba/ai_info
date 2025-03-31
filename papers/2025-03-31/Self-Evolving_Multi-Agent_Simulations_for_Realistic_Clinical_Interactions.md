# Self-Evolving Multi-Agent Simulations for Realistic Clinical Interactions

**URL**: http://arxiv.org/abs/2503.22678v1

## 原始摘要

In this work, we introduce MedAgentSim, an open-source simulated clinical
environment with doctor, patient, and measurement agents designed to evaluate
and enhance LLM performance in dynamic diagnostic settings. Unlike prior
approaches, our framework requires doctor agents to actively engage with
patients through multi-turn conversations, requesting relevant medical
examinations (e.g., temperature, blood pressure, ECG) and imaging results
(e.g., MRI, X-ray) from a measurement agent to mimic the real-world diagnostic
process. Additionally, we incorporate self improvement mechanisms that allow
models to iteratively refine their diagnostic strategies. We enhance LLM
performance in our simulated setting by integrating multi-agent discussions,
chain-of-thought reasoning, and experience-based knowledge retrieval,
facilitating progressive learning as doctor agents interact with more patients.
We also introduce an evaluation benchmark for assessing the LLM's ability to
engage in dynamic, context-aware diagnostic interactions. While MedAgentSim is
fully automated, it also supports a user-controlled mode, enabling human
interaction with either the doctor or patient agent. Comprehensive evaluations
in various simulated diagnostic scenarios demonstrate the effectiveness of our
approach. Our code, simulation tool, and benchmark are available at
\href{https://medagentsim.netlify.app/}.


## AI 摘要

本研究提出了MedAgentSim，一个开源的临床模拟环境，包含医生、患者和检测代理，用于评估和改进大语言模型(LLM)在动态诊断中的表现。该框架要求医生代理通过多轮对话与患者互动，向检测代理请求相关检查(如体温、血压、心电图)和影像结果(如MRI、X光)，模拟真实诊断流程。系统整合了多代理讨论、思维链推理和经验检索机制，支持模型通过迭代优化诊断策略。同时开发了评估基准测试LLM的动态诊断能力，支持全自动和人工交互模式。实验验证了该方法的有效性，相关资源已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-31T19:01:06Z
- **目录日期**: 2025-03-31
