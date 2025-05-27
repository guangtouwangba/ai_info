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

本文提出了一种评估神经网络功能相似性的新方法——功能潜在对齐（FuLA），通过借鉴知识蒸馏思想改进模型缝合范式。相比传统缝合方法，FuLA能更可靠地衡量独立训练网络的功能相似性，减少任务线索带来的干扰。实验在对抗训练、捷径训练和跨层缝合三个测试场景中验证了FuLA的优势：既能实现非平凡对齐，又避免了缝合层匹配的局限性。该研究为AI模型的功能比较提供了新工具，对理解神经网络表征学习具有重要意义。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T22:01:05Z
- **目录日期**: 2025-05-27
