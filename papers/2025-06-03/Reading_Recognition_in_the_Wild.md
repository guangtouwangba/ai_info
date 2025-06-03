# Reading Recognition in the Wild

**URL**: http://arxiv.org/abs/2505.24848v1

## 原始摘要

To enable egocentric contextual AI in always-on smart glasses, it is crucial
to be able to keep a record of the user's interactions with the world,
including during reading. In this paper, we introduce a new task of reading
recognition to determine when the user is reading. We first introduce the
first-of-its-kind large-scale multimodal Reading in the Wild dataset,
containing 100 hours of reading and non-reading videos in diverse and realistic
scenarios. We then identify three modalities (egocentric RGB, eye gaze, head
pose) that can be used to solve the task, and present a flexible transformer
model that performs the task using these modalities, either individually or
combined. We show that these modalities are relevant and complementary to the
task, and investigate how to efficiently and effectively encode each modality.
Additionally, we show the usefulness of this dataset towards classifying types
of reading, extending current reading understanding studies conducted in
constrained settings to larger scale, diversity and realism. Code, model, and
data will be public.


## AI 摘要

本文提出了一项新任务——阅读识别，用于判断用户是否在阅读，以支持智能眼镜中的情境感知AI。研究者首先发布了首个大规模多模态数据集"Reading in the Wild"，包含100小时不同真实场景下的阅读和非阅读视频。通过分析三种模态（第一视角RGB、视线追踪和头部姿态），开发了一个灵活的Transformer模型，可单独或组合使用这些模态。研究表明这些模态对任务具有相关性和互补性，并探讨了每种模态的有效编码方式。该数据集还能扩展当前受限环境下的阅读理解研究，支持更丰富的阅读类型分类。相关代码、模型和数据将公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T01:29:11Z
- **目录日期**: 2025-06-03
