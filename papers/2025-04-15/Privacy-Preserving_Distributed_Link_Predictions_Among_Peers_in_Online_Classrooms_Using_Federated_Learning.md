# Privacy-Preserving Distributed Link Predictions Among Peers in Online Classrooms Using Federated Learning

**URL**: http://arxiv.org/abs/2504.10456v1

## 原始摘要

Social interactions among classroom peers, represented as social learning
networks (SLNs), play a crucial role in enhancing learning outcomes. While SLN
analysis has recently garnered attention, most existing approaches rely on
centralized training, where data is aggregated and processed on a local/cloud
server with direct access to raw data. However, in real-world educational
settings, such direct access across multiple classrooms is often restricted due
to privacy concerns. Furthermore, training models on isolated classroom data
prevents the identification of common interaction patterns that exist across
multiple classrooms, thereby limiting model performance. To address these
challenges, we propose one of the first frameworks that integrates Federated
Learning (FL), a distributed and collaborative machine learning (ML) paradigm,
with SLNs derived from students' interactions in multiple classrooms' online
forums to predict future link formations (i.e., interactions) among students.
By leveraging FL, our approach enables collaborative model training across
multiple classrooms while preserving data privacy, as it eliminates the need
for raw data centralization. Recognizing that each classroom may exhibit unique
student interaction dynamics, we further employ model personalization
techniques to adapt the FL model to individual classroom characteristics. Our
results demonstrate the effectiveness of our approach in capturing both shared
and classroom-specific representations of student interactions in SLNs.
Additionally, we utilize explainable AI (XAI) techniques to interpret model
predictions, identifying key factors that influence link formation across
different classrooms. These insights unveil the drivers of social learning
interactions within a privacy-preserving, collaborative, and distributed ML
framework -- an aspect that has not been explored before.


## AI 摘要

该研究提出了一种结合联邦学习(FL)和社会学习网络(SLN)的创新框架，用于预测学生间未来互动。传统集中式训练面临隐私保护和数据孤岛问题，而该方案通过分布式协作学习，在保护各班级原始数据隐私的同时，识别跨班级的共性互动模式。研究采用模型个性化技术适应不同班级的交互特征，并通过可解释AI技术揭示影响链接形成的关键因素。结果表明，该方法能有效捕捉学生互动中的共性和个性特征，为隐私保护下的社交学习分析提供了新思路，首次探索了分布式机器学习框架中的社会学习驱动因素。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T07:01:27Z
- **目录日期**: 2025-04-15
