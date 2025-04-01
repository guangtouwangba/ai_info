# Value of Information-based Deceptive Path Planning Under Adversarial Interventions

**URL**: http://arxiv.org/abs/2503.24284v1

## 原始摘要

Existing methods for deceptive path planning (DPP) address the problem of
designing paths that conceal their true goal from a passive, external observer.
Such methods do not apply to problems where the observer has the ability to
perform adversarial interventions to impede the path planning agent. In this
paper, we propose a novel Markov decision process (MDP)-based model for the DPP
problem under adversarial interventions and develop new value of information
(VoI) objectives to guide the design of DPP policies. Using the VoI objectives
we propose, path planning agents deceive the adversarial observer into choosing
suboptimal interventions by selecting trajectories that are of low
informational value to the observer. Leveraging connections to the linear
programming theory for MDPs, we derive computationally efficient solution
methods for synthesizing policies for performing DPP under adversarial
interventions. In our experiments, we illustrate the effectiveness of the
proposed solution method in achieving deceptiveness under adversarial
interventions and demonstrate the superior performance of our approach to both
existing DPP methods and conservative path planning approaches on illustrative
gridworld problems.


## AI 摘要

本文提出了一种新的基于马尔可夫决策过程(MDP)的欺骗性路径规划(DPP)模型，用于解决对抗性干预下的路径规划问题。研究者开发了新的信息价值(VoI)目标来指导DPP策略设计，使规划者能够通过选择低信息价值的轨迹来欺骗对手采取次优干预措施。通过结合MDP的线性规划理论，该方法能高效计算对抗干预下的欺骗性策略。实验表明，该方法在网格世界问题中比现有DPP方法和保守路径规划方法表现更优，能有效实现对抗干预下的欺骗效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T16:03:11Z
- **目录日期**: 2025-04-01
