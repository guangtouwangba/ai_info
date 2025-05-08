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

本文提出了一种加权广义符合性测试赌博机(WCTM)方法，用于高风险AI/ML系统的持续监控。该方法扩展了现有非参数变化点检测技术的应用范围，能够在线监测数据分布中的任何异常变化点，同时控制误报率。针对实际应用，作者设计了特定WCTM算法，可适应轻度协变量偏移，并对严重偏移(如概念漂移或极端协变量偏移)发出警报。相比现有方法，WCTM突破了有限假设类别的限制，支持在线自适应调整。在真实数据集上的实验表明，该方法优于当前最先进的基线方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T21:01:15Z
- **目录日期**: 2025-05-08
