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

本文提出了一种评估神经网络功能相似性的新方法——功能潜在对齐(FuLA)，该方法受到知识蒸馏的启发，通过最优仿射变换对齐不同模型来解决任务。相比传统模型缝合方法，FuLA在对抗训练、捷径训练和跨层缝合等测试场景中表现出更高的可靠性，能有效减少任务线索带来的干扰，并捕捉到缝合级匹配所忽略的重要对齐关系。实验表明，FuLA在多种功能相似性测试平台上均优于现有方法，为解决AI领域长期存在的神经网络功能相似性评估问题提供了更可靠的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T21:01:05Z
- **目录日期**: 2025-05-27
