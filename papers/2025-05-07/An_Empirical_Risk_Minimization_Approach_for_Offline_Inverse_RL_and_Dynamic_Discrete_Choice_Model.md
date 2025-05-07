# An Empirical Risk Minimization Approach for Offline Inverse RL and Dynamic Discrete Choice Model

**URL**: http://arxiv.org/abs/2502.14131v3

## 原始摘要

We study the problem of estimating Dynamic Discrete Choice (DDC) models, also
known as offline Maximum Entropy-Regularized Inverse Reinforcement Learning
(offline MaxEnt-IRL) in machine learning. The objective is to recover reward or
$Q^*$ functions that govern agent behavior from offline behavior data. In this
paper, we propose a globally convergent gradient-based method for solving these
problems without the restrictive assumption of linearly parameterized rewards.
The novelty of our approach lies in introducing the Empirical Risk Minimization
(ERM) based IRL/DDC framework, which circumvents the need for explicit state
transition probability estimation in the Bellman equation. Furthermore, our
method is compatible with non-parametric estimation techniques such as neural
networks. Therefore, the proposed method has the potential to be scaled to
high-dimensional, infinite state spaces. A key theoretical insight underlying
our approach is that the Bellman residual satisfies the Polyak-Lojasiewicz (PL)
condition -- a property that, while weaker than strong convexity, is sufficient
to ensure fast global convergence guarantees. Through a series of synthetic
experiments, we demonstrate that our approach consistently outperforms
benchmark methods and state-of-the-art alternatives.


## AI 摘要

本文提出了一种基于梯度下降的全局收敛方法，用于估计动态离散选择(DDC)模型或离线最大熵正则化逆向强化学习(offline MaxEnt-IRL)。该方法创新性地采用基于经验风险最小化(ERM)的框架，避免了贝尔曼方程中显式状态转移概率估计的需求，且兼容神经网络等非参数估计技术，适用于高维无限状态空间。理论关键发现是贝尔曼残差满足Polyak-Lojasiewicz(PL)条件，虽弱于强凸性但足以保证快速全局收敛。合成实验表明，该方法在线性和非线性参数化设置下均优于基准方法和现有最优方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T05:02:33Z
- **目录日期**: 2025-05-07
