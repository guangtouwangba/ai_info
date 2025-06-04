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

DyTact是一种无标记的动态手-物体接触捕捉方法，用于精确重建手部操作中的接触动态。该方法采用基于2D高斯面元的动态关节表示，结合MANO网格的模板先验来优化稳定性和效率。通过细化模块处理高频变形，并采用接触引导的自适应采样策略提升遮挡区域的精度。实验表明，DyTact在动态接触估计精度和新视角合成质量上达到领先水平，同时具有快速优化和高效内存使用的优势。项目页面：https://oliver-cong02.github.io/DyTact.github.io/。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T14:02:03Z
- **目录日期**: 2025-06-04
