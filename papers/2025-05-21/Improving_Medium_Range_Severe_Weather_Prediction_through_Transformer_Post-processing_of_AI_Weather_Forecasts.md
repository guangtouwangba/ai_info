# Improving Medium Range Severe Weather Prediction through Transformer Post-processing of AI Weather Forecasts

**URL**: http://arxiv.org/abs/2505.11750v2

## 原始摘要

Improving the skill of medium-range (1-8 day) severe weather prediction is
crucial for mitigating societal impacts. This study introduces a novel approach
leveraging decoder-only transformer networks to post-process AI-based weather
forecasts, specifically from the Pangu-Weather model, for improved severe
weather guidance. Unlike traditional post-processing methods that use a dense
neural network to predict the probability of severe weather using discrete
forecast samples, our method treats forecast lead times as sequential
``tokens'', enabling the transformer to learn complex temporal relationships
within the evolving atmospheric state. We compare this approach against
post-processing of the Global Forecast System (GFS) using both a traditional
dense neural network and our transformer, as well as configurations that
exclude convective parameters to fairly evaluate the impact of using the
Pangu-Weather AI model. Results demonstrate that the transformer-based
post-processing significantly enhances forecast skill compared to dense neural
networks. Furthermore, AI-driven forecasts, particularly Pangu-Weather
initialized from high resolution analysis, exhibit superior performance to GFS
in the medium-range, even without explicit convective parameters. Our approach
offers improved accuracy, and reliability, which also provides interpretability
through feature attribution analysis, advancing medium-range severe weather
prediction capabilities.


## AI 摘要

本研究提出了一种基于解码器-仅Transformer网络的新方法，用于后处理AI天气预报（如盘古天气模型），以提升1-8天中尺度恶劣天气预测能力。该方法将预报时间序列视为"标记"，使Transformer能学习大气状态演变的复杂时序关系。实验表明，相比传统密集神经网络后处理方法，Transformer显著提高了预测技能。此外，即使不含对流参数，基于高分辨率分析的盘古天气AI预报在中尺度表现也优于全球预报系统（GFS）。该方法通过特征归因分析提供可解释性，兼具准确性、可靠性和预测能力提升。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T04:03:39Z
- **目录日期**: 2025-05-21
