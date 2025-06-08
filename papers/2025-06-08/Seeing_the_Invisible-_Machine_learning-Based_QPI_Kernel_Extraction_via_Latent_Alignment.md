# Seeing the Invisible: Machine learning-Based QPI Kernel Extraction via Latent Alignment

**URL**: http://arxiv.org/abs/2506.05325v1

## 原始摘要

Quasiparticle interference (QPI) imaging is a powerful tool for probing
electronic structures in quantum materials, but extracting the single-scatterer
QPI pattern (i.e., the kernel) from a multi-scatterer image remains a
fundamentally ill-posed inverse problem. In this work, we propose the first
AI-based framework for QPI kernel extraction. We introduce a two-step learning
strategy that decouples kernel representation learning from
observation-to-kernel inference. In the first step, we train a variational
autoencoder to learn a compact latent space of scattering kernels. In the
second step, we align the latent representation of QPI observations with those
of the pre-learned kernels using a dedicated encoder. This design enables the
model to infer kernels robustly even under complex, entangled scattering
conditions. We construct a diverse and physically realistic QPI dataset
comprising 100 unique kernels and evaluate our method against a direct one-step
baseline. Experimental results demonstrate that our approach achieves
significantly higher extraction accuracy, and improved generalization to unseen
kernels.


## AI 摘要

本文提出了一种基于人工智能的准粒子干涉(QPI)核提取新方法。针对多散射体图像中提取单散射QPI模式这一病态逆问题，研究团队开发了包含两个学习步骤的框架：首先通过变分自编码器学习散射核的紧凑潜在空间表示，随后使用专用编码器将QPI观测数据与预学习核的潜在表示对齐。这种设计使模型能在复杂纠缠散射条件下稳健推断核函数。在包含100种独特核的多样化QPI数据集上测试表明，该方法相比直接单步基线具有显著更高的提取精度，并对未见核具有更好的泛化能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-08T03:27:19Z
- **目录日期**: 2025-06-08
