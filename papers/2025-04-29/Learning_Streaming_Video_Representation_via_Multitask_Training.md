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

这篇论文提出了StreamFormer，一种用于实时视频流处理的新型架构。该模型基于预训练视觉Transformer，通过引入因果时序注意力机制，实现了高效流式视频处理，同时保持了图像表征能力。作者采用多任务视觉-语言对齐框架统一训练，使模型能同时学习全局语义、时序动态和细粒度空间关系。实验表明，StreamFormer在在线动作检测、视频实例分割和视频问答等任务中取得了有竞争力的结果，同时保持了高效性，展现了在自动驾驶、具身AI等实时应用中的潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T11:01:00Z
- **目录日期**: 2025-04-29
