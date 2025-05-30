# WATCH: Weighted Adaptive Testing for Changepoint Hypotheses via Weighted-Conformal Martingales

**URL**: http://arxiv.org/abs/2505.04608v1

## 原始摘要

Responsibly deploying artificial intelligence (AI) / machine learning (ML)
systems in high-stakes settings arguably requires not only proof of system
reliability, but moreover continual, post-deployment monitoring to quickly
detect and address any unsafe behavior. Statistical methods for nonparametric
change-point detection -- especially the tools of conformal test martingales
(CTMs) and anytime-valid inference -- offer promising approaches to this
monitoring task. However, existing methods are restricted to monitoring limited
hypothesis classes or ``alarm criteria,'' such as data shifts that violate
certain exchangeability assumptions, or do not allow for online adaptation in
response to shifts. In this paper, we expand the scope of these monitoring
methods by proposing a weighted generalization of conformal test martingales
(WCTMs), which lay a theoretical foundation for online monitoring for any
unexpected changepoints in the data distribution while controlling
false-alarms. For practical applications, we propose specific WCTM algorithms
that accommodate online adaptation to mild covariate shifts (in the marginal
input distribution) while raising alarms in response to more severe shifts,
such as concept shifts (in the conditional label distribution) or extreme
(out-of-support) covariate shifts that cannot be easily adapted to. On
real-world datasets, we demonstrate improved performance relative to
state-of-the-art baselines.


## AI 摘要

本文提出了一种加权广义符合性测试鞅（WCTM）方法，用于人工智能/机器学习系统部署后的持续监控。该方法扩展了现有非参数变点检测技术的应用范围，能够在线检测数据分布中的异常变化（如概念漂移或极端协变量漂移）并控制误报率。与现有方法相比，WCTM支持在线自适应调整以应对温和的协变量漂移，同时对严重漂移保持敏感性。实验表明，该方法在真实数据集上优于当前最优基线。这一技术为高风险场景下的AI系统可靠性监控提供了新的理论框架和实践方案。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T23:01:19Z
- **目录日期**: 2025-05-08
