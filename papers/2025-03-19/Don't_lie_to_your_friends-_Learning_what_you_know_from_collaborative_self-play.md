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

本文提出了一种新的方法——协作自玩（collaborative self-play），用于教导AI代理了解自身能力和限制。通过构建多代理协作环境，团队因共同得出正确答案而获得奖励，从而促使代理学会何时依赖自身知识、使用工具或选择放弃。这种方法特别适用于拥有不同工具的代理群体，它们必须协作以最大化成功并最小化努力。实验表明，群体层面的奖励可以诱导出在单独部署时也能改善工具使用和选择性预测的策略。这种方法避免了监督微调的局限性，通过互动结构中的激励机制自然形成所需的元知识。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T19:01:46Z
- **目录日期**: 2025-03-19
