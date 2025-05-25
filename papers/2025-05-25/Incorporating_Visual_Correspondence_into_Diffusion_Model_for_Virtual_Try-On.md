# Incorporating Visual Correspondence into Diffusion Model for Virtual Try-On

**URL**: http://arxiv.org/abs/2505.16977v1

## 原始摘要

Diffusion models have shown preliminary success in virtual try-on (VTON)
task. The typical dual-branch architecture comprises two UNets for implicit
garment deformation and synthesized image generation respectively, and has
emerged as the recipe for VTON task. Nevertheless, the problem remains
challenging to preserve the shape and every detail of the given garment due to
the intrinsic stochasticity of diffusion model. To alleviate this issue, we
novelly propose to explicitly capitalize on visual correspondence as the prior
to tame diffusion process instead of simply feeding the whole garment into UNet
as the appearance reference. Specifically, we interpret the fine-grained
appearance and texture details as a set of structured semantic points, and
match the semantic points rooted in garment to the ones over target person
through local flow warping. Such 2D points are then augmented into 3D-aware
cues with depth/normal map of target person. The correspondence mimics the way
of putting clothing on human body and the 3D-aware cues act as semantic point
matching to supervise diffusion model training. A point-focused diffusion loss
is further devised to fully take the advantage of semantic point matching.
Extensive experiments demonstrate strong garment detail preservation of our
approach, evidenced by state-of-the-art VTON performances on both VITON-HD and
DressCode datasets. Code is publicly available at:
https://github.com/HiDream-ai/SPM-Diff.


## AI 摘要

本文提出了一种基于扩散模型和语义点匹配(SPM)的虚拟试衣(VTON)新方法。针对现有双分支UNet架构在保留服装细节方面的不足，该方法创新性地利用视觉对应关系作为先验，通过将服装纹理细节表示为结构化语义点，并将其与目标人物的语义点进行局部流形匹配，同时结合深度/法线图增强3D感知线索。此外，还设计了专注于点的扩散损失函数来充分利用语义点匹配。在VITON-HD和DressCode数据集上的实验表明，该方法能有效保持服装细节，达到了最先进的性能表现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-25T11:02:34Z
- **目录日期**: 2025-05-25
