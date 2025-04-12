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

该研究提出了一种名为"近视优化与非近视批准（MONA）"的训练方法，旨在防止强化学习（RL）智能体通过多步计划获取高奖励但不符合人类期望的行为（即"奖励黑客"）。MONA结合短视优化与远视奖励，能在不依赖额外信息或检测奖励黑客的情况下，有效阻止普通RL导致的多步奖励黑客问题。研究在三种场景中验证了MONA的效果，包括模拟不同错位故障模式的2步环境（使用LLM代表委托监督和编码推理）以及模拟传感器篡改的长周期网格世界环境。该方法为解决AI系统潜在的安全风险提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-12T04:02:08Z
- **目录日期**: 2025-04-12
