# Don't lie to your friends: Learning what you know from collaborative self-play

**URL**: http://arxiv.org/abs/2503.14481v1

## 原始摘要

To be helpful assistants, AI agents must be aware of their own capabilities
and limitations. This includes knowing when to answer from parametric knowledge
versus using tools, when to trust tool outputs, and when to abstain or hedge.
Such capabilities are hard to teach through supervised fine-tuning because they
require constructing examples that reflect the agent's specific capabilities.
We therefore propose a radically new approach to teaching agents what they
know: \emph{collaborative self-play}. We construct multi-agent collaborations
in which the group is rewarded for collectively arriving at correct answers.
The desired meta-knowledge emerges from the incentives built into the structure
of the interaction. We focus on small societies of agents that have access to
heterogeneous tools (corpus-specific retrieval), and therefore must collaborate
to maximize their success while minimizing their effort. Experiments show that
group-level rewards for multi-agent communities can induce policies that
\emph{transfer} to improve tool use and selective prediction in settings where
individual agents are deployed in isolation.


## AI 摘要

本文提出了一种名为“协作自玩”的新方法，用于教导AI代理了解自身能力和限制。通过构建多代理协作环境，团队因共同得出正确答案而获得奖励，从而促使代理学会何时依赖自身知识、使用工具或选择放弃。该方法特别适用于拥有不同工具的小型代理群体，通过集体奖励机制，代理在单独部署时也能更好地使用工具和做出选择性预测。实验表明，这种群体奖励机制能有效提升代理的工具使用能力和预测准确性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T00:01:50Z
- **目录日期**: 2025-03-20
