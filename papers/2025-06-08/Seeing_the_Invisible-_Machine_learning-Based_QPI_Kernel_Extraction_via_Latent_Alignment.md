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

本研究提出了首个基于人工智能的准粒子干涉(QPI)核提取框架。针对多散射源图像中提取单散射QPI核这一病态逆问题，研究团队开发了一种两步学习策略：首先通过变分自编码器学习散射核的紧凑潜在空间表示，然后使用专用编码器将QPI观测的潜在表示与预学习核对齐。这种设计使模型即使在复杂纠缠散射条件下也能稳健推断核函数。研究构建了包含100个独特核的多样化QPI数据集，实验表明该方法相比直接一步法基线具有显著更高的提取精度，并对未见核表现出更好的泛化能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-08T12:01:29Z
- **目录日期**: 2025-06-08
