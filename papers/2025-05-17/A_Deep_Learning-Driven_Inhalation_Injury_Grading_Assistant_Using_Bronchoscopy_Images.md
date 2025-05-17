# A Deep Learning-Driven Inhalation Injury Grading Assistant Using Bronchoscopy Images

**URL**: http://arxiv.org/abs/2505.08517v2

## 原始摘要

Inhalation injuries present a challenge in clinical diagnosis and grading due
to Conventional grading methods such as the Abbreviated Injury Score (AIS)
being subjective and lacking robust correlation with clinical parameters like
mechanical ventilation duration and patient mortality. This study introduces a
novel deep learning-based diagnosis assistant tool for grading inhalation
injuries using bronchoscopy images to overcome subjective variability and
enhance consistency in severity assessment. Our approach leverages data
augmentation techniques, including graphic transformations, Contrastive
Unpaired Translation (CUT), and CycleGAN, to address the scarcity of medical
imaging data. We evaluate the classification performance of two deep learning
models, GoogLeNet and Vision Transformer (ViT), across a dataset significantly
expanded through these augmentation methods. The results demonstrate GoogLeNet
combined with CUT as the most effective configuration for grading inhalation
injuries through bronchoscopy images and achieves a classification accuracy of
97.8%. The histograms and frequency analysis evaluations reveal variations
caused by the augmentation CUT with distribution changes in the histogram and
texture details of the frequency spectrum. PCA visualizations underscore the
CUT substantially enhances class separability in the feature space. Moreover,
Grad-CAM analyses provide insight into the decision-making process; mean
intensity for CUT heatmaps is 119.6, which significantly exceeds 98.8 of the
original datasets. Our proposed tool leverages mechanical ventilation periods
as a novel grading standard, providing comprehensive diagnostic support.


## AI 摘要

本研究开发了一种基于深度学习的吸入性损伤分级辅助工具，利用支气管镜图像克服传统主观评分方法的局限性。通过数据增强技术（包括图像变换、CUT和CycleGAN）解决医学影像数据稀缺问题，评估了GoogLeNet和ViT模型的分类性能。结果显示，GoogLeNet结合CUT增强效果最佳，分类准确率达97.8%。CUT显著改善了特征空间中的类别区分度（Grad-CAM热图平均强度119.6 vs 原始数据98.8），并通过机械通气时长这一客观指标建立新型分级标准。该工具提升了吸入性损伤评估的客观性和一致性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-17T11:02:10Z
- **目录日期**: 2025-05-17
