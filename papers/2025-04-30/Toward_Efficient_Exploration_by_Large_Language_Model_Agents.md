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

本文探讨了基于大语言模型(LLM)的强化学习(RL)智能体在数据效率方面的挑战。研究表明，现有LLM智能体设计方案难以解决探索问题，而传统RL算法虽能有效处理探索但难以用纯自然语言实现。不同于依赖微调或上下文学习让LLM隐式模仿RL算法，作者提出让LLM显式实现已知的高效探索RL算法(后验采样强化学习)。实验证明，这种基于LLM的显式实现方法在需要谨慎探索的自然语言任务中表现显著更优。该研究为开发数据高效的LLM-RL智能体提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T20:00:56Z
- **目录日期**: 2025-04-30
