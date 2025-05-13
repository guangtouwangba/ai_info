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

本文提出了一个基于状态特征影响的强化学习解释理论框架，旨在解决智能体决策缺乏透明度的问题。该框架聚焦三个核心交互要素：行为（智能体动作）、性能（智能体成果）和价值估计（智能体预期），将状态特征视为合作产生这些要素的"参与者"，并运用合作博弈论中的Shapley值量化各特征影响力。这种方法提供了一类具有明确语义和理论保证的数学解释，能产生符合人类直觉的见解，统一并扩展了现有工作。该框架为建立更可解释、可信赖的强化学习系统奠定了理论基础，特别适用于安全关键领域。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T04:02:04Z
- **目录日期**: 2025-05-13
