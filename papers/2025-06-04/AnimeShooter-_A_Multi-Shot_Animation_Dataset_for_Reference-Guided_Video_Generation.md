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

研究人员提出了AnimeShooter数据集，用于解决AI生成动画中多镜头连贯性和角色一致性的问题。该数据集包含分层注释（故事级和镜头级）和视觉一致的参考图像，并通过自动化流程确保质量。此外，AnimeShooter-audio子集还提供同步音频轨道。为验证数据集效果，团队开发了AnimeShooterGen基线模型，结合多模态大语言模型和视频扩散模型，利用参考图像和先前生成的镜头来保持视觉一致性。实验表明，基于该数据集训练的模型在跨镜头一致性和参考视觉引导方面表现优异，为连贯动画视频生成提供了有力支持。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T20:01:22Z
- **目录日期**: 2025-06-04
