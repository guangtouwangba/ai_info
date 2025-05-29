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

本文提出了3DMem-Bench基准测试和3DLLM-Mem模型，旨在解决大语言模型(LLMs)在3D环境中空间-时间记忆推理的不足。3DMem-Bench包含26,000多条轨迹和2,892项任务，用于评估代理在3D环境中的长期记忆能力。3DLLM-Mem采用动态记忆管理机制，通过工作记忆标记查询并融合情景记忆中的关键空间-时间特征，使代理能专注于任务相关信息，同时保持复杂长期环境中的记忆效率。实验表明，该模型在最具挑战性的野外任务中成功率比基线高出16.5%，达到最先进水平。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T02:30:23Z
- **目录日期**: 2025-05-29
