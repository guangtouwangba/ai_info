# UniSync: A Unified Framework for Audio-Visual Synchronization

**URL**: http://arxiv.org/abs/2503.16357v1

## 原始摘要

Precise audio-visual synchronization in speech videos is crucial for content
quality and viewer comprehension. Existing methods have made significant
strides in addressing this challenge through rule-based approaches and
end-to-end learning techniques. However, these methods often rely on limited
audio-visual representations and suboptimal learning strategies, potentially
constraining their effectiveness in more complex scenarios. To address these
limitations, we present UniSync, a novel approach for evaluating audio-visual
synchronization using embedding similarities. UniSync offers broad
compatibility with various audio representations (e.g., Mel spectrograms,
HuBERT) and visual representations (e.g., RGB images, face parsing maps, facial
landmarks, 3DMM), effectively handling their significant dimensional
differences. We enhance the contrastive learning framework with a margin-based
loss component and cross-speaker unsynchronized pairs, improving discriminative
capabilities. UniSync outperforms existing methods on standard datasets and
demonstrates versatility across diverse audio-visual representations. Its
integration into talking face generation frameworks enhances synchronization
quality in both natural and AI-generated content.


## AI 摘要

UniSync是一种新颖的音频-视频同步评估方法，通过嵌入相似性来解决现有方法在复杂场景中的局限性。它兼容多种音频（如梅尔频谱、HuBERT）和视觉表示（如RGB图像、面部解析图、面部标志、3DMM），并有效处理其维度差异。通过引入基于边界的损失组件和跨说话者不同步对，UniSync增强了对比学习框架的判别能力。在标准数据集上，UniSync优于现有方法，并在自然和AI生成内容中提高了同步质量，展示了其在不同音频-视频表示中的多功能性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-22T23:02:13Z
- **目录日期**: 2025-03-22
