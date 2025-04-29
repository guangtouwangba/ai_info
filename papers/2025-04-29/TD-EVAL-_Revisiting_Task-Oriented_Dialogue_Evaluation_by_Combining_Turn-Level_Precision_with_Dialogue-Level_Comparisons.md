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

本文提出了TD-EVAL评估框架，用于改进任务导向对话(TOD)系统的评估方法。该框架包含两个层面：回合级评估从对话连贯性、知识一致性和策略合规性三个维度分析每个响应；对话级评估通过TOD Agent Arena进行两两比较。实验表明，TD-EVAL能有效识别传统指标遗漏的对话错误，且比传统和基于大语言模型(LLM)的指标更符合人类判断。这一可插拔的评估框架为TOD系统提供了兼顾细粒度分析和整体质量评估的新范式。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T23:01:51Z
- **目录日期**: 2025-04-29
