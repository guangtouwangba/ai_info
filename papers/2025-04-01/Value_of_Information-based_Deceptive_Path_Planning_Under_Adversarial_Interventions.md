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

本文提出了一种新的马尔可夫决策过程(MDP)模型，用于解决对抗性干预下的欺骗路径规划(DPP)问题。研究者设计了基于信息价值(VoI)的目标函数，使路径规划代理能通过选择信息价值低的轨迹来欺骗对抗性观察者，诱使其采取次优干预措施。通过结合MDP的线性规划理论，开发了计算高效的解决方案。实验表明，该方法在对抗性干预下能有效实现欺骗性，且在网格世界示例中表现优于现有DPP方法和保守路径规划方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T08:02:24Z
- **目录日期**: 2025-04-01
