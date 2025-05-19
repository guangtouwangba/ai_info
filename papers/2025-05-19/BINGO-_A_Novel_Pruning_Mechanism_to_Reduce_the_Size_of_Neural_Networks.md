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

过去十年，机器学习应用激增，模型复杂度与参数量大幅提升，导致训练成本高达数百万美元，阻碍了非富裕群体参与AI研发并推高用户成本。现有剪枝方法（如迭代幅度剪枝）虽能保持精度，但需要高计算量的迭代训练。为此，研究者提出BINGO方法：在训练过程中逐次分析神经网络子集，量化每个权重对模型精度的重要性，最终生成权重重要性评分，实现一次性高效剪枝。BINGO在降低计算负担的同时保持模型精度，为"AI发展不必伴随模型膨胀"提供了可行方案。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T16:02:05Z
- **目录日期**: 2025-05-19
