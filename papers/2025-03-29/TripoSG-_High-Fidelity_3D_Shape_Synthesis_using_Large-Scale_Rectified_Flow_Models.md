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

TripoSG是一种新型3D形状生成扩散模型，能根据输入图像生成高保真3D网格。该模型采用三大创新：1)基于大规模高质量数据训练的整流流transformer，实现最先进的生成质量；2)结合SDF、法线和eikonal损失的混合监督训练策略，提升3D重建性能；3)构建200万高质量3D样本的数据处理流程，验证数据质量与数量的重要性。实验表明，TripoSG在细节表现和输入图像匹配度上达到最优，并能适应多种图像风格，展现出色的泛化能力。该模型将公开以推动3D生成领域发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-29T23:02:02Z
- **目录日期**: 2025-03-29
