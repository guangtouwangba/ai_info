# TD-EVAL: Revisiting Task-Oriented Dialogue Evaluation by Combining Turn-Level Precision with Dialogue-Level Comparisons

**URL**: http://arxiv.org/abs/2504.19982v1

## 原始摘要

Task-oriented dialogue (TOD) systems are experiencing a revolution driven by
Large Language Models (LLMs), yet the evaluation methodologies for these
systems remain insufficient for their growing sophistication. While traditional
automatic metrics effectively assessed earlier modular systems, they focus
solely on the dialogue level and cannot detect critical intermediate errors
that can arise during user-agent interactions. In this paper, we introduce
TD-EVAL (Turn and Dialogue-level Evaluation), a two-step evaluation framework
that unifies fine-grained turn-level analysis with holistic dialogue-level
comparisons. At turn level, we evaluate each response along three TOD-specific
dimensions: conversation cohesion, backend knowledge consistency, and policy
compliance. Meanwhile, we design TOD Agent Arena that uses pairwise comparisons
to provide a measure of dialogue-level quality. Through experiments on MultiWOZ
2.4 and {\tau}-Bench, we demonstrate that TD-EVAL effectively identifies the
conversational errors that conventional metrics miss. Furthermore, TD-EVAL
exhibits better alignment with human judgments than traditional and LLM-based
metrics. These findings demonstrate that TD-EVAL introduces a new paradigm for
TOD system evaluation, efficiently assessing both turn and system levels with a
plug-and-play framework for future research.


## AI 摘要

本文提出TD-EVAL框架，用于评估任务导向对话系统(TOD)的性能。该框架包含两个层面：回合级评估（分析对话连贯性、知识一致性和策略合规性）和对话级评估（通过两两比较衡量整体质量）。实验表明，TD-EVAL能有效识别传统指标遗漏的对话错误，且比传统方法和基于大语言模型(LLM)的指标更符合人类判断。该框架为TOD系统评估提供了即插即用的新范式，可同时评估回合和系统层面的表现。研究基于MultiWOZ 2.4和τ-Bench数据集验证了其有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T15:02:12Z
- **目录日期**: 2025-04-29
