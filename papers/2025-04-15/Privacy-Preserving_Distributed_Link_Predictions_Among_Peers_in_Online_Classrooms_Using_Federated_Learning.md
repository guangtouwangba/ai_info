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

该研究提出了一种结合联邦学习(FL)和社会学习网络(SLN)的新框架，用于预测学生间未来互动关系。传统集中式训练面临隐私保护和数据隔离问题，而该分布式方法能在保护各班级原始数据隐私的前提下，通过协作训练识别跨班级的共同互动模式。研究还采用模型个性化技术适应不同班级的独特互动特征，并通过可解释AI技术揭示影响互动形成的关键因素。结果表明，该方法能有效捕捉学生互动的共享特征和班级特异性，为隐私保护下的社交学习研究提供了新思路。这是首个在分布式机器学习框架下探索社会学习驱动因素的研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T13:08:46Z
- **目录日期**: 2025-04-15
