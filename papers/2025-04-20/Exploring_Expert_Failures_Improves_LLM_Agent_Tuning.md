# Exploring Expert Failures Improves LLM Agent Tuning

**URL**: http://arxiv.org/abs/2504.13145v1

## 原始摘要

Large Language Models (LLMs) have shown tremendous potential as agents,
excelling at tasks that require multiple rounds of reasoning and interactions.
Rejection Sampling Fine-Tuning (RFT) has emerged as an effective method for
finetuning LLMs as agents: it first imitates expert-generated successful
trajectories and further improves agentic skills through iterative fine-tuning
on successful, self-generated trajectories. However, since the expert (e.g.,
GPT-4) succeeds primarily on simpler subtasks and RFT inherently favors simpler
scenarios, many complex subtasks remain unsolved and persistently
out-of-distribution (OOD). Upon investigating these challenging subtasks, we
discovered that previously failed expert trajectories can often provide
valuable guidance, e.g., plans and key actions, that can significantly improve
agent exploration efficiency and acquisition of critical skills. Motivated by
these observations, we propose Exploring Expert Failures (EEF), which
identifies beneficial actions from failed expert trajectories and integrates
them into the training dataset. Potentially harmful actions are meticulously
excluded to prevent contamination of the model learning process. By leveraging
the beneficial actions in expert failures, EEF successfully solves some
previously unsolvable subtasks and improves agent tuning performance.
Remarkably, our approach achieved a 62\% win rate in WebShop, outperforming RFT
(53. 6\%) and GPT-4 (35. 6\%), and to the best of our knowledge, setting a new
state-of-the-art as the first method to surpass a score of 0.81 in WebShop and
exceed 81 in SciWorld.


## AI 摘要

该研究提出了一种名为"探索专家失败"(EEF)的新方法，用于改进大型语言模型(LLMs)作为智能体的微调过程。传统拒绝采样微调(RFT)方法主要依赖专家成功轨迹，但会忽略复杂子任务。EEF创新性地从专家失败轨迹中提取有价值信息(如关键行动和计划)，同时排除有害内容，显著提高了智能体的探索效率和关键技能获取。实验表明，EEF在WebShop任务中达到62%的胜率，超越RFT(53.6%)和GPT-4(35.6%)，并在WebShop和SciWorld任务中创下新的最优性能记录。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-20T19:01:15Z
- **目录日期**: 2025-04-20
