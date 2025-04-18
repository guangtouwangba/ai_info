# Data Assimilation for Robust UQ Within Agent-Based Simulation on HPC Systems

**URL**: http://arxiv.org/abs/2504.12228v1

## 原始摘要

Agent-based simulation provides a powerful tool for in silico system
modeling. However, these simulations do not provide built-in methods for
uncertainty quantification (UQ). Within these types of models a typical
approach to UQ is to run multiple realizations of the model then compute
aggregate statistics. This approach is limited due to the compute time required
for a solution. When faced with an emerging biothreat, public health decisions
need to be made quickly and solutions for integrating near real-time data with
analytic tools are needed.
  We propose an integrated Bayesian UQ framework for agent-based models based
on sequential Monte Carlo sampling. Given streaming or static data about the
evolution of an emerging pathogen, this Bayesian framework provides a
distribution over the parameters governing the spread of a disease through a
population. These estimates of the spread of a disease may be provided to
public health agencies seeking to abate the spread.
  By coupling agent-based simulations with Bayesian modeling in a data
assimilation, our proposed framework provides a powerful tool for modeling
dynamical systems in silico. We propose a method which reduces model error and
provides a range of realistic possible outcomes. Moreover, our method addresses
two primary limitations of ABMs: the lack of UQ and an inability to assimilate
data. Our proposed framework combines the flexibility of an agent-based model
with UQ provided by the Bayesian paradigm in a workflow which scales well to
HPC systems. We provide algorithmic details and results on a simulated outbreak
with both static and streaming data.


## AI 摘要

本文提出了一种基于序贯蒙特卡罗采样的贝叶斯不确定性量化框架，用于解决基于智能体模型(ABM)缺乏不确定性量化(UQ)和数据同化能力的问题。该框架通过将ABM与贝叶斯建模相结合，能够在静态或实时流数据条件下，快速估计疾病传播参数的概率分布，为突发公共卫生事件决策提供支持。该方法降低了模型误差，能生成一系列可能的现实结果，同时解决了ABM的两大局限：缺乏UQ能力和无法同化数据。该框架兼具ABM的灵活性和贝叶斯方法的UQ优势，且适用于高性能计算系统。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-18T04:01:36Z
- **目录日期**: 2025-04-18
