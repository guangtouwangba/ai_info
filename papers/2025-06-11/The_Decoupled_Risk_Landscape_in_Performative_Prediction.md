# The Decoupled Risk Landscape in Performative Prediction

**URL**: http://arxiv.org/abs/2506.09044v1

## 原始摘要

Performative Prediction addresses scenarios where deploying a model induces a
distribution shift in the input data, such as individuals modifying their
features and reapplying for a bank loan after rejection. Literature has had a
theoretical perspective giving mathematical guarantees for convergence (either
to the stable or optimal point). We believe that visualization of the loss
landscape can complement this theoretical advances with practical insights.
Therefore, (1) we introduce a simple decoupled risk visualization method
inspired in the two-step process that performative prediction is. Our approach
visualizes the risk landscape with respect to two parameter vectors: model
parameters and data parameters. We use this method to propose new properties of
the interest points, to examine how existing algorithms traverse the risk
landscape and perform under more realistic conditions, including strategic
classification with non-linear models. (2) Building on this decoupled risk
visualization, we introduce a novel setting - extended Performative Prediction
- which captures scenarios where the distribution reacts to a model different
from the decision-making one, reflecting the reality that agents often lack
full access to the deployed model.


## AI 摘要

本文提出了一种可视化方法来分析执行预测（Performative Prediction）中的损失景观，以补充现有理论研究的不足。该方法将风险景观解耦为模型参数和数据参数两个维度进行可视化分析，揭示了关键点的性质，并评估了现有算法在非线性战略分类等实际场景中的表现。基于此，作者进一步提出了"扩展执行预测"的新设定，用于描述代理无法完全访问部署模型时，数据分布对非决策模型的反应情况。这一工作为理解模型部署引发的分布偏移提供了实践视角。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T19:01:34Z
- **目录日期**: 2025-06-11
