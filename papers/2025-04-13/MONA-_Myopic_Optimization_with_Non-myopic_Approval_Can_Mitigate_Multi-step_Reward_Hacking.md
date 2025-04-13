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

该研究提出了一种名为"近视优化与非近视批准"(MONA)的训练方法，旨在防止强化学习(RL)智能体通过多步计划获取不当高回报(即"回报黑客"行为)。MONA结合短视优化和远视奖励，能在人类无法检测到不当行为的情况下阻止多步回报黑客行为。实验证明，MONA在三种模拟不同错位故障模式的场景中有效，包括使用大语言模型模拟委托监督和编码推理的两步环境，以及模拟传感器篡改的长时网格世界环境。该方法无需额外信息即可防止普通RL导致的多步回报黑客问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T10:02:25Z
- **目录日期**: 2025-04-13
