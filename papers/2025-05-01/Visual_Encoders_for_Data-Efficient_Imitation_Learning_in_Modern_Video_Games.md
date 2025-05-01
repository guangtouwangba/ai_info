# Visual Encoders for Data-Efficient Imitation Learning in Modern Video Games

**URL**: http://arxiv.org/abs/2312.02312v2

## 原始摘要

Video games have served as useful benchmarks for the decision-making
community, but going beyond Atari games towards modern games has been
prohibitively expensive for the vast majority of the research community. Prior
work in modern video games typically relied on game-specific integration to
obtain game features and enable online training, or on existing large datasets.
An alternative approach is to train agents using imitation learning to play
video games purely from images. However, this setting poses a fundamental
question: which visual encoders obtain representations that retain information
critical for decision making? To answer this question, we conduct a systematic
study of imitation learning with publicly available pre-trained visual encoders
compared to the typical task-specific end-to-end training approach in
Minecraft, Counter-Strike: Global Offensive, and Minecraft Dungeons. Our
results show that end-to-end training can be effective with comparably
low-resolution images and only minutes of demonstrations, but significant
improvements can be gained by utilising pre-trained encoders such as DINOv2
depending on the game. In addition to enabling effective decision making, we
show that pre-trained encoders can make decision-making research in video games
more accessible by significantly reducing the cost of training.


## AI 摘要

研究表明，在视频游戏决策任务中，使用预训练视觉编码器（如DINOv2）可以显著提升模仿学习的效果，同时降低训练成本。通过对比《我的世界》《CS:GO》等游戏中的端到端训练方法，发现虽然低分辨率图像和少量演示数据也能实现有效训练，但预训练编码器能更好地保留决策关键信息。这一发现为更广泛的研究社区提供了经济高效的替代方案，无需依赖游戏特定集成或大规模数据集，仅需图像输入即可训练智能体，使视频游戏决策研究更具可及性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T20:02:19Z
- **目录日期**: 2025-05-01
