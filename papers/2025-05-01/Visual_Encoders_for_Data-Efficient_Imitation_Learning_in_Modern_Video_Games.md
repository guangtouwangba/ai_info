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

电子游戏是决策AI研究的重要测试平台，但现代游戏的开发成本过高。本研究对比了两种视觉编码方法：端到端训练和预训练编码器（如DINOv2）在《我的世界》《CS:GO》等游戏中的模仿学习效果。结果显示，端到端训练虽能用低分辨率图像和少量演示有效学习，但预训练编码器表现更优，能显著降低研究成本。预训练方法不仅提升决策效果，还使视频游戏决策研究更易普及，为资源有限的研究者提供了可行方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T07:02:22Z
- **目录日期**: 2025-05-01
