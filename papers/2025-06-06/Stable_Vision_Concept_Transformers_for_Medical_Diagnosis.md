# Stable Vision Concept Transformers for Medical Diagnosis

**URL**: http://arxiv.org/abs/2506.05286v1

## 原始摘要

Transparency is a paramount concern in the medical field, prompting
researchers to delve into the realm of explainable AI (XAI). Among these XAI
methods, Concept Bottleneck Models (CBMs) aim to restrict the model's latent
space to human-understandable high-level concepts by generating a conceptual
layer for extracting conceptual features, which has drawn much attention
recently. However, existing methods rely solely on concept features to
determine the model's predictions, which overlook the intrinsic feature
embeddings within medical images. To address this utility gap between the
original models and concept-based models, we propose Vision Concept Transformer
(VCT). Furthermore, despite their benefits, CBMs have been found to negatively
impact model performance and fail to provide stable explanations when faced
with input perturbations, which limits their application in the medical field.
To address this faithfulness issue, this paper further proposes the Stable
Vision Concept Transformer (SVCT) based on VCT, which leverages the vision
transformer (ViT) as its backbone and incorporates a conceptual layer. SVCT
employs conceptual features to enhance decision-making capabilities by fusing
them with image features and ensures model faithfulness through the integration
of Denoised Diffusion Smoothing. Comprehensive experiments on four medical
datasets demonstrate that our VCT and SVCT maintain accuracy while remaining
interpretable compared to baselines. Furthermore, even when subjected to
perturbations, our SVCT model consistently provides faithful explanations, thus
meeting the needs of the medical field.


## AI 摘要

本文针对医疗领域对AI可解释性的需求，提出两种改进的视觉概念转换模型。针对现有概念瓶颈模型(CBMs)仅依赖概念特征而忽略医学图像内在特征的问题，作者首先提出视觉概念转换器(VCT)。为进一步解决CBM模型在输入扰动下解释不稳定的问题，基于VCT开发了稳定视觉概念转换器(SVCT)。SVCT采用视觉转换器(ViT)架构，融合概念特征与图像特征提升决策能力，并引入去噪扩散平滑技术确保解释稳定性。在四个医学数据集上的实验表明，VCT和SVCT在保持准确性的同时具有良好可解释性，SVCT还能在扰动下提供稳定解释，满足医疗领域需求。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T20:02:07Z
- **目录日期**: 2025-06-06
