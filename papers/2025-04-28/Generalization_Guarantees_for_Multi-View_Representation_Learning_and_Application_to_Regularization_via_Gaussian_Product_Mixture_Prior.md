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

本文研究了分布式多视图表示学习问题，探讨了在无显式协调下各代理应如何提取必要且充分的视图特征以支持解码器正确估计隐藏标签。通过泛化误差视角，作者建立了基于相对熵的泛化界，并提出了一种正则化方法。研究发现，采用数据依赖的高斯混合先验能提升性能，在单视图任务中表现优于现有VIB和CDVIB方法，并自然产生加权注意力机制。对于多视图场景，联合高斯乘积混合先验会促使代理提取冗余特征，这一反直觉发现为多视图学习提供了新见解。实验验证了所提方法的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T23:01:30Z
- **目录日期**: 2025-04-28
