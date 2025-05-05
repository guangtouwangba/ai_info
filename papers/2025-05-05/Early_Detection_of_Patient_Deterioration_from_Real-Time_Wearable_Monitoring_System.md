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

该研究提出了一种名为TARL的创新方法，用于通过可穿戴设备监测心率数据，实现患者病情恶化的早期检测。TARL通过建模心率时间序列中关键子序列（shapelets）的结构关系，构建知识图谱来追踪病情发展。该方法采用转移感知知识嵌入技术，强化子序列间关联并量化缺失值影响，从而生成全面的心率表征，既能解释当前状态又能预测未来趋势。通过与医护人员合作收集ICU患者数据验证，TARL展现出高可靠性和早期预警能力。案例研究证实其可解释性检测过程，有望成为辅助临床识别病情恶化的AI工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T23:01:43Z
- **目录日期**: 2025-05-05
