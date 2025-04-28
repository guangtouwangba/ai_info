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

本文研究了分布式多视角表示学习问题，探讨了在无显式协调的情况下，各智能体如何从统计相关的不同视角中提取必要且充分的特征表示，以帮助解码器正确估计隐藏标签。作者从泛化误差角度出发，建立了基于相对熵的泛化界，并提出了一种正则化方法。实验表明，采用数据相关的高斯混合先验能取得良好性能，在单视角场景下优于现有VIB和CDVIB方法，并自然产生了加权注意力机制。多视角研究中发现，选择高斯乘积混合作为联合先验会隐式鼓励智能体提取冗余特征，这一反直觉现象值得关注。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T05:02:23Z
- **目录日期**: 2025-04-28
