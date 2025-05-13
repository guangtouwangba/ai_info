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

本文研究了医疗领域的文本到图像生成技术，比较了两种方法：(1)微调大型预训练模型(FLUX、Kandinsky)和(2)训练小型专用模型(MSDM)。作者提出了一种基于Stable Diffusion优化的MSDM模型，整合临床文本编码器、变分自编码器和交叉注意力机制，以更好地对齐医学文本与生成图像。在结肠镜(MedVQA-GI)和放射学(ROCOv2)数据集上的评估显示，大型模型图像保真度更高，但MSDM能以更低计算成本实现相近质量。定量指标和专家评估揭示了各方法的优缺点，为解决医疗AI数据稀缺问题提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T09:02:26Z
- **目录日期**: 2025-05-13
