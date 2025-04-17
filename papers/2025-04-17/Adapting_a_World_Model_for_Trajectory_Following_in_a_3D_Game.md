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

模仿学习通过专家知识训练智能体，轨迹复现是其核心环节。针对3D游戏等复杂环境中的分布偏移和随机性，本研究在游戏《Bleeding Edge》中测试了不同编码器和策略头的逆向动力学模型（IDM），并探索了应对随机不确定性和智能体缺陷的未来对齐策略。实验通过轨迹偏离距离和首次显著偏离点评估性能，结果表明：多样数据场景下，GPT式策略头配合从头训练的编码器最优；低数据量时DINOv2编码器配GPT策略头最佳；在预训练后微调场景中，GPT式与MLP式策略头表现相当。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T11:01:40Z
- **目录日期**: 2025-04-17
