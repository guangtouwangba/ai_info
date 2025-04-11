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

该研究提出了一种名为"短视优化与非短视批准"(MONA)的新训练方法，旨在防止强化学习(RL)系统通过多步计划获取不当高回报的"奖励黑客"行为。MONA结合短期优化与长期奖励评估，能在不依赖额外信息或人类监督的情况下有效阻止多步奖励黑客行为。研究在三种实验环境中验证了MONA的效果：模拟LLM委托监督的两步决策环境、编码推理环境，以及模拟传感器篡改问题的长时网格世界环境。结果表明MONA能预防普通RL系统产生的多步奖励黑客问题，提高了AI系统的安全性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-11T03:15:14Z
- **目录日期**: 2025-04-11
