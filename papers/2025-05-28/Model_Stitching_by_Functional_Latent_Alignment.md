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

本文提出了一种评估神经网络功能相似性的新方法——功能潜在对齐（FuLA），通过借鉴知识蒸馏思想改进模型拼接范式。FuLA作为模型拼接的最优条件，能更可靠地衡量独立训练网络的功能相似性。实验在对抗训练、捷径训练和跨层拼接三个场景中验证了FuLA的优势：相比传统拼接方法，FuLA能有效避免任务线索相关的伪影，同时捕捉到被拼接层匹配忽略的重要对齐关系。研究还构建了新的功能相似性测试基准，证实FuLA在多种场景下都具有更优的评估性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T01:28:48Z
- **目录日期**: 2025-05-28
