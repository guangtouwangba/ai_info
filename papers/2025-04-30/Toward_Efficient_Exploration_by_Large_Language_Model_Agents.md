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

这篇论文探讨了如何利用大语言模型(LLMs)实现高效探索的强化学习(RL)算法。研究表明，当前许多LLM智能体设计在数据效率方面存在探索不足的问题。作者提出了一种创新方法：不是通过微调或上下文学习让LLMs隐式模仿RL算法，而是直接基于LLMs显式实现经典的"后验采样强化学习"算法。该方法在需要谨慎探索的自然语言任务中展现出显著优势，相比现有方案能更有效地实现数据高效的强化学习。实验结果表明，这种明确实现已知高效RL算法的方式比隐式方法更具优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T21:00:54Z
- **目录日期**: 2025-04-30
