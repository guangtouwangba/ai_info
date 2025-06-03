# AdaWorld: Learning Adaptable World Models with Latent Actions

**URL**: http://arxiv.org/abs/2503.18938v4

## 原始摘要

World models aim to learn action-controlled future prediction and have proven
essential for the development of intelligent agents. However, most existing
world models rely heavily on substantial action-labeled data and costly
training, making it challenging to adapt to novel environments with
heterogeneous actions through limited interactions. This limitation can hinder
their applicability across broader domains. To overcome this limitation, we
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

AdaWorld提出了一种创新的世界模型学习方法，通过自监督方式从视频中提取潜在动作信息，捕捉帧间关键变化，并构建基于这些潜在动作的自回归世界模型。相比依赖大量动作标记数据的传统方法，AdaWorld能在有限交互和微调下高效适应新环境中的异构动作。实验表明，AdaWorld在模拟质量和视觉规划方面均表现优异，实现了跨领域的高效迁移学习。该方法降低了模型对标注数据和训练成本的依赖，提升了智能体在新环境中的适应能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T21:02:25Z
- **目录日期**: 2025-06-03
