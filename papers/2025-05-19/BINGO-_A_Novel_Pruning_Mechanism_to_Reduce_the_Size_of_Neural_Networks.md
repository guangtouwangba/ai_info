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

过去十年，机器学习模型规模呈指数级增长，导致训练和运行成本高昂，阻碍了个人参与并推高AI服务价格。现有剪枝方法（如迭代幅度剪枝）虽能保持精度，但需多次迭代训练，计算和环境成本极高。为此，研究者提出BINGO方法：在单次训练过程中，通过评估神经网络各权重对模型精度的影响程度，直接生成权重重要性评分，从而一次性剪除冗余权重。BINGO在保证精度的前提下，大幅降低了计算复杂度，为"AI发展不必伴随模型膨胀"提供了可行方案。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T20:02:09Z
- **目录日期**: 2025-05-19
