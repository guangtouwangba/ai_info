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

该研究提出了一种基于状态特征影响的强化学习解释理论框架，通过博弈论中的Shapley值量化各环境特征对智能体决策的贡献。该框架聚焦三个核心交互要素：行为（智能体行动）、表现（智能体成就）和估值（智能体预期），将状态特征视为合作玩家，提供数学严谨、语义明确的解释。研究统一了现有方法，揭示其隐含假设，并通过案例验证解释结果符合人类直觉且能发现新见解。该理论为增强强化学习的可解释性和可信度奠定了原则性基础，特别适用于需要人类信任的安全关键领域。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T21:00:53Z
- **目录日期**: 2025-05-13
