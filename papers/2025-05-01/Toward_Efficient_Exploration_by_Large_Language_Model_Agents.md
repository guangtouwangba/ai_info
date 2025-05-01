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

这篇摘要探讨了基于大语言模型(LLM)的强化学习(RL)智能体的探索效率问题。研究发现，当前许多LLM智能体设计方案难以有效解决探索挑战，而传统RL算法虽能优雅处理探索问题，却难以在纯自然语言环境中实现。作者提出了一种创新方法：不是通过微调或上下文学习让LLM隐式模仿RL算法，而是直接利用LLM显式实现一个已知的数据高效RL算法(后验采样强化学习)。实验结果表明，这种基于LLM的已知算法实现，在需要谨慎探索的自然语言任务中表现显著更优。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T00:01:06Z
- **目录日期**: 2025-05-01
