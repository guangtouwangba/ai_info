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

本文提出了一种基于量子处理单元（QPU）的支持向量机（SVM）分类模型，通过将训练过程重新表述为二次无约束二进制优化（QUBO）问题，解决了医疗AI领域的数据隐私保护难题。该模型无需将敏感数据传输至云端即可完成训练，避免了传统隐私保护技术带来的性能折衷。研究在公开乳腺癌数据集上进行了性能和扩展性分析，包括理想与噪声模拟训练，并成功在实际中性原子QPU上测试。这种方法为医疗AI应用提供了兼顾隐私保护与算法性能的创新解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T21:02:04Z
- **目录日期**: 2025-05-08
