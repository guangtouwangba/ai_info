# JARVIS-VLA: Post-Training Large-Scale Vision Language Models to Play Visual Games with Keyboards and Mouse

**URL**: http://arxiv.org/abs/2503.16365v1

## 原始摘要

Recently, action-based decision-making in open-world environments has gained
significant attention. Visual Language Action (VLA) models, pretrained on
large-scale web datasets, have shown promise in decision-making tasks. However,
previous work has primarily focused on action post-training, often neglecting
enhancements to the foundational model itself. In response, we introduce a
novel approach, Act from Visual Language Post-Training, which refines Visual
Language Models (VLMs) through visual and linguistic guidance in a
self-supervised manner. This enhancement improves the models' capabilities in
world knowledge, visual recognition, and spatial grounding in open-world
environments. Following the above post-training paradigms, we obtain the first
VLA models in Minecraft that can follow human instructions on over 1k different
atomic tasks, including crafting, smelting, cooking, mining, and killing. Our
experiments demonstrate that post-training on non-trajectory tasks leads to a
significant 40% improvement over the best agent baseline on a diverse set of
atomic tasks. Furthermore, we demonstrate that our approach surpasses
traditional imitation learning-based policies in Minecraft, achieving
state-of-the-art performance. We have open-sourced the code, models, and
datasets to foster further research. The project page can be found in
https://craftjarvis.github.io/JarvisVLA.


## AI 摘要

近年来，开放世界环境中的基于行动的决策引起了广泛关注。预训练于大规模网络数据集的视觉语言行动（VLA）模型在决策任务中表现出潜力。然而，以往的研究主要集中在行动后训练，往往忽视了对基础模型的改进。为此，我们提出了一种新方法，通过视觉和语言指导以自监督方式改进视觉语言模型（VLMs），从而提升其在世界知识、视觉识别和空间定位方面的能力。实验表明，该方法在Minecraft中超越了传统模仿学习策略，实现了最先进的性能。代码、模型和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-21T23:02:03Z
- **目录日期**: 2025-03-21
