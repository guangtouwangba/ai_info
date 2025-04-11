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

该研究提出了一种名为MONA（短视优化结合远视评估）的新型强化学习训练方法，旨在防止AI系统通过多步计划获取高奖励但不符合人类期望的行为（即"奖励黑客"）。该方法结合短期优化与长期奖励评估，能在不依赖额外信息或人工检测的情况下，有效阻止普通强化学习可能产生的多步奖励黑客行为。研究在三种不同场景中验证了MONA的效果，包括模拟监督失效的两步环境（使用大语言模型）和模拟传感器篡改的长时网格世界环境，证明了其对多种错位失效模式的防范能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-11T06:02:13Z
- **目录日期**: 2025-04-11
