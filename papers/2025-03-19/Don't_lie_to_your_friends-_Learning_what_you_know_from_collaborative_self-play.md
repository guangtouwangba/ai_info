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

AI助手需要了解自身能力和局限，以决定何时依赖内部知识、使用工具或选择不回答。传统监督微调难以教授这些能力，因此提出“协作自玩”方法，通过多智能体协作，集体获得正确答案来激励元知识的形成。实验表明，群体奖励能促使智能体在单独部署时更有效地使用工具和做出选择性预测。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T12:02:05Z
- **目录日期**: 2025-03-19
