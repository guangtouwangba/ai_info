# TesserAct: Learning 4D Embodied World Models

**URL**: http://arxiv.org/abs/2504.20995v1

## 原始摘要

This paper presents an effective approach for learning novel 4D embodied
world models, which predict the dynamic evolution of 3D scenes over time in
response to an embodied agent's actions, providing both spatial and temporal
consistency. We propose to learn a 4D world model by training on RGB-DN (RGB,
Depth, and Normal) videos. This not only surpasses traditional 2D models by
incorporating detailed shape, configuration, and temporal changes into their
predictions, but also allows us to effectively learn accurate inverse dynamic
models for an embodied agent. Specifically, we first extend existing robotic
manipulation video datasets with depth and normal information leveraging
off-the-shelf models. Next, we fine-tune a video generation model on this
annotated dataset, which jointly predicts RGB-DN (RGB, Depth, and Normal) for
each frame. We then present an algorithm to directly convert generated RGB,
Depth, and Normal videos into a high-quality 4D scene of the world. Our method
ensures temporal and spatial coherence in 4D scene predictions from embodied
scenarios, enables novel view synthesis for embodied environments, and
facilitates policy learning that significantly outperforms those derived from
prior video-based world models.


## AI 摘要

本文提出了一种学习新型4D具身世界模型的有效方法，通过训练RGB-DN（RGB、深度和法线）视频预测3D场景随时间的动态演变。该方法超越了传统2D模型，将形状、构型和时间变化纳入预测，并能准确学习具身智能体的逆向动力学模型。研究首先利用现成模型扩展机器人操作视频数据集，添加深度和法线信息；然后微调视频生成模型以联合预测每帧的RGB-DN；最后提出算法将生成的RGB、深度和法线视频直接转换为高质量4D场景。该方法确保了时空一致性，支持新视角合成，并显著提升了基于视频世界模型的策略学习效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T21:01:03Z
- **目录日期**: 2025-04-30
