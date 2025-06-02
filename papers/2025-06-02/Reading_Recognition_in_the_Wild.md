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

本文提出了一种用于智能眼镜的阅读识别任务，通过判断用户是否在阅读来记录其与世界的互动。研究团队首次构建了大规模多模态数据集"Reading in the Wild"，包含100小时真实场景下的阅读/非阅读视频。研究确定了三种相关且互补的模态数据(第一视角RGB、眼动追踪、头部姿态)，并提出灵活的Transformer模型处理单模态或多模态组合。实验验证了各模态的有效性，并探索了高效编码方法。该数据集还能扩展用于阅读类型分类，将现有受限环境下的阅读理解研究推向更大规模、多样性和真实性。相关代码、模型和数据将公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T17:02:27Z
- **目录日期**: 2025-06-02
