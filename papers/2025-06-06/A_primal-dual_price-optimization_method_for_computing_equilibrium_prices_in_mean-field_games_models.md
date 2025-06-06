# A primal-dual price-optimization method for computing equilibrium prices in mean-field games models

**URL**: http://arxiv.org/abs/2506.04169v1

## 原始摘要

We develop a simple yet efficient Lagrangian method for computing equilibrium
prices in a mean-field game price-formation model. We prove that equilibrium
prices are optimal in terms of a suitable criterion and derive a primal-dual
gradient-based algorithm for computing them. One of the highlights of our
computational framework is the efficient, simple, and flexible implementation
of the algorithm using modern automatic differentiation techniques. Our
implementation is modular and admits a seamless extension to high-dimensional
settings with more complex dynamics, costs, and equilibrium conditions.
Additionally, automatic differentiation enables a versatile algorithm that
requires only coding the cost functions of agents. It automatically handles the
gradients of the costs, thereby eliminating the need to manually form the
adjoint equations.


## AI 摘要

本文提出了一种基于拉格朗日方法的简单高效算法，用于计算均值场博弈价格形成模型中的均衡价格。研究证明这些均衡价格在特定标准下是最优的，并开发了一种基于原始-对偶梯度的方法进行计算。该计算框架的关键优势在于利用现代自动微分技术实现了高效、简单且灵活的实现方案。该实现具有模块化特点，可无缝扩展到具有更复杂动态、成本和均衡条件的高维场景。自动微分技术使算法仅需编写代理成本函数即可自动处理梯度计算，无需手动构建伴随方程。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T02:31:22Z
- **目录日期**: 2025-06-06
