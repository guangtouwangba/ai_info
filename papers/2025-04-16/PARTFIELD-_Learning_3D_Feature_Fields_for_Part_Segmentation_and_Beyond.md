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

PartField是一种基于部件学习的3D特征提取方法，无需依赖预定义模板或文本标签即可自动识别部件及其层次结构，适用于多种模态的开放世界3D形状。该方法仅需单次前馈推理，显著提升了运行速度和鲁棒性。通过对比学习框架，结合标注数据集和大规模无监督数据集的2D/3D部件提案进行训练，生成连续特征场，可聚类得到层次化部件分解。实验表明，PartField的准确率最高提升20%，且速度比其他无类别部件分割方法快数个数量级。其一致性特征场还支持跨形状的协同分割与对应任务。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T16:02:12Z
- **目录日期**: 2025-04-16
