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

本文提出EgoCHARM算法，通过单个头戴式惯性测量单元(IMU)实现高效的人体活动识别。该分层算法采用半监督学习策略，主要利用高级活动标签进行训练，同时学习可泛化的低级运动嵌入特征。在9种高级和3种低级活动识别中分别取得0.826和0.855的F1分数，模型参数仅63k(高级)和22k(低级)，可直接部署在当前IMU芯片上。研究还分析了基于头戴式IMU的活动识别的机会与限制，为资源(内存、计算、功耗、样本)高效的智能眼镜活动识别提供了解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-27T12:01:21Z
- **目录日期**: 2025-04-27
