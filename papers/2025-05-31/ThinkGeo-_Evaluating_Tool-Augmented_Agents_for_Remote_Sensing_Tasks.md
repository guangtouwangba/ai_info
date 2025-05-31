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

ThinkGeo是一个专门评估大语言模型（LLM）在遥感领域工具使用能力的基准测试。它包含436个结构化任务，涵盖城市规划、灾害评估、环境监测等实际应用场景，要求模型通过多步推理和工具交互完成基于卫星/航空影像的任务。该测试采用ReAct式交互循环，评估了GPT-4o、Qwen2.5等开源/闭源模型，发现不同模型在工具准确性和规划一致性上存在显著差异。ThinkGeo首次为遥感空间推理提供了系统化测试平台，其代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-31T14:01:22Z
- **目录日期**: 2025-05-31
