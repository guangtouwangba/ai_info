# RL-DAUNCE: Reinforcement Learning-Driven Data Assimilation with Uncertainty-Aware Constrained Ensembles

**URL**: http://arxiv.org/abs/2505.05452v1

## 原始摘要

Machine learning has become a powerful tool for enhancing data assimilation.
While supervised learning remains the standard method, reinforcement learning
(RL) offers unique advantages through its sequential decision-making framework,
which naturally fits the iterative nature of data assimilation by dynamically
balancing model forecasts with observations. We develop RL-DAUNCE, a new
RL-based method that enhances data assimilation with physical constraints
through three key aspects. First, RL-DAUNCE inherits the computational
efficiency of machine learning while it uniquely structures its agents to
mirror ensemble members in conventional data assimilation methods. Second,
RL-DAUNCE emphasizes uncertainty quantification by advancing multiple ensemble
members, moving beyond simple mean-state optimization. Third, RL-DAUNCE's
ensemble-as-agents design facilitates the enforcement of physical constraints
during the assimilation process, which is crucial to improving the state
estimation and subsequent forecasting. A primal-dual optimization strategy is
developed to enforce constraints, which dynamically penalizes the reward
function to ensure constraint satisfaction throughout the learning process.
Also, state variable bounds are respected by constraining the RL action space.
Together, these features ensure physical consistency without sacrificing
efficiency. RL-DAUNCE is applied to the Madden-Julian Oscillation, an
intermittent atmospheric phenomenon characterized by strongly non-Gaussian
features and multiple physical constraints. RL-DAUNCE outperforms the standard
ensemble Kalman filter (EnKF), which fails catastrophically due to the
violation of physical constraints. Notably, RL-DAUNCE matches the performance
of constrained EnKF, particularly in recovering intermittent signals, capturing
extreme events, and quantifying uncertainties, while requiring substantially
less computational effort.


## AI 摘要

本文提出了一种基于强化学习的新型数据同化方法RL-DAUNCE，通过三个创新点提升传统方法：1) 采用类似集合成员的智能体设计保持计算效率；2) 通过多集合成员推进实现不确定性量化；3) 创新性地在训练过程中通过对偶优化策略和动作空间约束来保证物理约束。应用于具有强非高斯特征的MJO现象时，RL-DAUNCE显著优于标准集合卡尔曼滤波（EnKF），在恢复间歇信号、捕捉极端事件和量化不确定性方面与约束EnKF相当，但计算量大幅降低。该方法成功解决了物理约束违反问题，同时保持了计算效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-11T00:01:31Z
- **目录日期**: 2025-05-11
