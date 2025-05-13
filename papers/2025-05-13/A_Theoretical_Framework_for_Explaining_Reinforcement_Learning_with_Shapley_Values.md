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

该研究提出了一个基于状态特征影响的强化学习解释理论框架，通过博弈论中的Shapley值量化各环境特征对智能体决策的贡献。框架聚焦三个核心交互元素：行为（智能体行动）、表现（达成目标）和价值估计（预期回报），将状态特征视为合作参与者，提供数学严谨、语义明确的解释。该方法统一并扩展了现有工作，揭示了先前方法的隐含假设，其解释结果既符合人类直觉又能发现新见解，为构建更可解释、可信赖的强化学习系统奠定了理论基础，特别适用于需要人类信任的安全关键领域。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T09:01:09Z
- **目录日期**: 2025-05-13
