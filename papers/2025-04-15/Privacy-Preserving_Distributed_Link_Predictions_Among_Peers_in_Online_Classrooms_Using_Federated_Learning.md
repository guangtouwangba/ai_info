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

本研究提出首个结合联邦学习(FL)与社交学习网络(SLN)的框架，用于预测学生在线论坛中的未来互动。传统集中式训练面临隐私限制且无法识别跨教室的共性模式。通过FL实现多教室协同建模而无需共享原始数据，并采用个性化技术适配各教室独特互动特征。结果表明，该方法能有效捕捉SLN中共性和教室特定的交互表征。结合可解释AI技术，揭示了影响不同教室链接形成的关键因素，在保护隐私的分布式机器学习框架下，首次探索了社交学习互动的驱动机制。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T23:02:34Z
- **目录日期**: 2025-04-15
