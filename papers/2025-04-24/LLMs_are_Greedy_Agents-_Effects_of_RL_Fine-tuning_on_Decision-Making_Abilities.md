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

大型语言模型（LLMs）在决策任务中存在探索不足和"知行差距"（知道但无法有效执行）的问题。研究发现其失败模式主要包括贪婪性、频率偏差和知行差距。通过基于强化学习（RL）对模型自生成的思维链（CoT）进行微调，可显著提升LLMs的探索能力和决策表现。实验在多臂老虎机、上下文老虎机和井字棋等任务中验证了该方法的有效性。研究还比较了传统探索机制（如ε-贪婪）和LLM特有方法（如自我校正和一致性）对决策优化的作用。结果表明RL微调能有效缩小知行差距并改善决策质量。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T01:28:39Z
- **目录日期**: 2025-04-24
