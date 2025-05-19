# Signal attenuation enables scalable decentralized multi-agent reinforcement learning over networks

**URL**: http://arxiv.org/abs/2505.11461v1

## 原始摘要

Classic multi-agent reinforcement learning (MARL) methods require that agents
enjoy global state observability, preventing development of decentralized
algorithms and limiting scalability. Recent work has shown that, under
assumptions on decaying inter-agent influence, global observability can be
replaced by local neighborhood observability at each agent, enabling
decentralization and scalability. Real-world applications enjoying such decay
properties remain underexplored, however, despite the fact that signal power
decay, or signal attenuation, due to path loss is an intrinsic feature of many
problems in wireless communications and radar networks. In this paper, we show
that signal attenuation enables decentralization in MARL by considering the
illustrative special case of performing power allocation for target detection
in a radar network. To achieve this, we propose two new constrained multi-agent
Markov decision process formulations of this power allocation problem, derive
local neighborhood approximations for global value function and gradient
estimates and establish corresponding error bounds, and develop decentralized
saddle point policy gradient algorithms for solving the proposed problems. Our
approach, though oriented towards the specific radar network problem we
consider, provides a useful model for future extensions to additional problems
in wireless communications and radar networks.


## AI 摘要

传统多智能体强化学习（MARL）方法依赖全局状态可观测性，限制了算法的去中心化和可扩展性。研究表明，在智能体间影响衰减的假设下，局部邻域观测可替代全局观测。本文聚焦无线通信和雷达网络中的信号衰减特性，以雷达网络目标检测的功率分配问题为例，提出两种约束多智能体马尔可夫决策过程模型，推导了全局价值函数和梯度估计的局部邻域近似及其误差界，并开发了去中心化鞍点策略梯度算法。该方法为无线通信和雷达网络中的其他问题提供了可扩展的解决框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T14:01:50Z
- **目录日期**: 2025-05-19
