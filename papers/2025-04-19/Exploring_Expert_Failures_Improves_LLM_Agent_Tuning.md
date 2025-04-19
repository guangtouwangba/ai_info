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

研究发现，大语言模型（LLM）作为智能体时，通过拒绝采样微调（RFT）方法虽能提升性能，但在复杂子任务上仍表现不佳。专家（如GPT-4）的失败轨迹中往往包含有价值的指导信息（如关键行动和计划）。为此，研究者提出"探索专家失败"（EEF）方法，筛选失败轨迹中的有益行动加入训练数据，同时剔除有害行动以避免干扰模型学习。EEF成功解决了部分RFT无法处理的子任务，在WebShop任务中以62%胜率超越RFT（53.6%）和GPT-4（35.6%），并在WebShop和SciWorld中创下81分以上的新纪录。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-19T17:01:25Z
- **目录日期**: 2025-04-19
