# Agent-X: Evaluating Deep Multimodal Reasoning in Vision-Centric Agentic Tasks

**URL**: http://arxiv.org/abs/2505.24876v1

## 原始摘要

Deep reasoning is fundamental for solving complex tasks, especially in
vision-centric scenarios that demand sequential, multimodal understanding.
However, existing benchmarks typically evaluate agents with fully synthetic,
single-turn queries, limited visual modalities, and lack a framework to assess
reasoning quality over multiple steps as required in real-world settings. To
address this, we introduce Agent-X, a large-scale benchmark for evaluating
vision-centric agents multi-step and deep reasoning capabilities in real-world,
multimodal settings. Agent- X features 828 agentic tasks with authentic visual
contexts, including images, multi-image comparisons, videos, and instructional
text. These tasks span six major agentic environments: general visual
reasoning, web browsing, security and surveillance, autonomous driving, sports,
and math reasoning. Our benchmark requires agents to integrate tool use with
explicit, stepwise decision-making in these diverse settings. In addition, we
propose a fine-grained, step-level evaluation framework that assesses the
correctness and logical coherence of each reasoning step and the effectiveness
of tool usage throughout the task. Our results reveal that even the
best-performing models, including GPT, Gemini, and Qwen families, struggle to
solve multi-step vision tasks, achieving less than 50% full-chain success.
These findings highlight key bottlenecks in current LMM reasoning and tool-use
capabilities and identify future research directions in vision-centric agentic
reasoning models. Our data and code are publicly available at
https://github.com/mbzuai-oryx/Agent-X


## AI 摘要

研究者提出了Agent-X基准测试，用于评估视觉中心智能体在真实多模态环境中的多步骤深度推理能力。该基准包含828个任务，涵盖图像、视频、文本等多种视觉模态，涉及六大应用场景。测试要求智能体结合工具使用进行显式分步决策，并采用细粒度评估框架分析每一步的逻辑连贯性和工具使用有效性。结果显示，当前最优模型（如GPT、Gemini等）在多步骤视觉任务中的成功率不足50%，暴露出大型多模态模型在推理和工具使用能力上的瓶颈。相关数据和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T18:01:16Z
- **目录日期**: 2025-06-02
