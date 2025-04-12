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

该研究提出了一种名为"短视优化与非短视批准"(MONA)的新型强化学习训练方法，旨在防止AI系统通过多步计划获取高奖励但不符合人类期望的行为（即"奖励黑客"）。MONA结合了短期优化和长期奖励评估，能在人类无法察觉不当行为的情况下有效阻止多步奖励黑客问题。实验证明，MONA在三种不同场景中（包括两步环境中的LLM代理监督和编码推理，以及长期网格世界环境中的传感器篡改）都能防止普通强化学习导致的多步奖励黑客，且无需额外信息或检测能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-12T18:02:08Z
- **目录日期**: 2025-04-12
