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

MR. Video提出了一种基于MapReduce原理的长视频理解框架：先独立密集感知短视频片段（Map），再联合聚合所有片段信息（Reduce）。相比序列模型，它能突破上下文长度限制；相比依赖关键片段选择的视频代理，其并行感知更简单高效。该框架通过两阶段实现：1）字幕生成阶段对片段生成描述并统一命名；2）分析阶段针对用户问题提取相关信息并整合答案。在LVBench测试中，MR. Video比现有最优方法准确率提升超过10%，验证了MapReduce在视频理解中的有效性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T14:01:05Z
- **目录日期**: 2025-04-23
