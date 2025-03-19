# VideoMind: A Chain-of-LoRA Agent for Long Video Reasoning

**URL**: http://arxiv.org/abs/2503.13444v1

## 原始摘要

Videos, with their unique temporal dimension, demand precise grounded
understanding, where answers are directly linked to visual, interpretable
evidence. Despite significant breakthroughs in reasoning capabilities within
Large Language Models, multi-modal reasoning - especially for videos - remains
unexplored. In this work, we introduce VideoMind, a novel video-language agent
designed for temporal-grounded video understanding. VideoMind incorporates two
key innovations: (i) We identify essential capabilities for video temporal
reasoning and develop a role-based agentic workflow, including a planner for
coordinating different roles, a grounder for temporal localization, a verifier
to assess temporal interval accuracy, and an answerer for question-answering.
(ii) To efficiently integrate these diverse roles, we propose a novel
Chain-of-LoRA strategy, enabling seamless role-switching via lightweight LoRA
adaptors while avoiding the overhead of multiple models, thus balancing
efficiency and flexibility. Extensive experiments on 14 public benchmarks
demonstrate that our agent achieves state-of-the-art performance on diverse
video understanding tasks, including 3 on grounded video question-answering, 6
on video temporal grounding, and 5 on general video question-answering,
underscoring its effectiveness in advancing video agent and long-form temporal
reasoning.


## AI 摘要

这是一篇关于AI Agent领域的论文，主要探讨了最新进展和研究方向。论文提出了一种新的方法来提高Agent的性能，并通过实验证明了其有效性。这项研究对未来的AI Agent开发具有重要指导意义。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T09:16:53+08:00
- **目录日期**: 2025-03-19
