# MR. Video: "MapReduce" is the Principle for Long Video Understanding

**URL**: http://arxiv.org/abs/2504.16082v1

## 原始摘要

We propose MR. Video, an agentic long video understanding framework that
demonstrates the simple yet effective MapReduce principle for processing long
videos: (1) Map: independently and densely perceiving short video clips, and
(2) Reduce: jointly aggregating information from all clips. Compared with
sequence-to-sequence vision-language models (VLMs), MR. Video performs detailed
short video perception without being limited by context length. Compared with
existing video agents that typically rely on sequential key segment selection,
the Map operation enables simpler and more scalable sequence parallel
perception of short video segments. Its Reduce step allows for more
comprehensive context aggregation and reasoning, surpassing explicit key
segment retrieval. This MapReduce principle is applicable to both VLMs and
video agents, and we use LLM agents to validate its effectiveness.
  In practice, MR. Video employs two MapReduce stages: (A) Captioning:
generating captions for short video clips (map), then standardizing repeated
characters and objects into shared names (reduce); (B) Analysis: for each user
question, analyzing relevant information from individual short videos (map),
and integrating them into a final answer (reduce). MR. Video achieves over 10%
accuracy improvement on the challenging LVBench compared to state-of-the-art
VLMs and video agents.
  Code is available at: https://github.com/ziqipang/MR-Video


## AI 摘要

MR. Video 是一个基于MapReduce原理的长视频理解框架，通过"Map"步骤独立感知短视频片段，再通过"Reduce"步骤聚合所有片段信息。相比传统方法，它不受上下文长度限制，能并行处理片段并实现更全面的推理。该框架包含两个阶段：(A) 首先生成片段字幕并统一命名；(B) 然后分析用户问题并整合信息生成最终答案。在LVBench测试中，MR. Video比现有最优方法准确率提升超过10%。这一原理适用于视觉语言模型和视频代理，研究通过LLM代理验证了其有效性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T15:00:59Z
- **目录日期**: 2025-04-23
