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

模仿学习通过专家知识训练智能体，轨迹复现是其关键环节。在复杂3D游戏（如《Bleeding Edge》）中，分布偏移和随机性需要超越简单动作重放的鲁棒方法。本研究采用不同编码器和策略头的逆动力学模型（IDM）进行轨迹跟踪，并探索多种未来对齐策略以解决随机不确定性和智能体缺陷导致的分布偏移。实验通过轨迹偏差距离和首次显著偏离点评估性能，结果显示：多样数据场景下，GPT风格策略头+从头训练编码器最优；少数据时DINOv2编码器+GPT头最佳；预训练后微调时，GPT头和MLP头表现相当。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T22:00:54Z
- **目录日期**: 2025-04-17
