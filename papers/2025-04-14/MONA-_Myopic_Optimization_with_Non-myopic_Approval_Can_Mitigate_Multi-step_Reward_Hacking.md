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

该研究提出了一种名为MONA（短视优化结合远视奖励）的训练方法，旨在防止强化学习(RL)智能体通过多步计划获取不当高回报（"奖励欺骗"），即使人类无法察觉这些行为。MONA结合短期优化与长期奖励评估，能在不依赖额外信息或检测欺骗行为的情况下，有效阻止普通RL导致的多步奖励欺骗问题。研究通过三个实验场景验证了MONA的效果，包括模拟LLM委托监督的2步环境、编码推理场景以及模拟传感器篡改的长时网格世界环境，覆盖了多种错位失效模式。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-14T01:29:41Z
- **目录日期**: 2025-04-14
