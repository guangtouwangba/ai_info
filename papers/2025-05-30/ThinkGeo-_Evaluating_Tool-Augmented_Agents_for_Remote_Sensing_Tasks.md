# ThinkGeo: Evaluating Tool-Augmented Agents for Remote Sensing Tasks

**URL**: http://arxiv.org/abs/2505.23752v1

## 原始摘要

Recent progress in large language models (LLMs) has enabled tool-augmented
agents capable of solving complex real-world tasks through step-by-step
reasoning. However, existing evaluations often focus on general-purpose or
multimodal scenarios, leaving a gap in domain-specific benchmarks that assess
tool-use capabilities in complex remote sensing use cases. We present ThinkGeo,
an agentic benchmark designed to evaluate LLM-driven agents on remote sensing
tasks via structured tool use and multi-step planning. Inspired by
tool-interaction paradigms, ThinkGeo includes human-curated queries spanning a
wide range of real-world applications such as urban planning, disaster
assessment and change analysis, environmental monitoring, transportation
analysis, aviation monitoring, recreational infrastructure, and industrial site
analysis. Each query is grounded in satellite or aerial imagery and requires
agents to reason through a diverse toolset. We implement a ReAct-style
interaction loop and evaluate both open and closed-source LLMs (e.g., GPT-4o,
Qwen2.5) on 436 structured agentic tasks. The benchmark reports both step-wise
execution metrics and final answer correctness. Our analysis reveals notable
disparities in tool accuracy and planning consistency across models. ThinkGeo
provides the first extensive testbed for evaluating how tool-enabled LLMs
handle spatial reasoning in remote sensing. Our code and dataset are publicly
available


## AI 摘要

ThinkGeo是一个针对遥感领域设计的智能体基准测试，用于评估大语言模型（LLM）在复杂遥感任务中的工具使用和多步规划能力。该基准包含436个人工设计的真实场景任务，涵盖城市规划、灾害评估、环境监测等应用领域，要求智能体通过卫星/航空影像进行空间推理。研究采用ReAct式交互框架测试了开源和闭源模型（如GPT-4o、Qwen2.5），发现不同模型在工具准确性和规划一致性上存在显著差异。ThinkGeo首次为遥感领域的空间推理工具链评估提供了系统测试平台，相关代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T04:02:03Z
- **目录日期**: 2025-05-30
