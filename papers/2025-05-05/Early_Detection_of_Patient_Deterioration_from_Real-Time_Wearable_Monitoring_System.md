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

TARL是一种创新的早期患者恶化检测方法，通过分析心率时间序列中的关键子序列（shapelets）构建动态知识图谱。该方法利用过渡感知知识嵌入技术强化shapelets间的关系，并量化缺失值的影响，从而创建全面的心率表征。这些表征能捕捉解释性结构并预测未来心率趋势，帮助早期疾病识别。研究团队与医护人员合作，收集ICU患者可穿戴设备的心率数据和疾病严重程度指标进行评估。真实ICU数据实验表明，TARL兼具高可靠性和早期检测能力。案例研究还展示了其可解释的检测过程，有望成为辅助临床决策的AI工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T12:01:54Z
- **目录日期**: 2025-05-05
