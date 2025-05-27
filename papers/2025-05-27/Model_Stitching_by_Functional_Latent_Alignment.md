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

本文提出了一种评估神经网络功能相似性的新方法——功能潜在对齐(FuLA)。该方法受知识蒸馏启发，将模型缝合的最优性条件重新定义为潜在空间的对齐，从而更可靠地衡量不同网络学习到的功能表征相似度。研究通过对抗训练、捷径学习和跨层缝合三个测试平台验证了FuLA的优势：相比传统缝合方法，FuLA能有效避免任务线索带来的伪影，同时捕捉到被缝合级匹配忽略的重要对齐关系。实验表明，FuLA在多种场景下都能提供更稳健的功能相似性评估，为AI模型表征分析提供了新工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T17:01:18Z
- **目录日期**: 2025-05-27
