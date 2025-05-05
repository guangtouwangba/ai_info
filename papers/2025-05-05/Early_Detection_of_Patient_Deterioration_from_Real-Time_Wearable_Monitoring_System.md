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

TARL是一种创新方法，通过分析心率时间序列中的关键子序列（shapelets）来早期检测患者病情恶化。它构建了一个shapelet转换知识图谱，模拟心率动态变化，揭示疾病进展和未来趋势。该方法采用转换感知知识嵌入，强化shapelet间关联并量化缺失值影响，从而形成全面的心率表征。这些表征能解释心率结构并预测未来趋势，辅助早期病情识别。TARL在真实ICU数据测试中表现出高可靠性和早期检测能力，其可解释性检测流程有望成为AI辅助工具，帮助临床医生识别患者恶化早期迹象。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T17:01:41Z
- **目录日期**: 2025-05-05
