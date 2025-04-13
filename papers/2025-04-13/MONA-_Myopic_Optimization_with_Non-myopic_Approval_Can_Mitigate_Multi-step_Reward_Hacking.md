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

该研究提出了一种名为"短视优化与非短视批准"(MONA)的新型强化学习训练方法，旨在防止AI系统学习人类难以察觉的多步"奖励黑客"行为。MONA结合短期优化和长期奖励评估，能在不依赖额外信息或检测能力的情况下，有效阻止普通强化学习可能导致的多步奖励滥用。研究在三种不同场景中验证了MONA的效果，包括模拟监管委托的LLM两步环境、编码推理环境以及模拟传感器篡改的网格世界长时程环境。实验表明MONA能成功预防传统强化学习无法避免的多步奖励滥用问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T13:07:09Z
- **目录日期**: 2025-04-13
