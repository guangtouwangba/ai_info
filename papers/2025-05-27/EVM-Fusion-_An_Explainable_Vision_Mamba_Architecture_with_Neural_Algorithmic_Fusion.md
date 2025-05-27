# EVM-Fusion: An Explainable Vision Mamba Architecture with Neural Algorithmic Fusion

**URL**: http://arxiv.org/abs/2505.17367v2

## 原始摘要

Medical image classification is critical for clinical decision-making, yet
demands for accuracy, interpretability, and generalizability remain
challenging. This paper introduces EVM-Fusion, an Explainable Vision Mamba
architecture featuring a novel Neural Algorithmic Fusion (NAF) mechanism for
multi-organ medical image classification. EVM-Fusion leverages a multipath
design, where DenseNet and U-Net based pathways, enhanced by Vision Mamba (Vim)
modules, operate in parallel with a traditional feature pathway. These diverse
features are dynamically integrated via a two-stage fusion process: cross-modal
attention followed by the iterative NAF block, which learns an adaptive fusion
algorithm. Intrinsic explainability is embedded through path-specific spatial
attention, Vim {\Delta}-value maps, traditional feature SE-attention, and
cross-modal attention weights. Experiments on a diverse 9-class multi-organ
medical image dataset demonstrate EVM-Fusion's strong classification
performance, achieving 99.75% test accuracy and provide multi-faceted insights
into its decision-making process, highlighting its potential for trustworthy AI
in medical diagnostics.


## AI 摘要

本文提出EVM-Fusion，一种基于可解释视觉Mamba架构的医学图像分类方法。该方法采用多路径设计，结合DenseNet、U-Net和传统特征路径，通过Vision Mamba模块增强特征提取。创新性地使用两阶段融合机制：跨模态注意力与自适应神经算法融合(NAF)块动态整合多源特征。系统内置多种可解释性组件，包括路径特异性空间注意力、Vim Δ值图等。在9类多器官医学图像数据集上取得99.75%的测试准确率，同时提供多维决策依据，展现了其在可信医疗AI诊断中的应用潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T20:02:18Z
- **目录日期**: 2025-05-27
