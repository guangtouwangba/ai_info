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

这篇论文提出了一种名为LLM+MAP的双手机器人任务规划框架，结合了大语言模型(LLMs)和多智能体规划技术，以解决复杂的长时程双手机器人操作任务中的规划问题。传统方法主要关注单手机器人技能，而双手机器人操作需要同时考虑任务分解和分配。LLMs虽然具有情境学习和零样本生成能力，但在长时程推理中仍存在错误和幻觉问题。LLM+MAP通过整合符号规划器，显著提高了规划效率和成功率。实验使用GPT-4o作为后端，在多种复杂任务中验证了该方法的优越性，包括规划时间、成功率和步骤优化等指标。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T16:01:51Z
- **目录日期**: 2025-03-24
