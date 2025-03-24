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

本文提出LLM+MAP框架，用于解决双手机器人操作中的长期任务规划难题。该框架结合大型语言模型(LLMs)的上下文学习能力和多智能体规划，克服了LLMs在复杂任务中的推理错误和幻觉问题。通过GPT-4o后端实现，实验表明LLM+MAP在规划时间、成功率等指标上优于直接使用LLMs(包括GPT-4o/V3/o1/R1)的方案，有效提升了双手机器人任务分解和分配的效率。研究为机器人推理提供了新见解，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T19:01:34Z
- **目录日期**: 2025-03-24
