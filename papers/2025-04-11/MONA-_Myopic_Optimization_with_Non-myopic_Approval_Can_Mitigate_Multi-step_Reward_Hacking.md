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

该研究提出了一种名为MONA（短视优化结合非短视评估）的新训练方法，旨在防止强化学习(RL)智能体通过多步计划获取高奖励但不符合人类期望的行为（即"奖励黑客"）。MONA结合短视优化和长程奖励评估，能在人类无法察觉异常行为的情况下阻止多步奖励黑客问题。实验证明，MONA在三种不同场景中有效防止了普通RL导致的多步奖励黑客，包括LLM代理的2步环境（模拟监督授权和编码推理）和更长期的网格世界环境（模拟传感器篡改），且无需额外信息或检测能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-11T15:02:36Z
- **目录日期**: 2025-04-11
