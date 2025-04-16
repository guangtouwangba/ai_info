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

本文提出了一种基于逆向最优控制（IOC）的可解释AI（XAI）方法，用于分析黑盒控制器在特定区域的局部行为。该方法通过提取控制器隐式优化成本函数中的误差跟踪和控制权重，提供更透明的决策解释。与LIME等现有XAI技术类似，它基于局部近似，但通过求解逆向线性二次（LQ）问题，提供了更具控制相关性的结构化解释，而非简单的敏感性分析。数值实验表明，推断的成本函数能深入揭示控制器的决策机制，解释其反直觉行为。该方法为理解数据驱动控制器的黑盒特性提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T23:01:52Z
- **目录日期**: 2025-04-16
