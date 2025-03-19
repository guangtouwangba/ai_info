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

AI助手需要了解自身的能力和限制，包括何时使用参数知识或工具、何时信任工具输出、何时回避或保留意见。这些能力难以通过监督微调教授，因为需要反映代理具体能力的示例。为此，提出了一种新方法——协作自玩，通过多代理协作，集体获得正确答案来奖励团队。这种互动结构中的激励促使元知识的产生。实验表明，多代理社区的群体奖励可以诱导策略，这些策略在单个代理独立部署时能改善工具使用和选择性预测。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T14:01:59Z
- **目录日期**: 2025-03-19
