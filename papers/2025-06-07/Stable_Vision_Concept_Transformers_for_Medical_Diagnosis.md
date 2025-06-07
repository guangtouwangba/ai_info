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

本文提出两种可解释AI模型——视觉概念转换器(VCT)和稳定视觉概念转换器(SVCT)，以解决医学AI领域的透明性问题。VCT通过融合图像特征和概念特征提升模型性能，SVCT在此基础上结合去噪扩散平滑技术增强模型鲁棒性。实验表明，这两种模型在四个医学数据集上均保持高准确性和可解释性，SVCT尤其能在输入扰动情况下提供稳定的解释。相比传统概念瓶颈模型(CBMs)，新方法既保留了概念特征的可解释优势，又通过整合图像特征提升了模型性能，更符合医学领域对可靠解释的需求。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T19:01:53Z
- **目录日期**: 2025-06-07
