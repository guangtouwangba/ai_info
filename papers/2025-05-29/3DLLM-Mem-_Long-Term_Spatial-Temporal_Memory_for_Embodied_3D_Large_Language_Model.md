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

本文提出了3DLLM-Mem模型，通过动态记忆管理解决大语言模型在3D环境中的时空推理问题。研究首先构建了包含26,000+轨迹和2,892项任务的3DMem-Bench基准测试，评估智能体在3D环境中的长期记忆能力。模型采用工作记忆token表示当前观察，通过选择性注意力机制融合来自存储过去观察的片段记忆中的时空特征，使智能体能专注于任务相关信息。实验表明，3DLLM-Mem在最具挑战性的野外任务中成功率比最强基线高16.5%，实现了最先进的性能表现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T03:20:05Z
- **目录日期**: 2025-05-29
