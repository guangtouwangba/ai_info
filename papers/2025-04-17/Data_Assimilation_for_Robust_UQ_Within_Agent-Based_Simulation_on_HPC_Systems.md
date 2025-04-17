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

本文提出了一种基于序贯蒙特卡罗采样的贝叶斯不确定性量化框架，用于改进基于智能体的传染病传播模拟模型。该框架通过将智能体模型与贝叶斯建模相结合，能够处理静态和实时流数据，量化疾病传播参数的不确定性，为公共卫生决策提供更可靠的预测。相比传统多次运行取统计值的方法，新框架解决了计算效率低和数据同化能力不足两大局限，在保持模型灵活性的同时实现了高效的不确定性量化。研究展示了该框架在模拟疫情暴发中的应用效果，证明其能减少模型误差并提供更现实的预测范围。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T11:02:36Z
- **目录日期**: 2025-04-17
