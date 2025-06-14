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

本文针对医疗领域对可解释AI的需求，提出了Vision Concept Transformer (VCT)和其改进版Stable Vision Concept Transformer (SVCT)。现有概念瓶颈模型(CBMs)仅依赖概念特征而忽略医学图像内在特征，且对输入扰动敏感。VCT通过融合概念特征和图像特征提升性能；SVCT基于VCT架构，结合视觉Transformer和去噪扩散平滑技术，确保模型鲁棒性和解释稳定性。在四个医学数据集上的实验表明，VCT和SVCT在保持准确性的同时具有可解释性。SVCT尤其能在扰动下提供稳定解释，满足医疗领域对可靠AI的需求。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T23:02:08Z
- **目录日期**: 2025-06-06
