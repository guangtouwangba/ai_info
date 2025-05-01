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

该研究探讨了在视频游戏决策任务中视觉编码器的选择问题。通过对比公开预训练视觉编码器与端到端训练方法在《我的世界》、《CS:GO》等游戏中的表现，发现端到端训练在低分辨率图像和少量演示下仍有效，但预训练编码器（如DINOv2）能显著提升性能。预训练编码器不仅能提升决策效果，还大幅降低了训练成本，使基于视觉的决策研究更易普及。研究为游戏AI领域提供了实用的视觉表示选择指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T19:02:14Z
- **目录日期**: 2025-05-01
