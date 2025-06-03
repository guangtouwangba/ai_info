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

我们提出了一种基于模型的强化学习方法，在Craftax-classic开放世界2D生存游戏基准测试中取得了新突破。该方法通过改进样本效率的关键设计，仅用100万步训练就达到69.66%的奖励分数，显著超过DreamerV3的53.2%和人类水平的65%。核心创新包括：(1)结合CNN和RNN的新型策略架构；(2)"预热Dyna"混合真实与想象数据训练；(3)基于图像块的最近邻标记化方法改进Transformer世界模型输入；(4)"块式教师强制"技术增强未来时间步的联合推理能力。这些改进使算法在泛化、探索和长期推理方面表现优异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T16:03:00Z
- **目录日期**: 2025-06-03
