# Adapting a World Model for Trajectory Following in a 3D Game

**URL**: http://arxiv.org/abs/2504.12299v1

## 原始摘要

Imitation learning is a powerful tool for training agents by leveraging
expert knowledge, and being able to replicate a given trajectory is an integral
part of it. In complex environments, like modern 3D video games, distribution
shift and stochasticity necessitate robust approaches beyond simple action
replay. In this study, we apply Inverse Dynamics Models (IDM) with different
encoders and policy heads to trajectory following in a modern 3D video game --
Bleeding Edge. Additionally, we investigate several future alignment strategies
that address the distribution shift caused by the aleatoric uncertainty and
imperfections of the agent. We measure both the trajectory deviation distance
and the first significant deviation point between the reference and the agent's
trajectory and show that the optimal configuration depends on the chosen
setting. Our results show that in a diverse data setting, a GPT-style policy
head with an encoder trained from scratch performs the best, DINOv2 encoder
with the GPT-style policy head gives the best results in the low data regime,
and both GPT-style and MLP-style policy heads had comparable results when
pre-trained on a diverse setting and fine-tuned for a specific behaviour
setting.


## AI 摘要

本研究探讨了在复杂3D游戏环境中使用逆动力学模型（IDM）进行轨迹跟随的方法。针对分布偏移和随机性问题，研究测试了不同编码器和策略头的组合，并提出了多种未来对齐策略。实验在《Bleeding Edge》游戏中进行，通过测量轨迹偏差距离和首次显著偏离点评估性能。结果显示：在多样化数据场景下，从头训练的GPT风格策略头表现最佳；低数据量时，DINOv2编码器配合GPT策略头最优；而在预训练后微调的场景中，GPT和MLP风格策略头表现相当。研究为游戏AI的模仿学习提供了实用见解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-18T01:28:16Z
- **目录日期**: 2025-04-18
