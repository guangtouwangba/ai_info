# AnimeShooter: A Multi-Shot Animation Dataset for Reference-Guided Video Generation

**URL**: http://arxiv.org/abs/2506.03126v1

## 原始摘要

Recent advances in AI-generated content (AIGC) have significantly accelerated
animation production. To produce engaging animations, it is essential to
generate coherent multi-shot video clips with narrative scripts and character
references. However, existing public datasets primarily focus on real-world
scenarios with global descriptions, and lack reference images for consistent
character guidance. To bridge this gap, we present AnimeShooter, a
reference-guided multi-shot animation dataset. AnimeShooter features
comprehensive hierarchical annotations and strong visual consistency across
shots through an automated pipeline. Story-level annotations provide an
overview of the narrative, including the storyline, key scenes, and main
character profiles with reference images, while shot-level annotations
decompose the story into consecutive shots, each annotated with scene,
characters, and both narrative and descriptive visual captions. Additionally, a
dedicated subset, AnimeShooter-audio, offers synchronized audio tracks for each
shot, along with audio descriptions and sound sources. To demonstrate the
effectiveness of AnimeShooter and establish a baseline for the reference-guided
multi-shot video generation task, we introduce AnimeShooterGen, which leverages
Multimodal Large Language Models (MLLMs) and video diffusion models. The
reference image and previously generated shots are first processed by MLLM to
produce representations aware of both reference and context, which are then
used as the condition for the diffusion model to decode the subsequent shot.
Experimental results show that the model trained on AnimeShooter achieves
superior cross-shot visual consistency and adherence to reference visual
guidance, which highlight the value of our dataset for coherent animated video
generation.


## AI 摘要

该研究提出了AnimeShooter数据集，用于解决AI生成动画中多镜头连贯性和角色一致性问题。该数据集包含分层标注（故事级概述和镜头级分解）和视觉一致的参考图像，并配有同步音频的子集AnimeShooter-audio。研究者还开发了基线模型AnimeShooterGen，结合多模态大语言模型和视频扩散模型，利用参考图像和上下文生成连贯后续镜头。实验表明，基于该数据集训练的模型在跨镜头视觉一致性和参考遵循方面表现优异，为连贯动画视频生成提供了有效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T14:01:40Z
- **目录日期**: 2025-06-04
