# RT-GAN: Recurrent Temporal GAN for Adding Lightweight Temporal Consistency to Frame-Based Domain Translation Approaches

**URL**: http://arxiv.org/abs/2310.00868v2

## 原始摘要

Fourteen million colonoscopies are performed annually just in the U.S.
However, the videos from these colonoscopies are not saved due to storage
constraints (each video from a high-definition colonoscope camera can be in
tens of gigabytes). Instead, a few relevant individual frames are saved for
documentation/reporting purposes and these are the frames on which most current
colonoscopy AI models are trained on. While developing new unsupervised domain
translation methods for colonoscopy (e.g. to translate between real optical and
virtual/CT colonoscopy), it is thus typical to start with approaches that
initially work for individual frames without temporal consistency. Once an
individual-frame model has been finalized, additional contiguous frames are
added with a modified deep learning architecture to train a new model from
scratch for temporal consistency. This transition to temporally-consistent deep
learning models, however, requires significantly more computational and memory
resources for training. In this paper, we present a lightweight solution with a
tunable temporal parameter, RT-GAN (Recurrent Temporal GAN), for adding
temporal consistency to individual frame-based approaches that reduces training
requirements by a factor of 5. We demonstrate the effectiveness of our approach
on two challenging use cases in colonoscopy: haustral fold segmentation
(indicative of missed surface) and realistic colonoscopy simulator video
generation. We also release a first-of-its kind temporal dataset for
colonoscopy for the above use cases. The datasets, accompanying code, and
pretrained models will be made available on our Computational Endoscopy
Platform GitHub (https://github.com/nadeemlab/CEP). The supplementary video is
available at https://youtu.be/UMVP-uIXwWk.


## AI 摘要

美国每年进行1400万例结肠镜检查，但高清视频因存储限制无法保存，仅保留少量关键帧用于报告和AI训练。现有方法主要针对单帧图像，缺乏时间一致性。本文提出轻量级RT-GAN模型，通过可调时间参数将时间一致性引入单帧方法，训练资源需求降低80%。该方案在结肠皱襞分割（漏检区域指示）和模拟视频生成两个应用场景验证有效，并发布了首个结肠镜时序数据集。代码、预训练模型和数据集已在GitHub开源（https://github.com/nadeemlab/CEP）。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T04:03:12Z
- **目录日期**: 2025-05-14
