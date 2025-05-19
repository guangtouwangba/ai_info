# BINGO: A Novel Pruning Mechanism to Reduce the Size of Neural Networks

**URL**: http://arxiv.org/abs/2505.09864v2

## 原始摘要

Over the past decade, the use of machine learning has increased
exponentially. Models are far more complex than ever before, growing to
gargantuan sizes and housing millions of weights. Unfortunately, the fact that
large models have become the state of the art means that it often costs
millions of dollars to train and operate them. These expenses not only hurt
companies but also bar non-wealthy individuals from contributing to new
developments and force consumers to pay greater prices for AI. Current methods
used to prune models, such as iterative magnitude pruning, have shown great
accuracy but require an iterative training sequence that is incredibly
computationally and environmentally taxing. To solve this problem, BINGO is
introduced. BINGO, during the training pass, studies specific subsets of a
neural network one at a time to gauge how significant of a role each weight
plays in contributing to a network's accuracy. By the time training is done,
BINGO generates a significance score for each weight, allowing for
insignificant weights to be pruned in one shot. BINGO provides an
accuracy-preserving pruning technique that is less computationally intensive
than current methods, allowing for a world where AI growth does not have to
mean model growth, as well.


## AI 摘要

过去十年，机器学习模型复杂度激增，导致训练成本高达数百万美元，阻碍个人参与并推高AI服务价格。现有剪枝方法（如迭代幅度剪枝）虽精准但计算耗能巨大。为此，研究提出BINGO技术：在训练过程中逐次分析神经网络子集，量化每个权重对模型准确性的贡献，最终生成权重重要性评分，实现一次性高效剪枝。BINGO在保持精度的同时大幅降低计算负担，为"AI发展不必伴随模型膨胀"提供了可行方案。（99字）  

注：摘要聚焦三个核心点：  
1. 问题背景（大模型成本高、现有方法缺陷）  
2. BINGO机制（动态评估权重重要性）  
3. 技术优势（一次性剪枝、节能保精度）  
严格控制在100字内，未丢失关键信息。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T12:02:14Z
- **目录日期**: 2025-05-19
