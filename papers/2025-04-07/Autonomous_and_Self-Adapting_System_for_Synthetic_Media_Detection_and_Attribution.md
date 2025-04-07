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

这篇论文提出了一种自主自适应的合成图像识别系统，能够检测合成图像并识别其来源。与现有静态系统不同，该系统采用开放集识别策略和可演化的嵌入空间，不仅能识别已知生成器，还能通过无监督聚类自动发现新生成器并更新决策边界。实验表明，该方法显著优于现有方案，在生成模型快速发展的背景下，为构建通用、自适应的取证系统迈出了关键一步。该系统可有效应对合成图像带来的虚假信息、欺诈等恶意应用风险。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-07T15:01:17Z
- **目录日期**: 2025-04-07
