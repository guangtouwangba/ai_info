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

DyTact是一种无标记的动态手-物体接触捕捉方法，通过基于2D高斯面元的动态铰接表示来建模复杂操作。该方法将面元绑定到MANO网格上，利用模板模型的归纳偏置来稳定和加速优化。其细化模块处理高频变形，而接触引导的自适应采样策略在接触区域选择性增加面元密度以应对严重遮挡。实验表明，DyTact不仅实现了最先进的动态接触估计精度，还显著提升了新视角合成质量，同时优化速度快且内存效率高。项目页面：https://oliver-cong02.github.io/DyTact.github.io/。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T17:01:48Z
- **目录日期**: 2025-06-04
