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

研究表明，大语言模型(LLM)作为智能体在需要多轮推理的任务中表现优异。拒绝采样微调(RFT)是提升LLM智能体能力的有效方法，但受限于专家模型(如GPT-4)在复杂子任务上的失败率。新提出的探索专家失败(EEF)方法创新性地从失败的专家轨迹中提取有价值的行动指导(如关键计划和动作)，同时剔除有害动作以避免干扰学习过程。实验证明，EEF成功解决了部分RFT无法处理的子任务，在WebShop任务中达到62%的胜率(优于RFT的53.6%和GPT-4的35.6%)，并在WebShop和SciWorld基准测试中创下新纪录。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-18T23:01:18Z
- **目录日期**: 2025-04-18
