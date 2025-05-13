# Prompt to Polyp: Medical Text-Conditioned Image Synthesis with Diffusion Models

**URL**: http://arxiv.org/abs/2505.05573v2

## 原始摘要

The generation of realistic medical images from text descriptions has
significant potential to address data scarcity challenges in healthcare AI
while preserving patient privacy. This paper presents a comprehensive study of
text-to-image synthesis in the medical domain, comparing two distinct
approaches: (1) fine-tuning large pre-trained latent diffusion models and (2)
training small, domain-specific models. We introduce a novel model named MSDM,
an optimized architecture based on Stable Diffusion that integrates a clinical
text encoder, variational autoencoder, and cross-attention mechanisms to better
align medical text prompts with generated images. Our study compares two
approaches: fine-tuning large pre-trained models (FLUX, Kandinsky) versus
training compact domain-specific models (MSDM). Evaluation across colonoscopy
(MedVQA-GI) and radiology (ROCOv2) datasets reveals that while large models
achieve higher fidelity, our optimized MSDM delivers comparable quality with
lower computational costs. Quantitative metrics and qualitative evaluations by
medical experts reveal strengths and limitations of each approach.


## AI 摘要

本研究比较了两种医学文本到图像生成方法：微调大型预训练模型（如FLUX、Kandinsky）与训练小型专用模型（MSDM）。提出的MSDM模型基于Stable Diffusion优化架构，整合临床文本编码器、变分自编码器和交叉注意力机制，能更好对齐医学文本与生成图像。在结肠镜（MedVQA-GI）和放射学（ROCOv2）数据集上的评估显示，大型模型生成质量更高，但MSDM能以更低计算成本达到相近效果。定量指标和医学专家评估揭示了两种方法的优劣，为解决医疗AI数据稀缺问题提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T03:20:30Z
- **目录日期**: 2025-05-13
