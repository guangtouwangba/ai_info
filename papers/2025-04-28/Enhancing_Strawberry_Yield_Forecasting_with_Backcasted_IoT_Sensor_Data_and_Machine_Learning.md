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

全球人口快速增长下，数字农业对可持续粮食生产和资源管理决策至关重要。本研究在草莓种植温室部署物联网（IoT）传感器，收集两季环境数据（温湿度、灌溉等），并结合四年产量记录。针对前两季缺失的IoT数据，团队提出基于AI的回溯方法，利用附近气象站数据和现有观测生成合成传感器数据。通过构建产量预测模型验证，发现结合合成数据能显著提升预测准确性，其表现优于仅使用历史产量、气象记录和真实传感器数据的模型。这表明合成数据可有效弥补农业AI模型的数据缺口。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T20:02:07Z
- **目录日期**: 2025-04-28
