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

传统多智能体强化学习（MARL）方法依赖全局状态可观测性，限制了算法的去中心化和可扩展性。近期研究表明，在智能体间影响衰减的假设下，局部邻域观测可替代全局观测，从而支持去中心化。尽管无线通信和雷达网络中的信号衰减现象天然符合这一特性，但其实际应用仍待探索。本文以雷达网络目标检测的功率分配问题为例，提出两种约束MARL模型，推导局部邻域对全局价值函数和梯度的近似估计及其误差界，并开发去中心化策略梯度算法。该方法为无线通信和雷达网络的类似问题提供了可扩展的解决框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T18:02:59Z
- **目录日期**: 2025-05-19
