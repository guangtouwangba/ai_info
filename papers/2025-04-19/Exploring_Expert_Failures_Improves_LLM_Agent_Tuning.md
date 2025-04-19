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

该研究提出了一种名为"探索专家失败"(EEF)的新方法，用于提升大语言模型(LLM)作为智能体的性能。传统拒绝采样微调(RFT)方法主要依赖专家成功轨迹，但难以解决复杂子任务。研究发现，专家失败轨迹中的有用信息(如关键行动和计划)能显著提升智能体探索效率和关键技能获取。EEF通过精选失败轨迹中的有益行动并排除有害行动来优化训练数据。实验表明，EEF在WebShop任务中取得62%胜率，超越RFT(53.6%)和GPT-4(35.6%)，并在WebShop和SciWorld任务中创造了新的最优性能记录。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-19T20:01:27Z
- **目录日期**: 2025-04-19
