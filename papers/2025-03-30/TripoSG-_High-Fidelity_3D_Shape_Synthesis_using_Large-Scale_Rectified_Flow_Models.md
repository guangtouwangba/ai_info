# TripoSG: High-Fidelity 3D Shape Synthesis using Large-Scale Rectified Flow Models

**URL**: http://arxiv.org/abs/2502.06608v3

## 原始摘要

Recent advancements in diffusion techniques have propelled image and video
generation to unprecedented levels of quality, significantly accelerating the
deployment and application of generative AI. However, 3D shape generation
technology has so far lagged behind, constrained by limitations in 3D data
scale, complexity of 3D data processing, and insufficient exploration of
advanced techniques in the 3D domain. Current approaches to 3D shape generation
face substantial challenges in terms of output quality, generalization
capability, and alignment with input conditions. We present TripoSG, a new
streamlined shape diffusion paradigm capable of generating high-fidelity 3D
meshes with precise correspondence to input images. Specifically, we propose:
1) A large-scale rectified flow transformer for 3D shape generation, achieving
state-of-the-art fidelity through training on extensive, high-quality data. 2)
A hybrid supervised training strategy combining SDF, normal, and eikonal losses
for 3D VAE, achieving high-quality 3D reconstruction performance. 3) A data
processing pipeline to generate 2 million high-quality 3D samples, highlighting
the crucial rules for data quality and quantity in training 3D generative
models. Through comprehensive experiments, we have validated the effectiveness
of each component in our new framework. The seamless integration of these parts
has enabled TripoSG to achieve state-of-the-art performance in 3D shape
generation. The resulting 3D shapes exhibit enhanced detail due to
high-resolution capabilities and demonstrate exceptional fidelity to input
images. Moreover, TripoSG demonstrates improved versatility in generating 3D
models from diverse image styles and contents, showcasing strong generalization
capabilities. To foster progress and innovation in the field of 3D generation,
we will make our model publicly available.


## AI 摘要

TripoSG是一种新型3D形状扩散生成框架，通过大规模整流流变换器和混合监督训练策略（结合SDF、法线和Eikonal损失），实现了高保真3D网格生成。该技术解决了当前3D生成在质量、泛化能力和输入对齐方面的挑战，利用200万高质量3D样本训练，支持从多样图像风格生成细节丰富的3D模型。实验验证了各组件有效性，其高分辨率能力显著提升了输出细节与输入图像的匹配精度。该模型将开源以推动3D生成领域发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-30T10:01:56Z
- **目录日期**: 2025-03-30
