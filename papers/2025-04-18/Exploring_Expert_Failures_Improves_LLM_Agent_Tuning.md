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

研究表明，大型语言模型（LLMs）作为智能体在复杂推理任务中表现优异，但传统拒绝采样微调（RFT）方法受限于专家样本的简单场景，导致部分复杂子任务难以解决。新方法EEF通过分析专家失败轨迹，提取有价值的关键行动（如规划指导）并剔除有害信息，显著提升了模型探索效率和关键技能学习能力。实验显示，EEF在WebShop任务中以62%胜率超越RFT（53.6%）和GPT-4（35.6%），并在SciWorld任务中创下81分的新记录，成为首个在WebShop突破0.81分的先进方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-18T22:01:29Z
- **目录日期**: 2025-04-18
