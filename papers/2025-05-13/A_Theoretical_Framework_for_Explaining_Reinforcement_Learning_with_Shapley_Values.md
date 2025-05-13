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

本文提出了一种基于状态特征影响力的强化学习解释性理论框架，通过分析智能体与环境交互中的三个核心要素（行为、性能和价值估计），利用合作博弈论中的Shapley值量化各状态特征的贡献。该方法提供了数学严谨、语义清晰的解释机制，能揭示传统方法难以发现的洞察，并与人类直觉一致。该框架统一并扩展了现有研究，明确了各种解释方法背后的假设，为构建更可解释、可信赖的强化学习系统奠定了理论基础，尤其适用于需要人类信任与问责制的安全关键领域。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T18:01:15Z
- **目录日期**: 2025-05-13
