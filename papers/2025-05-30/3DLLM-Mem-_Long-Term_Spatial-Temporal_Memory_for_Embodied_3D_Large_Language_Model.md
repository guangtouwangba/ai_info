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

人类擅长利用长期记忆处理复杂任务，而当前大语言模型（LLMs）在动态3D环境中表现欠佳。为此，研究者提出3DMem-Bench基准测试（含26,000+轨迹和2,892项任务）来评估3D空间-时间记忆能力，并开发了3DLLM-Mem模型。该模型通过工作记忆标记整合当前观察与过去记忆，选择性融合关键时空特征，在保持高效的同时聚焦任务相关信息。实验显示，3DLLM-Mem在最具挑战性的任务上成功率超越基线16.5%，实现了最先进性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-30T00:01:42Z
- **目录日期**: 2025-05-30
