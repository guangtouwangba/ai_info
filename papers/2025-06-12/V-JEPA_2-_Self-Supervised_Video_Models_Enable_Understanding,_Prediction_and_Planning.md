# V-JEPA 2: Self-Supervised Video Models Enable Understanding, Prediction and Planning

**URL**: http://arxiv.org/abs/2506.09985v1

## 原始摘要

A major challenge for modern AI is to learn to understand the world and learn
to act largely by observation. This paper explores a self-supervised approach
that combines internet-scale video data with a small amount of interaction data
(robot trajectories), to develop models capable of understanding, predicting,
and planning in the physical world. We first pre-train an action-free
joint-embedding-predictive architecture, V-JEPA 2, on a video and image dataset
comprising over 1 million hours of internet video. V-JEPA 2 achieves strong
performance on motion understanding (77.3 top-1 accuracy on Something-Something
v2) and state-of-the-art performance on human action anticipation (39.7
recall-at-5 on Epic-Kitchens-100) surpassing previous task-specific models.
Additionally, after aligning V-JEPA 2 with a large language model, we
demonstrate state-of-the-art performance on multiple video question-answering
tasks at the 8 billion parameter scale (e.g., 84.0 on PerceptionTest, 76.9 on
TempCompass). Finally, we show how self-supervised learning can be applied to
robotic planning tasks by post-training a latent action-conditioned world
model, V-JEPA 2-AC, using less than 62 hours of unlabeled robot videos from the
Droid dataset. We deploy V-JEPA 2-AC zero-shot on Franka arms in two different
labs and enable picking and placing of objects using planning with image goals.
Notably, this is achieved without collecting any data from the robots in these
environments, and without any task-specific training or reward. This work
demonstrates how self-supervised learning from web-scale data and a small
amount of robot interaction data can yield a world model capable of planning in
the physical world.


## AI 摘要

本文提出了一种结合互联网视频数据与少量机器人交互数据的自监督学习方法V-JEPA 2。该模型首先在超100万小时视频数据上预训练，在动作理解（Something-Something v2准确率77.3）和人类动作预测（Epic-Kitchens-100召回率39.7）达到领先水平。与语言模型结合后，在视频问答任务中表现优异（PerceptionTest得分84.0）。进一步开发的V-JEPA 2-AC模型仅用62小时未标注机器人视频数据，就能在Franka机械臂上实现零样本物体抓取任务，无需环境特定数据或任务训练。研究表明，网络规模自监督学习可构建能进行物理世界规划的通用世界模型。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T14:01:28Z
- **目录日期**: 2025-06-12
