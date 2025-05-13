# WATCH: Adaptive Monitoring for AI Deployments via Weighted-Conformal Martingales

**URL**: http://arxiv.org/abs/2505.04608v2

## 原始摘要

Responsibly deploying artificial intelligence (AI) / machine learning (ML)
systems in high-stakes settings arguably requires not only proof of system
reliability, but moreover continual, post-deployment monitoring to quickly
detect and address any unsafe behavior. Statistical methods for nonparametric
change-point detection -- especially the tools of conformal test martingales
(CTMs) and anytime-valid inference -- offer promising approaches to this
monitoring task. However, existing methods are restricted to monitoring limited
hypothesis classes or ``alarm criteria'' (such as data shifts that violate
certain exchangeability assumptions), do not allow for online adaptation in
response to shifts, and/or do not enable root-cause analysis of any
degradation. In this paper, we expand the scope of these monitoring methods by
proposing a weighted generalization of conformal test martingales (WCTMs),
which lay a theoretical foundation for online monitoring for any unexpected
changepoints in the data distribution while controlling false-alarms. For
practical applications, we propose specific WCTM algorithms that adapt online
to mild covariate shifts (in the marginal input distribution) while quickly
detecting and diagnosing more severe shifts, such as concept shifts (in the
conditional label distribution) or extreme (out-of-support) covariate shifts
that cannot be easily adapted to. On real-world datasets, we demonstrate
improved performance relative to state-of-the-art baselines.


## AI 摘要

本文提出了一种加权广义符合测试鞅(WCTM)方法，用于人工智能/机器学习系统在高风险场景下的持续监控。该方法通过非参数变化点检测技术，能够在线监测数据分布中的异常变化并控制误报率。相比现有方法，WCTM突破了监测假设类的限制，允许在线适应轻度协变量偏移，同时快速检测更严重的概念偏移或极端协变量偏移。实验表明，该方法在真实数据集上优于现有最优基线，为系统可靠性监控提供了理论基础和实用工具，支持对性能退化的根源分析。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T10:02:13Z
- **目录日期**: 2025-05-13
