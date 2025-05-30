# RedMotion: Motion Prediction via Redundancy Reduction

**URL**: http://arxiv.org/abs/2306.10840v4

## 原始摘要

We introduce RedMotion, a transformer model for motion prediction in
self-driving vehicles that learns environment representations via redundancy
reduction. Our first type of redundancy reduction is induced by an internal
transformer decoder and reduces a variable-sized set of local road environment
tokens, representing road graphs and agent data, to a fixed-sized global
embedding. The second type of redundancy reduction is obtained by
self-supervised learning and applies the redundancy reduction principle to
embeddings generated from augmented views of road environments. Our experiments
reveal that our representation learning approach outperforms PreTraM, Traj-MAE,
and GraphDINO in a semi-supervised setting. Moreover, RedMotion achieves
competitive results compared to HPTR or MTR++ in the Waymo Motion Prediction
Challenge. Our open-source implementation is available at:
https://github.com/kit-mrt/future-motion


## AI 摘要

RedMotion是一种用于自动驾驶车辆运动预测的Transformer模型，通过冗余减少学习环境表征。该模型采用两种冗余减少方法：一是通过内部Transformer解码器将可变大小的局部道路环境标记（包含道路图和智能体数据）压缩为固定大小的全局嵌入；二是通过自监督学习对增强视图生成的环境嵌入应用冗余减少原则。实验表明，在半监督设置下，RedMotion优于PreTraM、Traj-MAE和GraphDINO，并在Waymo运动预测挑战中取得与HPTR和MTR++相当的竞争性结果。项目已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T23:02:27Z
- **目录日期**: 2025-04-02
