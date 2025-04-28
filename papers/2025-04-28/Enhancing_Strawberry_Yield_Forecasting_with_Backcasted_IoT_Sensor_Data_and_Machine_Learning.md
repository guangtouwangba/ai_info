# Enhancing Strawberry Yield Forecasting with Backcasted IoT Sensor Data and Machine Learning

**URL**: http://arxiv.org/abs/2504.18451v1

## 原始摘要

Due to rapid population growth globally, digitally-enabled agricultural
sectors are crucial for sustainable food production and making informed
decisions about resource management for farmers and various stakeholders. The
deployment of Internet of Things (IoT) technologies that collect real-time
observations of various environmental (e.g., temperature, humidity, etc.) and
operational factors (e.g., irrigation) influencing production is often seen as
a critical step to enable additional novel downstream tasks, such as AI-based
yield forecasting. However, since AI models require large amounts of data, this
creates practical challenges in a real-world dynamic farm setting where IoT
observations would need to be collected over a number of seasons. In this
study, we deployed IoT sensors in strawberry production polytunnels for two
growing seasons to collect environmental data, including water usage, external
and internal temperature, external and internal humidity, soil moisture, soil
temperature, and photosynthetically active radiation. The sensor observations
were combined with manually provided yield records spanning a period of four
seasons. To bridge the gap of missing IoT observations for two additional
seasons, we propose an AI-based backcasting approach to generate synthetic
sensor observations using historical weather data from a nearby weather station
and the existing polytunnel observations. We built an AI-based yield
forecasting model to evaluate our approach using the combination of real and
synthetic observations. Our results demonstrated that incorporating synthetic
data improved yield forecasting accuracy, with models incorporating synthetic
data outperforming those trained only on historical yield, weather records, and
real sensor data.


## AI 摘要

随着全球人口快速增长，数字农业对可持续粮食生产和资源管理决策至关重要。本研究在草莓种植温室部署物联网传感器，收集了两季的环境数据（温湿度、土壤状况等）和四季产量记录。针对早期缺失的传感器数据，研究者提出基于AI的回溯方法，利用历史气象数据生成合成观测值。实验表明，结合合成数据训练的AI产量预测模型，其准确性优于仅使用真实传感器数据、历史产量和气象记录的模型。这为解决农业AI模型训练数据不足问题提供了可行方案，尤其适用于需要长期数据积累的动态农场环境。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T23:01:50Z
- **目录日期**: 2025-04-28
