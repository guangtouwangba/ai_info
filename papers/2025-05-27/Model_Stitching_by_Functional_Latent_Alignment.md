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

评估神经网络功能相似性是一个重要但尚未解决的AI问题。模型缝合（model stitching）是目前的主流方法，通过最优仿射变换对齐两个模型来评估相似性。本文受知识蒸馏启发，提出功能潜在对齐（FuLA）作为新的缝合最优条件。实验表明，FuLA在对抗训练、捷径学习和跨层缝合等场景中，能更可靠地评估功能相似性，减少任务线索带来的干扰，并识别出传统缝合方法遗漏的重要对齐关系。相比现有方法，FuLA展现出更优的综合性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T03:19:06Z
- **目录日期**: 2025-05-27
