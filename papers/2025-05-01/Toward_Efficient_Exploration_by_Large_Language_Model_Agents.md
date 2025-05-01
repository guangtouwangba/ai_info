# Toward Efficient Exploration by Large Language Model Agents

**URL**: http://arxiv.org/abs/2504.20997v1

## 原始摘要

A burgeoning area within reinforcement learning (RL) is the design of
sequential decision-making agents centered around large language models (LLMs).
While autonomous decision-making agents powered by modern LLMs could facilitate
numerous real-world applications, such successes demand agents that are capable
of data-efficient RL. One key obstacle to achieving data efficiency in RL is
exploration, a challenge that we demonstrate many recent proposals for LLM
agent designs struggle to contend with. Meanwhile, classic algorithms from the
RL literature known to gracefully address exploration require technical
machinery that can be challenging to operationalize in purely natural language
settings. In this work, rather than relying on finetuning or in-context
learning to coax LLMs into implicitly imitating a RL algorithm, we illustrate
how LLMs can be used to explicitly implement an existing RL algorithm
(Posterior Sampling for Reinforcement Learning) whose capacity for
statistically-efficient exploration is already well-studied. We offer empirical
results demonstrating how our LLM-based implementation of a known,
data-efficient RL algorithm can be considerably more effective in natural
language tasks that demand prudent exploration.


## AI 摘要

本文探讨了基于大语言模型（LLM）的强化学习（RL）智能体在数据效率方面的挑战。尽管LLM驱动的自主决策智能体具有广泛应用前景，但现有方案往往难以解决探索效率问题。研究者提出了一种新方法：不依赖微调或上下文学习让LLM隐式模仿RL算法，而是直接利用LLM显式实现经典RL算法（后验采样强化学习），该算法具有已知的统计高效探索能力。实验表明，这种基于LLM的显式实现方法在需要谨慎探索的自然语言任务中显著提升了数据效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T01:28:50Z
- **目录日期**: 2025-05-01
