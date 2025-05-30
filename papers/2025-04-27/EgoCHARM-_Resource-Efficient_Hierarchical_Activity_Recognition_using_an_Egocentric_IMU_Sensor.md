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

本文提出EgoCHARM算法，用于智能眼镜上的单IMU设备进行高效人体活动识别。该分层算法采用半监督学习策略，主要利用高级活动标签训练，同时学习可泛化的低级运动嵌入特征，实现高/低级活动识别。在9种高级和3种低级活动测试中，F1分数分别达0.826和0.855，模型参数量仅63k（高级）和22k（低级），可直接部署于现有IMU芯片。研究还通过敏感性分析揭示了头戴式IMU活动识别的优势与局限，为健康监测和情境感知AI提供了轻量级解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-27T23:01:23Z
- **目录日期**: 2025-04-27
