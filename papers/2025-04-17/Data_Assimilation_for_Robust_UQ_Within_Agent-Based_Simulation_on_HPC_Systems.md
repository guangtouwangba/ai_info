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

该研究提出了一种基于序贯蒙特卡罗采样的贝叶斯不确定性量化框架，用于改进基于智能体的传染病传播模型。该框架通过将智能体模拟与贝叶斯建模相结合，能够实时整合静态或流式数据，量化疾病传播参数的不确定性，为公共卫生决策提供可能结果的范围。这一方法解决了传统智能体模型缺乏不确定性量化和实时数据同化能力的局限性，同时保持了模型的灵活性。研究还展示了该框架在高性能计算系统上的可扩展性，并通过模拟疫情爆发案例验证了其在静态和流式数据场景下的应用效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T18:01:43Z
- **目录日期**: 2025-04-17
