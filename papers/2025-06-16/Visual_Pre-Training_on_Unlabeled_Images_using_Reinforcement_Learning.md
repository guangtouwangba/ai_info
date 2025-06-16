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

这篇论文提出了一种将无监督图像预训练视为强化学习问题的新方法。研究者发现，自监督图像预训练与基于价值的强化学习算法存在相似性——两者都学习将输入（图像裁剪或观测）与潜在状态（邻近视图或未来状态）相关联。他们设计了一个动态系统，通过代理改变图像视图或添加增强来训练通用价值函数。这种方法类似于裁剪一致性自监督学习，但通过奖励函数可以利用精选图像或弱标签来引导特征学习。实验表明，该方法在EpicKitchens、COCO和CC12M等无标注数据集上能学习到更好的表征。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T03:26:16Z
- **目录日期**: 2025-06-16
