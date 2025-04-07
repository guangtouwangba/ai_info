# Autonomous and Self-Adapting System for Synthetic Media Detection and Attribution

**URL**: http://arxiv.org/abs/2504.03615v1

## 原始摘要

Rapid advances in generative AI have enabled the creation of highly realistic
synthetic images, which, while beneficial in many domains, also pose serious
risks in terms of disinformation, fraud, and other malicious applications.
Current synthetic image identification systems are typically static, relying on
feature representations learned from known generators; as new generative models
emerge, these systems suffer from severe performance degradation. In this
paper, we introduce the concept of an autonomous self-adaptive synthetic media
identification system -- one that not only detects synthetic images and
attributes them to known sources but also autonomously identifies and
incorporates novel generators without human intervention. Our approach
leverages an open-set identification strategy with an evolvable embedding space
that distinguishes between known and unknown sources. By employing an
unsupervised clustering method to aggregate unknown samples into
high-confidence clusters and continuously refining its decision boundaries, our
system maintains robust detection and attribution performance even as the
generative landscape evolves. Extensive experiments demonstrate that our method
significantly outperforms existing approaches, marking a crucial step toward
universal, adaptable forensic systems in the era of rapidly advancing
generative models.


## AI 摘要

这篇论文提出了一种自主自适应的合成图像识别系统，能够检测合成图像并追踪其来源。与现有静态系统不同，该方法采用开放集识别策略，通过可演化的嵌入空间区分已知和未知生成器。系统利用无监督聚类将未知样本聚集成高置信度簇，并持续优化决策边界，从而在新型生成模型出现时仍保持强大检测能力。实验表明，该方法显著优于现有技术，为应对快速发展的生成模型提供了更通用的取证解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-07T06:01:16Z
- **目录日期**: 2025-04-07
