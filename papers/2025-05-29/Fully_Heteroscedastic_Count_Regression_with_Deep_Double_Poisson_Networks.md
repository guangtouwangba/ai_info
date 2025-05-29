# Fully Heteroscedastic Count Regression with Deep Double Poisson Networks

**URL**: http://arxiv.org/abs/2406.09262v4

## 原始摘要

Neural networks capable of accurate, input-conditional uncertainty
representation are essential for real-world AI systems. Deep ensembles of
Gaussian networks have proven highly effective for continuous regression due to
their ability to flexibly represent aleatoric uncertainty via unrestricted
heteroscedastic variance, which in turn enables accurate epistemic uncertainty
estimation. However, no analogous approach exists for count regression, despite
many important applications. To address this gap, we propose the Deep Double
Poisson Network (DDPN), a novel neural discrete count regression model that
outputs the parameters of the Double Poisson distribution, enabling arbitrarily
high or low predictive aleatoric uncertainty for count data and improving
epistemic uncertainty estimation when ensembled. We formalize and prove that
DDPN exhibits robust regression properties similar to heteroscedastic Gaussian
models via learnable loss attenuation, and introduce a simple loss modification
to control this behavior. Experiments on diverse datasets demonstrate that DDPN
outperforms current baselines in accuracy, calibration, and out-of-distribution
detection, establishing a new state-of-the-art in deep count regression.


## AI 摘要

该研究提出了一种新型深度双泊松网络(DDPN)，用于解决离散计数回归中的不确定性估计问题。DDPN通过输出双泊松分布参数，能够灵活表示计数数据的任意程度预测不确定性，并在集成时提升认知不确定性估计。研究证明DDPN具有类似异方差高斯模型的可学习损失衰减特性，并提出了简单的损失修改方法来控制这一行为。实验表明，DDPN在准确性、校准性和分布外检测方面优于现有基线方法，为深度计数回归设立了新的性能标杆。这一成果填补了连续回归与离散计数回归在不确定性表征方法上的重要空白。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T21:02:26Z
- **目录日期**: 2025-05-29
