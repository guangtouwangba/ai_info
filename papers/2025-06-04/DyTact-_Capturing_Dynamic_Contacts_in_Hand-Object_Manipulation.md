# DyTact: Capturing Dynamic Contacts in Hand-Object Manipulation

**URL**: http://arxiv.org/abs/2506.03103v1

## 原始摘要

Reconstructing dynamic hand-object contacts is essential for realistic
manipulation in AI character animation, XR, and robotics, yet it remains
challenging due to heavy occlusions, complex surface details, and limitations
in existing capture techniques. In this paper, we introduce DyTact, a
markerless capture method for accurately capturing dynamic contact in
hand-object manipulations in a non-intrusive manner. Our approach leverages a
dynamic, articulated representation based on 2D Gaussian surfels to model
complex manipulations. By binding these surfels to MANO meshes, DyTact
harnesses the inductive bias of template models to stabilize and accelerate
optimization. A refinement module addresses time-dependent high-frequency
deformations, while a contact-guided adaptive sampling strategy selectively
increases surfel density in contact regions to handle heavy occlusion.
Extensive experiments demonstrate that DyTact not only achieves
state-of-the-art dynamic contact estimation accuracy but also significantly
improves novel view synthesis quality, all while operating with fast
optimization and efficient memory usage. Project Page:
https://oliver-cong02.github.io/DyTact.github.io/ .


## AI 摘要

DyTact是一种新型无标记捕捉方法，用于精确重建动态手-物接触交互。该方法采用基于2D高斯面片的动态铰接式表征，通过绑定到MANO网格利用模板模型的归纳偏置来优化过程。其创新点包括：高频形变细化模块处理时序变形，接触引导的自适应采样策略提升接触区域的面片密度以应对严重遮挡。实验表明，DyTact不仅实现了最先进的动态接触估计精度，还显著提升了新视角合成质量，同时具有快速优化和高效内存使用的优势。项目页面：https://oliver-cong02.github.io/DyTact.github.io/。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T19:01:42Z
- **目录日期**: 2025-06-04
