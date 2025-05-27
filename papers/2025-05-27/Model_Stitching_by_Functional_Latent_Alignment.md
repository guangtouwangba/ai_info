# Model Stitching by Functional Latent Alignment

**URL**: http://arxiv.org/abs/2505.20142v1

## 原始摘要

Evaluating functional similarity involves quantifying the degree to which
independently trained neural networks learn functionally similar
representations. Reliably inferring the functional similarity of these networks
remains an open problem with far-reaching implications for AI. Model stitching
has emerged as a promising paradigm, where an optimal affine transformation
aligns two models to solve a task, with the stitched model serving as a proxy
for functional similarity. In this work, we draw inspiration from the knowledge
distillation literature and propose Functional Latent Alignment (FuLA) as a
novel optimality condition for model stitching. We revisit previously explored
functional similarity testbeds and introduce a new one, based on which FuLA
emerges as an overall more reliable method of functional similarity.
Specifically, our experiments in (a) adversarial training, (b) shortcut
training and, (c) cross-layer stitching, reveal that FuLA is less prone to
artifacts tied to training on task cues while achieving non-trivial alignments
that are missed by stitch-level matching.


## AI 摘要

本文提出了一种评估神经网络功能相似性的新方法——功能潜在对齐(FuLA)，该方法受知识蒸馏启发，通过最优仿射变换对齐两个模型来完成任务。相比传统模型拼接方法，FuLA在对抗训练、捷径训练和跨层拼接等测试场景中表现出更强的可靠性，能减少任务线索带来的干扰，捕捉传统方法遗漏的重要对齐关系。研究重新评估了现有测试基准并引入新基准，证实FuLA在量化独立训练网络的功能相似性方面更具优势，为AI系统功能相似性评估提供了更可靠的方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T11:01:05Z
- **目录日期**: 2025-05-27
