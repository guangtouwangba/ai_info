# Generalization Guarantees for Multi-View Representation Learning and Application to Regularization via Gaussian Product Mixture Prior

**URL**: http://arxiv.org/abs/2504.18455v1

## 原始摘要

We study the problem of distributed multi-view representation learning. In
this problem, $K$ agents observe each one distinct, possibly statistically
correlated, view and independently extracts from it a suitable representation
in a manner that a decoder that gets all $K$ representations estimates
correctly the hidden label. In the absence of any explicit coordination between
the agents, a central question is: what should each agent extract from its view
that is necessary and sufficient for a correct estimation at the decoder? In
this paper, we investigate this question from a generalization error
perspective. First, we establish several generalization bounds in terms of the
relative entropy between the distribution of the representations extracted from
training and "test" datasets and a data-dependent symmetric prior, i.e., the
Minimum Description Length (MDL) of the latent variables for all views and
training and test datasets. Then, we use the obtained bounds to devise a
regularizer; and investigate in depth the question of the selection of a
suitable prior. In particular, we show and conduct experiments that illustrate
that our data-dependent Gaussian mixture priors with judiciously chosen weights
lead to good performance. For single-view settings (i.e., $K=1$), our
experimental results are shown to outperform existing prior art Variational
Information Bottleneck (VIB) and Category-Dependent VIB (CDVIB) approaches.
Interestingly, we show that a weighted attention mechanism emerges naturally in
this setting. Finally, for the multi-view setting, we show that the selection
of the joint prior as a Gaussians product mixture induces a Gaussian mixture
marginal prior for each marginal view and implicitly encourages the agents to
extract and output redundant features, a finding which is somewhat
counter-intuitive.


## AI 摘要

本文研究了分布式多视图表示学习问题，探讨了在无显式协调的情况下，各代理应如何从相关视图中提取必要且充分的特征以支持解码器正确预测标签。通过泛化误差分析，作者提出了基于相对熵和最小描述长度（MDL）的泛化界限，并设计了一种正则化方法。实验表明，数据依赖的高斯混合先验优于现有变分信息瓶颈方法（VIB/CDVIB），且单视图场景中自然涌现加权注意力机制。多视图场景下，高斯乘积混合先验会隐式促使代理提取冗余特征，这一反直觉发现为多视图学习提供了新视角。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T14:01:52Z
- **目录日期**: 2025-04-28
