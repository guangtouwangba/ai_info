# LLMs are Greedy Agents: Effects of RL Fine-tuning on Decision-Making Abilities

**URL**: http://arxiv.org/abs/2504.16078v1

## 原始摘要

The success of Large Language Models (LLMs) has sparked interest in various
agentic applications. A key hypothesis is that LLMs, leveraging common sense
and Chain-of-Thought (CoT) reasoning, can effectively explore and efficiently
solve complex domains. However, LLM agents have been found to suffer from
sub-optimal exploration and the knowing-doing gap, the inability to effectively
act on knowledge present in the model. In this work, we systematically study
why LLMs perform sub-optimally in decision-making scenarios. In particular, we
closely examine three prevalent failure modes: greediness, frequency bias, and
the knowing-doing gap. We propose mitigation of these shortcomings by
fine-tuning via Reinforcement Learning (RL) on self-generated CoT rationales.
Our experiments across multi-armed bandits, contextual bandits, and
Tic-tac-toe, demonstrate that RL fine-tuning enhances the decision-making
abilities of LLMs by increasing exploration and narrowing the knowing-doing
gap. Finally, we study both classic exploration mechanisms, such as
$\epsilon$-greedy, and LLM-specific approaches, such as self-correction and
self-consistency, to enable more effective fine-tuning of LLMs for
decision-making.


## AI 摘要

大型语言模型（LLMs）在复杂决策场景中存在探索不足、贪婪行为、频率偏差和知行差距等问题。研究表明，通过强化学习（RL）对LLMs生成的思维链（CoT）进行微调，可显著提升其决策能力，增强探索效率并缩小知行差距。实验涵盖多臂老虎机、上下文老虎机和井字棋等任务，验证了RL微调的有效性。此外，研究对比了传统探索机制（如ε-贪婪）与LLM特有方法（如自我纠正和一致性）的效果，为优化LLMs的决策能力提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T14:01:17Z
- **目录日期**: 2025-04-23
