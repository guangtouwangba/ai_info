# Privacy-preserving neutral atom-based quantum classifier towards real healthcare applications

**URL**: http://arxiv.org/abs/2505.04570v1

## 原始摘要

Technological advances in Artificial Intelligence (AI) and Machine Learning
(ML) for the healthcare domain are rapidly arising, with a growing discussion
regarding the ethical management of their development. In general, ML
healthcare applications crucially require performance, interpretability of
data, and respect for data privacy. The latter is an increasingly debated topic
as commercial cloud computing services become more and more widespread.
Recently, dedicated methods are starting to be developed aiming to protect data
privacy. However, these generally result in a trade-off forcing one to balance
the level of data privacy and the algorithm performance. Here, a Support Vector
Machine (SVM) classifier model is proposed whose training is reformulated into
a Quadratic Unconstrained Binary Optimization (QUBO) problem, and adapted to a
neutral atom-based Quantum Processing Unit (QPU). Our final model does not
require anonymization techniques to protect data privacy since the sensitive
data are not needed to be transferred to the cloud-available QPU. Indeed, the
latter is used only during the training phase, hence allowing a future concrete
application in a real-world scenario. Finally, performance and scaling analyses
on a publicly available breast cancer dataset are discussed, both using ideal
and noisy simulations for the training process, and also successfully tested on
a currently available real neutral-atom QPU.


## AI 摘要

人工智能和机器学习在医疗领域的应用快速发展，但数据隐私保护成为关键挑战。本研究提出一种基于量子处理单元(QPU)的支持向量机(SVM)分类模型，通过将训练过程转化为二次无约束二进制优化(QUBO)问题，在不需传输敏感数据到云端的情况下完成训练，从而避免隐私泄露风险。研究使用乳腺癌数据集测试了理想和噪声模拟下的性能表现，并成功在实际中性原子量子处理器上验证。该方法消除了传统隐私保护技术与算法性能之间的权衡问题，为医疗AI的实际应用提供了可行的隐私保护解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T23:02:03Z
- **目录日期**: 2025-05-08
