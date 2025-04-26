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

该研究提出了一种名为EgoCHARM的资源高效机器学习算法，用于通过头戴式惯性测量单元(IMU)识别高级和低级活动。该分层算法采用半监督学习策略，主要依赖高级活动标签进行训练，同时学习可泛化的低级运动嵌入特征。实验在9种高级和3种低级活动上分别获得0.826和0.855的F1分数，模型仅需63k(高级)和22k(低级)参数，可直接部署在当前IMU芯片上。研究还分析了使用头戴式IMU进行活动识别的机会与限制，为智能眼镜上的活动识别提供了高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-26T17:01:06Z
- **目录日期**: 2025-04-26
