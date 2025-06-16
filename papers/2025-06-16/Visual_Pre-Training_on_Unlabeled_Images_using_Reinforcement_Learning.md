# Visual Pre-Training on Unlabeled Images using Reinforcement Learning

**URL**: http://arxiv.org/abs/2506.11967v1

## 原始摘要

In reinforcement learning (RL), value-based algorithms learn to associate
each observation with the states and rewards that are likely to be reached from
it. We observe that many self-supervised image pre-training methods bear
similarity to this formulation: learning features that associate crops of
images with those of nearby views, e.g., by taking a different crop or color
augmentation. In this paper, we complete this analogy and explore a method that
directly casts pre-training on unlabeled image data like web crawls and video
frames as an RL problem. We train a general value function in a dynamical
system where an agent transforms an image by changing the view or adding image
augmentations. Learning in this way resembles crop-consistency
self-supervision, but through the reward function, offers a simple lever to
shape feature learning using curated images or weakly labeled captions when
they exist. Our experiments demonstrate improved representations when training
on unlabeled images in the wild, including video data like EpicKitchens, scene
data like COCO, and web-crawl data like CC12M.


## AI 摘要

这篇论文提出了一种将自监督图像预训练类比为强化学习(RL)问题的新方法。研究者观察到，许多自监督预训练方法(如图像裁剪一致性)与RL中的价值函数学习具有相似性：都学习从当前观察预测未来状态的特征。基于此，他们将无标签图像预训练直接建模为RL问题，通过设计动态系统让智能体通过变换视图或添加图像增强来学习通用价值函数。这种方法不仅保留了裁剪一致性的优点，还能通过奖励函数利用标注数据优化特征学习。实验表明，该方法在EpicKitchens、COCO和CC12M等数据集上能学习到更好的表征。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T02:34:06Z
- **目录日期**: 2025-06-16
