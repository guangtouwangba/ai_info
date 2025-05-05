# Early Detection of Patient Deterioration from Real-Time Wearable Monitoring System

**URL**: http://arxiv.org/abs/2505.01305v1

## 原始摘要

Early detection of patient deterioration is crucial for reducing mortality
rates. Heart rate data has shown promise in assessing patient health, and
wearable devices offer a cost-effective solution for real-time monitoring.
However, extracting meaningful insights from diverse heart rate data and
handling missing values in wearable device data remain key challenges. To
address these challenges, we propose TARL, an innovative approach that models
the structural relationships of representative subsequences, known as
shapelets, in heart rate time series. TARL creates a shapelet-transition
knowledge graph to model shapelet dynamics in heart rate time series,
indicating illness progression and potential future changes. We further
introduce a transition-aware knowledge embedding to reinforce relationships
among shapelets and quantify the impact of missing values, enabling the
formulation of comprehensive heart rate representations. These representations
capture explanatory structures and predict future heart rate trends, aiding
early illness detection. We collaborate with physicians and nurses to gather
ICU patient heart rate data from wearables and diagnostic metrics assessing
illness severity for evaluating deterioration. Experiments on real-world ICU
data demonstrate that TARL achieves both high reliability and early detection.
A case study further showcases TARL's explainable detection process,
highlighting its potential as an AI-driven tool to assist clinicians in
recognizing early signs of patient deterioration.


## AI 摘要

TARL是一种创新方法，通过建模心率时间序列中的关键子序列（shapelets）及其动态变化，构建知识图谱来预测患者病情发展。该方法利用过渡感知知识嵌入强化shapelet间关系，量化缺失值影响，形成全面心率表征，从而捕捉病情解释性结构并预测未来趋势。结合可穿戴设备采集的ICU患者心率和诊断数据，TARL实现了高可靠性和早期病情恶化检测。案例研究验证了其可解释性，展示了作为AI辅助工具帮助临床识别早期恶化征兆的潜力。该方法解决了可穿戴数据多样性和缺失值处理的关键挑战。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T04:02:48Z
- **目录日期**: 2025-05-05
