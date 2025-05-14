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

每年美国进行1400万例结肠镜检查，但高清视频因存储限制（单视频可达数十GB）通常不被保存，仅保留少量关键帧用于报告和AI训练。现有结肠镜AI模型多基于单帧数据，开发新方法（如光学与虚拟结肠镜转换）时也常从单帧模型入手，再扩展至时序一致性模型，但这显著增加计算资源需求。本文提出轻量级RT-GAN（循环时序生成对抗网络），通过可调时序参数将单帧模型升级为时序一致性模型，训练需求降低80%。方法在结肠皱襞分割和模拟视频生成任务中验证有效，并开源首个结肠镜时序数据集和代码。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T12:02:46Z
- **目录日期**: 2025-05-14
