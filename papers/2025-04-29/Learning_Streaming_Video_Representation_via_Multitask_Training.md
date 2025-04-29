# Learning Streaming Video Representation via Multitask Training

**URL**: http://arxiv.org/abs/2504.20041v1

## 原始摘要

Understanding continuous video streams plays a fundamental role in real-time
applications including embodied AI and autonomous driving. Unlike offline video
understanding, streaming video understanding requires the ability to process
video streams frame by frame, preserve historical information, and make
low-latency decisions.To address these challenges, our main contributions are
three-fold. (i) We develop a novel streaming video backbone, termed as
StreamFormer, by incorporating causal temporal attention into a pre-trained
vision transformer. This enables efficient streaming video processing while
maintaining image representation capability.(ii) To train StreamFormer, we
propose to unify diverse spatial-temporal video understanding tasks within a
multitask visual-language alignment framework. Hence, StreamFormer learns
global semantics, temporal dynamics, and fine-grained spatial relationships
simultaneously. (iii) We conduct extensive experiments on online action
detection, online video instance segmentation, and video question answering.
StreamFormer achieves competitive results while maintaining efficiency,
demonstrating its potential for real-time applications.


## AI 摘要

本文提出StreamFormer，一种新型流式视频处理架构，通过将因果时间注意力机制集成到预训练视觉Transformer中，实现了高效流式处理并保持图像表示能力。采用多任务视觉-语言对齐框架统一训练，同时学习全局语义、时序动态和细粒度空间关系。在在线动作检测、视频实例分割和视频问答任务上的实验表明，StreamFormer在保持高效性的同时取得了有竞争力的结果，展现了其在实时应用中的潜力。该研究解决了流式视频理解中的帧处理、历史信息保存和低延迟决策等关键挑战。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T15:00:57Z
- **目录日期**: 2025-04-29
