# Dreamland: Controllable World Creation with Simulator and Generative Models

**URL**: http://arxiv.org/abs/2506.08006v1

## 原始摘要

Large-scale video generative models can synthesize diverse and realistic
visual content for dynamic world creation, but they often lack element-wise
controllability, hindering their use in editing scenes and training embodied AI
agents. We propose Dreamland, a hybrid world generation framework combining the
granular control of a physics-based simulator and the photorealistic content
output of large-scale pretrained generative models. In particular, we design a
layered world abstraction that encodes both pixel-level and object-level
semantics and geometry as an intermediate representation to bridge the
simulator and the generative model. This approach enhances controllability,
minimizes adaptation cost through early alignment with real-world
distributions, and supports off-the-shelf use of existing and future pretrained
generative models. We further construct a D3Sim dataset to facilitate the
training and evaluation of hybrid generation pipelines. Experiments demonstrate
that Dreamland outperforms existing baselines with 50.8% improved image
quality, 17.9% stronger controllability, and has great potential to enhance
embodied agent training. Code and data will be made available.


## AI 摘要

Dreamland提出了一种结合物理模拟器精细控制和大规模预训练生成模型逼真输出的混合世界生成框架。其核心是设计了一个分层世界抽象层，通过编码像素级和对象级的语义与几何信息作为中间表示，连接模拟器与生成模型。这种方法提升了可控性，通过早期对齐真实世界分布降低适应成本，并支持现有及未来预训练模型的即插即用。实验表明，Dreamland在图像质量上提升50.8%，可控性增强17.9%，并能显著提升具身智能体训练效果。配套发布的D3Sim数据集支持混合生成管道的训练与评估。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T22:01:06Z
- **目录日期**: 2025-06-10
