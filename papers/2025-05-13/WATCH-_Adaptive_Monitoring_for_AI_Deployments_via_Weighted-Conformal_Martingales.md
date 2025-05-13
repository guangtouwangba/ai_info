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

本文提出了一种加权泛化的共形测试鞅(WCTM)方法，用于高风险AI/ML系统的持续部署后监控。该方法扩展了非参数变化点检测的统计工具，能在线监控数据分布中的任何意外变化点并控制误报。具体算法可在线适应轻度协变量偏移，同时快速检测和诊断更严重的变化(如概念漂移或极端协变量偏移)。相比现有方法，WCTM突破了有限假设类别的限制，支持在线自适应调整，并能进行性能退化的根因分析。在真实数据集上的实验表明，其性能优于现有最优基线方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T14:02:32Z
- **目录日期**: 2025-05-13
