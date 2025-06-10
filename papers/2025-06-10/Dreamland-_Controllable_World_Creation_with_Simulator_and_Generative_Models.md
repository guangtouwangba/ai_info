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

Dreamland是一个结合物理模拟器精细控制和大规模预训练生成模型真实内容输出的混合世界生成框架。它设计了一个分层世界抽象，将像素级和对象级语义与几何编码为中间表示，以桥接模拟器和生成模型。这种方法增强了可控性，通过与真实世界分布的早期对齐降低适应成本，并支持现有及未来预训练生成模型的即插即用。实验表明，Dreamland在图像质量上提升50.8%，可控性提高17.9%，并能有效增强具身智能体训练。相关代码和数据将公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T05:01:00Z
- **目录日期**: 2025-06-10
