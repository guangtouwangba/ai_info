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

本研究提出了一种基于解码器-Transformer网络的新方法，用于后处理AI天气预报（如盘古天气模型），以提升1-8天中短期强天气预测能力。与传统密集神经网络不同，该方法将预报时间序列视为"标记"，利用Transformer捕捉大气状态的时序关系。实验表明，Transformer后处理显著优于传统方法，且基于高分辨率分析的盘古天气模型即使不含对流参数，其表现也优于全球预报系统（GFS）。该方法提高了预测准确性、可靠性，并通过特征归因分析增强可解释性，推动了中短期强天气预测技术的发展。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T21:02:25Z
- **目录日期**: 2025-05-21
