# A Lightweight Dual-Branch System for Weakly-Supervised Video Anomaly Detection on Consumer Edge Devices

**URL**: http://arxiv.org/abs/2410.21991v7

## 原始摘要

The growing demand for intelligent security in consumer electronics, such as
smart home cameras and personal monitoring systems, is often hindered by the
high computational cost and large model sizes of advanced AI. These limitations
prevent the effective deployment of real-time Video Anomaly Detection (VAD) on
resource-constrained edge devices. To bridge this gap, this paper introduces
Rule-based Video Anomaly Detection (RuleVAD), a novel, lightweight system
engineered for high-efficiency and low-complexity threat detection directly on
consumer hardware. RuleVAD features an innovative decoupled dual-branch
architecture to minimize computational load. An implicit branch uses visual
features for rapid, coarse-grained binary classification, efficiently filtering
out normal activity to avoid unnecessary processing. For potentially anomalous
or complex events, a multimodal explicit branch takes over. This branch
leverages YOLO-World to detect objects and applies data mining to generate
interpretable, text-based association rules from the scene. By aligning these
rules with visual data, RuleVAD achieves a more nuanced, fine-grained
classification, significantly reducing the false alarms common in vision-only
systems. Extensive experiments on the XD-Violence and UCF-Crime benchmark
datasets show that RuleVAD achieves superior performance, surpassing existing
state-of-the-art methods in both accuracy and speed. Crucially, the entire
system is optimized for low-power operation and is fully deployable on an
NVIDIA Jetson Nano board, demonstrating its practical feasibility for bringing
advanced, real-time security monitoring to everyday consumer electronic
devices.


## AI 摘要

本文提出了一种轻量级视频异常检测系统RuleVAD，专为消费电子设备设计。该系统采用解耦双分支架构：隐式分支通过视觉特征快速过滤正常活动；显式分支结合YOLO-World目标检测和数据挖掘生成可解释的文本关联规则，实现细粒度分类。实验表明，RuleVAD在XD-Violence和UCF-Crime数据集上性能优于现有方法，且能高效部署于NVIDIA Jetson Nano等边缘设备，显著降低了误报率，为智能家居等场景提供了实时、低功耗的安全监控解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T05:02:53Z
- **目录日期**: 2025-06-09
