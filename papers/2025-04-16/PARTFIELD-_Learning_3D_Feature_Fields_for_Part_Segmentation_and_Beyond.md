# PARTFIELD: Learning 3D Feature Fields for Part Segmentation and Beyond

**URL**: http://arxiv.org/abs/2504.11451v1

## 原始摘要

We propose PartField, a feedforward approach for learning part-based 3D
features, which captures the general concept of parts and their hierarchy
without relying on predefined templates or text-based names, and can be applied
to open-world 3D shapes across various modalities. PartField requires only a 3D
feedforward pass at inference time, significantly improving runtime and
robustness compared to prior approaches. Our model is trained by distilling 2D
and 3D part proposals from a mix of labeled datasets and image segmentations on
large unsupervised datasets, via a contrastive learning formulation. It
produces a continuous feature field which can be clustered to yield a
hierarchical part decomposition. Comparisons show that PartField is up to 20%
more accurate and often orders of magnitude faster than other recent
class-agnostic part-segmentation methods. Beyond single-shape part
decomposition, consistency in the learned field emerges across shapes, enabling
tasks such as co-segmentation and correspondence, which we demonstrate in
several applications of these general-purpose, hierarchical, and consistent 3D
feature fields. Check our Webpage!
https://research.nvidia.com/labs/toronto-ai/partfield-release/


## AI 摘要

PartField是一种前馈式方法，用于学习基于部件的3D特征，无需依赖预定义模板或文本标签即可捕捉部件及其层次结构的通用概念，适用于多种模态的开放世界3D形状。该方法仅需单次3D前馈推理，显著提升了运行速度和鲁棒性。通过对比学习框架，模型从标注数据集和大规模无监督数据集的图像分割中蒸馏2D/3D部件提案，生成连续特征场，可聚类为层次化部件分解。实验表明，PartField在准确率上最高提升20%，且速度比其他无类别部件分割方法快数个数量级。此外，学习到的特征场在形状间具有一致性，支持协同分割和对应关系等任务。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T20:01:52Z
- **目录日期**: 2025-04-16
