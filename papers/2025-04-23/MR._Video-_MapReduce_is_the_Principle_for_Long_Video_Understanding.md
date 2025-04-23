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

MR. Video 是一个基于 MapReduce 原理的长视频理解框架。它通过两个步骤处理长视频：(1) Map 阶段：独立密集感知短视频片段；(2) Reduce 阶段：联合聚合所有片段信息。相比序列到序列的视觉语言模型（VLMs），MR. Video 不受上下文长度限制，能更详细地感知短视频。相比依赖关键片段选择的视频代理，其并行感知更简单且可扩展。该框架采用两级 MapReduce：(A) 字幕生成：为短视频生成字幕并标准化重复内容；(B) 分析：从短视频中提取相关信息并整合为最终答案。在 LVBench 上，MR. Video 比现有方法准确率提升超过 10%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T19:01:02Z
- **目录日期**: 2025-04-23
