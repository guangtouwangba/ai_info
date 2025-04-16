# eXplainable AI for data driven control: an inverse optimal control approach

**URL**: http://arxiv.org/abs/2504.11446v1

## 原始摘要

Understanding the behavior of black-box data-driven controllers is a key
challenge in modern control design. In this work, we propose an eXplainable AI
(XAI) methodology based on Inverse Optimal Control (IOC) to obtain local
explanations for the behavior of a controller operating around a given region.
Specifically, we extract the weights assigned to tracking errors and control
effort in the implicit cost function that a black-box controller is optimizing,
offering a more transparent and interpretable representation of the
controller's underlying objectives. This approach presents connections with
well-established XAI techniques, such as Local Interpretable Model-agnostic
Explanations (LIME) since it is still based on a local approximation of the
control policy. However, rather being limited to a standard sensitivity
analysis, the explanation provided by our method relies on the solution of an
inverse Linear Quadratic (LQ) problem, offering a structured and more
control-relevant perspective. Numerical examples demonstrate that the inferred
cost function consistently provides a deeper understanding of the controller's
decision-making process, shedding light on otherwise counterintuitive or
unexpected phenomena.


## AI 摘要

本文提出了一种基于逆向最优控制(IOC)的可解释AI(XAI)方法，用于分析黑盒数据驱动控制器在特定区域的局部行为。该方法通过提取控制器隐含成本函数中的跟踪误差和控制力度权重，揭示其底层优化目标，从而提高控制决策的透明度。与LIME等传统XAI技术相比，该方法通过求解逆向线性二次(LQ)问题，提供了更具结构化和控制相关性的解释。数值实验表明，推断出的成本函数能有效增强对控制器决策过程的理解，特别是对一些反直觉现象的解释。该方法为现代控制设计中的黑盒系统行为分析提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T08:02:15Z
- **目录日期**: 2025-04-16
