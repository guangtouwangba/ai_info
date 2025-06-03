# Improving Transformer World Models for Data-Efficient RL

**URL**: http://arxiv.org/abs/2502.01591v2

## 原始摘要

We present an approach to model-based RL that achieves a new state of the art
performance on the challenging Craftax-classic benchmark, an open-world 2D
survival game that requires agents to exhibit a wide range of general abilities
-- such as strong generalization, deep exploration, and long-term reasoning.
With a series of careful design choices aimed at improving sample efficiency,
our MBRL algorithm achieves a reward of 69.66% after only 1M environment steps,
significantly outperforming DreamerV3, which achieves 53.2%, and, for the first
time, exceeds human performance of 65.0%. Our method starts by constructing a
SOTA model-free baseline, using a novel policy architecture that combines CNNs
and RNNs. We then add three improvements to the standard MBRL setup: (a) "Dyna
with warmup", which trains the policy on real and imaginary data, (b) "nearest
neighbor tokenizer" on image patches, which improves the scheme to create the
transformer world model (TWM) inputs, and (c) "block teacher forcing", which
allows the TWM to reason jointly about the future tokens of the next timestep.


## AI 摘要

该研究提出了一种基于模型的强化学习方法（MBRL），在Craftax-classic开放世界2D生存游戏中实现了69.66%的奖励表现（仅需100万步训练），首次超越人类水平（65.0%）并显著优于DreamerV3（53.2%）。方法结合了CNN和RNN的新型策略架构作为无模型基线，并引入三项改进：(1) "预热Dyna"混合真实与想象数据训练策略；(2) "最近邻图像块标记器"优化Transformer世界模型(TWM)输入；(3) "块式教师强制"使TWM能联合推理未来标记。这些设计显著提升了样本效率与长期推理能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T05:02:14Z
- **目录日期**: 2025-06-03
