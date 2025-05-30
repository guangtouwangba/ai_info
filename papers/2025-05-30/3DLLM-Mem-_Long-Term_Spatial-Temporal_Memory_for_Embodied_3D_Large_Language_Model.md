# 3DLLM-Mem: Long-Term Spatial-Temporal Memory for Embodied 3D Large Language Model

**URL**: http://arxiv.org/abs/2505.22657v1

## 原始摘要

Humans excel at performing complex tasks by leveraging long-term memory
across temporal and spatial experiences. In contrast, current Large Language
Models (LLMs) struggle to effectively plan and act in dynamic, multi-room 3D
environments. We posit that part of this limitation is due to the lack of
proper 3D spatial-temporal memory modeling in LLMs. To address this, we first
introduce 3DMem-Bench, a comprehensive benchmark comprising over 26,000
trajectories and 2,892 embodied tasks, question-answering and captioning,
designed to evaluate an agent's ability to reason over long-term memory in 3D
environments. Second, we propose 3DLLM-Mem, a novel dynamic memory management
and fusion model for embodied spatial-temporal reasoning and actions in LLMs.
Our model uses working memory tokens, which represents current observations, as
queries to selectively attend to and fuse the most useful spatial and temporal
features from episodic memory, which stores past observations and interactions.
Our approach allows the agent to focus on task-relevant information while
maintaining memory efficiency in complex, long-horizon environments.
Experimental results demonstrate that 3DLLM-Mem achieves state-of-the-art
performance across various tasks, outperforming the strongest baselines by
16.5% in success rate on 3DMem-Bench's most challenging in-the-wild embodied
tasks.


## AI 摘要

该研究针对大型语言模型(LLM)在动态3D环境中空间-时间记忆建模的不足，提出了3DMem-Bench基准测试(含26,000+轨迹和2,892项任务)和3DLLM-Mem记忆管理模型。该模型通过工作记忆标记查询并融合情景记忆中的关键时空特征，使智能体能在复杂环境中高效聚焦任务相关信息。实验显示，3DLLM-Mem在最具挑战性的野外任务上以16.5%的优势超越现有最佳基线，实现了最先进的性能表现。研究为提升LLM在3D环境中的长期记忆推理和行动能力提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T01:28:26Z
- **目录日期**: 2025-05-30
