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

本文提出了一种名为TARL的创新方法，通过分析心率时间序列中的关键子序列（shapelets）来早期检测患者病情恶化。TARL构建了一个shapelet转换知识图谱，建模心率动态变化以预测病情发展，并引入转换感知知识嵌入来强化shapelet间关系及处理缺失值。该方法能生成全面的心率表征，解释病情结构并预测未来趋势。通过与医护人员合作收集ICU患者可穿戴设备的心率数据和诊断指标，实验证明TARL具有高可靠性和早期检测能力。案例研究展示了其可解释的检测过程，凸显其作为AI辅助工具的临床潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T06:01:41Z
- **目录日期**: 2025-05-05
