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

随着全球人口快速增长，数字农业对可持续粮食生产和资源管理决策至关重要。物联网(IoT)技术可实时监测温湿度等环境因素和灌溉等生产操作，为AI产量预测提供数据支持。但AI模型需要大量数据，而农场动态环境中需多季节收集数据。本研究在草莓种植园部署IoT传感器收集两季环境数据，结合四季产量记录。针对缺失的两季数据，提出AI回测方法，利用附近气象站历史天气数据和现有观测生成合成传感器数据。实验表明，结合合成数据的产量预测模型优于仅使用真实数据和历史记录的模型。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T18:01:56Z
- **目录日期**: 2025-04-28
