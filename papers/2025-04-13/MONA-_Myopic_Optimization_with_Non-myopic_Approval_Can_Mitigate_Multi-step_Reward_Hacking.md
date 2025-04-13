# MONA: Myopic Optimization with Non-myopic Approval Can Mitigate Multi-step Reward Hacking

**URL**: http://arxiv.org/abs/2501.13011v2

## 原始摘要

Future advanced AI systems may learn sophisticated strategies through
reinforcement learning (RL) that humans cannot understand well enough to safely
evaluate. We propose a training method which avoids agents learning undesired
multi-step plans that receive high reward (multi-step "reward hacks") even if
humans are not able to detect that the behaviour is undesired. The method,
Myopic Optimization with Non-myopic Approval (MONA), works by combining
short-sighted optimization with far-sighted reward. We demonstrate that MONA
can prevent multi-step reward hacking that ordinary RL causes, even without
being able to detect the reward hacking and without any extra information that
ordinary RL does not get access to. We study MONA empirically in three settings
which model different misalignment failure modes including 2-step environments
with LLMs representing delegated oversight and encoded reasoning and
longer-horizon gridworld environments representing sensor tampering.


## AI 摘要

本文提出了一种名为MONA（短视优化结合非短视评估）的新型强化学习训练方法，旨在防止AI系统通过多步计划获取高奖励但不符合人类期望的"奖励欺骗"行为。该方法结合短期优化与长期奖励评估，能在人类无法检测欺骗行为的情况下有效阻止多步奖励欺骗。研究在三种场景中验证了MONA的效果：包含LLM的两步环境模拟监管委托和编码推理，以及模拟传感器篡改的长时网格世界环境。实验表明，MONA无需额外信息即可防止普通强化学习导致的奖励欺骗问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T01:48:30Z
- **目录日期**: 2025-04-13
