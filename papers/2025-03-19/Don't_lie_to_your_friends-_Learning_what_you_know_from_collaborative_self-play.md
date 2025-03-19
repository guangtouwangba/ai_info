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

AI助手需要了解自身能力和局限，以决定何时依赖内部知识、使用工具或选择回避。传统监督学习方法难以教授这种能力，因此提出了一种新方法——协作自玩。通过构建多智能体协作环境，团队因共同得出正确答案而获得奖励，从而促使智能体学会使用工具和选择性预测。实验表明，这种群体奖励机制能有效提升单个智能体在独立部署时的工具使用和预测能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T11:49:07Z
- **目录日期**: 2025-03-19
