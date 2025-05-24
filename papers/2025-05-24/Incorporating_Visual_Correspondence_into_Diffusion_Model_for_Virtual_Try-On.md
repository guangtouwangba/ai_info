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

本文提出了一种基于扩散模型的虚拟试衣(VTON)新方法SPM-Diff，通过显式利用视觉对应关系作为先验来改进细节保留。该方法将服装细节表示为结构化语义点，通过局部流变形匹配目标人体的语义点，并利用深度/法线图增强为3D感知线索。这种对应关系模拟了穿衣过程，3D线索作为监督信号指导扩散模型训练。此外还设计了点聚焦扩散损失以充分利用语义点匹配。在VITON-HD和DressCode数据集上的实验表明，该方法在服装细节保留方面表现优异，达到了最先进的VTON性能。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-24T08:02:29Z
- **目录日期**: 2025-05-24
