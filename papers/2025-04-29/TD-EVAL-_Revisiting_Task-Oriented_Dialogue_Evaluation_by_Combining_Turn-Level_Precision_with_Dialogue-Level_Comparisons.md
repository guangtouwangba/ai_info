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

本文提出TD-EVAL框架，用于评估基于大语言模型的任务导向对话系统(TOD)。传统自动评估仅关注对话层面，无法检测交互过程中的关键中间错误。TD-EVAL采用两步评估：在轮次层面评估对话连贯性、知识一致性和策略合规性；在对话层面通过TOD Agent Arena进行成对比较。实验表明，TD-EVAL能有效识别传统指标遗漏的错误，且比传统和基于LLM的指标更符合人类判断。该框架为TOD系统评估提供了即插即用的新范式，可同时评估轮次和系统层面性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T17:02:03Z
- **目录日期**: 2025-04-29
