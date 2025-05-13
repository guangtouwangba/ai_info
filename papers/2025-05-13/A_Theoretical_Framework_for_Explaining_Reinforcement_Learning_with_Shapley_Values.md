# A Theoretical Framework for Explaining Reinforcement Learning with Shapley Values

**URL**: http://arxiv.org/abs/2505.07797v1

## 原始摘要

Reinforcement learning agents can achieve superhuman performance, but their
decisions are often difficult to interpret. This lack of transparency limits
deployment, especially in safety-critical settings where human trust and
accountability are essential. In this work, we develop a theoretical framework
for explaining reinforcement learning through the influence of state features,
which represent what the agent observes in its environment. We identify three
core elements of the agent-environment interaction that benefit from
explanation: behaviour (what the agent does), performance (what the agent
achieves), and value estimation (what the agent expects to achieve). We treat
state features as players cooperating to produce each element and apply Shapley
values, a principled method from cooperative game theory, to identify the
influence of each feature. This approach yields a family of mathematically
grounded explanations with clear semantics and theoretical guarantees. We use
illustrative examples to show how these explanations align with human intuition
and reveal novel insights. Our framework unifies and extends prior work, making
explicit the assumptions behind existing approaches, and offers a principled
foundation for more interpretable and trustworthy reinforcement learning.


## AI 摘要

这篇论文提出了一个基于状态特征影响力的强化学习解释性理论框架。研究者通过将状态特征视为合作产生智能体行为的"玩家"，并应用博弈论中的Shapley值方法，量化每个特征对智能体行为、表现和价值估计三个核心要素的影响。该框架提供了一组具有明确语义和理论保证的数学解释方法，能够产生符合人类直觉的见解，同时统一并扩展了现有方法。这种基于博弈论的解释方法为增强强化学习的可解释性和可信度奠定了理论基础，特别适用于需要人类信任和问责制的安全关键领域。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T22:01:11Z
- **目录日期**: 2025-05-13
