# LLM+MAP: Bimanual Robot Task Planning using Large Language Models and Planning Domain Definition Language

**URL**: http://arxiv.org/abs/2503.17309v1

## 原始摘要

Bimanual robotic manipulation provides significant versatility, but also
presents an inherent challenge due to the complexity involved in the spatial
and temporal coordination between two hands. Existing works predominantly focus
on attaining human-level manipulation skills for robotic hands, yet little
attention has been paid to task planning on long-horizon timescales. With their
outstanding in-context learning and zero-shot generation abilities, Large
Language Models (LLMs) have been applied and grounded in diverse robotic
embodiments to facilitate task planning. However, LLMs still suffer from errors
in long-horizon reasoning and from hallucinations in complex robotic tasks,
lacking a guarantee of logical correctness when generating the plan. Previous
works, such as LLM+P, extended LLMs with symbolic planners. However, none have
been successfully applied to bimanual robots. New challenges inevitably arise
in bimanual manipulation, necessitating not only effective task decomposition
but also efficient task allocation. To address these challenges, this paper
introduces LLM+MAP, a bimanual planning framework that integrates LLM reasoning
and multi-agent planning, automating effective and efficient bimanual task
planning. We conduct simulated experiments on various long-horizon manipulation
tasks of differing complexity. Our method is built using GPT-4o as the backend,
and we compare its performance against plans generated directly by LLMs,
including GPT-4o, V3 and also recent strong reasoning models o1 and R1. By
analyzing metrics such as planning time, success rate, group debits, and
planning-step reduction rate, we demonstrate the superior performance of
LLM+MAP, while also providing insights into robotic reasoning. Code is
available at https://github.com/Kchu/LLM-MAP.


## AI 摘要

本文提出LLM+MAP框架，结合大语言模型(LLM)推理与多智能体规划，解决双手机器人操作中的长期任务规划难题。现有方法多关注单手机器人技能，而双手机器人需要复杂的时空协调和任务分配。传统LLM在长期推理中存在错误和幻觉问题，LLM+MAP通过整合符号规划器提升逻辑正确性。实验使用GPT-4o作为后端，在多种复杂任务中对比GPT-4o、V3等模型，结果显示LLM+MAP在规划时间、成功率等指标上表现更优。该框架为双手机器人任务规划提供了新思路，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T21:01:32Z
- **目录日期**: 2025-03-24
