# EgoCHARM: Resource-Efficient Hierarchical Activity Recognition using an Egocentric IMU Sensor

**URL**: http://arxiv.org/abs/2504.17735v1

## 原始摘要

Human activity recognition (HAR) on smartglasses has various use cases,
including health/fitness tracking and input for context-aware AI assistants.
However, current approaches for egocentric activity recognition suffer from low
performance or are resource-intensive. In this work, we introduce a resource
(memory, compute, power, sample) efficient machine learning algorithm,
EgoCHARM, for recognizing both high level and low level activities using a
single egocentric (head-mounted) Inertial Measurement Unit (IMU). Our
hierarchical algorithm employs a semi-supervised learning strategy, requiring
primarily high level activity labels for training, to learn generalizable low
level motion embeddings that can be effectively utilized for low level activity
recognition. We evaluate our method on 9 high level and 3 low level activities
achieving 0.826 and 0.855 F1 scores on high level and low level activity
recognition respectively, with just 63k high level and 22k low level model
parameters, allowing the low level encoder to be deployed directly on current
IMU chips with compute. Lastly, we present results and insights from a
sensitivity analysis and highlight the opportunities and limitations of
activity recognition using egocentric IMUs.


## AI 摘要

该研究提出了一种资源高效的机器学习算法EgoCHARM，用于智能眼镜上的单IMU设备进行人类活动识别。该分层算法采用半监督学习策略，主要使用高级活动标签训练，同时学习可泛化的低级运动嵌入特征，能同时识别高级和低级活动。在9种高级和3种低级活动测试中，分别获得0.826和0.855的F1分数，模型参数仅63k（高级）和22k（低级），可直接部署在当前IMU芯片上。研究还分析了IMU活动识别的机会与限制，为资源受限设备的实时活动识别提供了高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T06:01:12Z
- **目录日期**: 2025-04-25
