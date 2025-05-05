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

本文提出了一种名为TARL的创新方法，用于通过心率数据早期检测患者病情恶化。TARL通过分析心率时间序列中的关键子序列（shapelets）构建知识图谱，建模其动态变化以预测病情发展。该方法采用转移感知知识嵌入技术强化shapelet间关联，并量化缺失值影响，从而形成全面的心率表征。这些表征能解释病情结构并预测未来趋势。通过与医护人员合作收集ICU患者可穿戴设备数据和诊断指标，实验证明TARL具有高可靠性和早期检测能力。案例研究展示了其可解释的检测过程，凸显其作为AI辅助工具的临床潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T09:01:51Z
- **目录日期**: 2025-05-05
