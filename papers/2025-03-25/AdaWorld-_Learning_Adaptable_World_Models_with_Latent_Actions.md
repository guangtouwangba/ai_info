# AdaWorld: Learning Adaptable World Models with Latent Actions

**URL**: http://arxiv.org/abs/2503.18938v1

## 原始摘要

World models aim to learn action-controlled prediction models and have proven
essential for the development of intelligent agents. However, most existing
world models rely heavily on substantial action-labeled data and costly
training, making it challenging to adapt to novel environments with
heterogeneous actions through limited interactions. This limitation can hinder
their applicability across broader domains. To overcome this challenge, we
propose AdaWorld, an innovative world model learning approach that enables
efficient adaptation. The key idea is to incorporate action information during
the pretraining of world models. This is achieved by extracting latent actions
from videos in a self-supervised manner, capturing the most critical
transitions between frames. We then develop an autoregressive world model that
conditions on these latent actions. This learning paradigm enables highly
adaptable world models, facilitating efficient transfer and learning of new
actions even with limited interactions and finetuning. Our comprehensive
experiments across multiple environments demonstrate that AdaWorld achieves
superior performance in both simulation quality and visual planning.


## AI 摘要

AdaWorld是一种创新的世界模型学习方法，通过自监督方式从视频中提取潜在动作信息，捕捉帧间关键变化，并在预训练阶段融入动作信息。该方法构建了一个基于潜在动作的自回归世界模型，显著提升了模型在新环境中的适应能力。相比传统方法依赖大量标注数据和昂贵训练，AdaWorld仅需有限交互和微调即可高效学习新动作，实现跨领域迁移。实验表明，AdaWorld在仿真质量和视觉规划任务中均表现出优越性能，解决了现有世界模型在异构动作环境中适应性不足的核心挑战。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T19:01:07Z
- **目录日期**: 2025-03-25
