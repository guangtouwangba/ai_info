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

研究人员开发了RL-DAUNCE，一种基于强化学习的数据同化新方法。该方法通过三个关键创新提升传统技术：1）采用类似集合成员的智能体结构保持计算效率；2）通过多集合成员推进实现不确定性量化；3）独特设计允许在过程中实施物理约束。研究者采用原始对偶优化策略动态调整奖励函数，确保约束满足。应用于具有强非高斯特征的Madden-Julian振荡现象时，RL-DAUNCE显著优于标准集合卡尔曼滤波（EnKF），性能与计算量更大的约束EnKF相当，在恢复间歇信号、捕捉极端事件和量化不确定性方面表现突出，同时大幅降低计算成本。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T09:01:32Z
- **目录日期**: 2025-05-09
