# Towards Reliable and Interpretable Traffic Crash Pattern Prediction and Safety Interventions Using Customized Large Language Models

**URL**: http://arxiv.org/abs/2505.12545v2

## 原始摘要

Predicting crash events is crucial for understanding crash distributions and
their contributing factors, thereby enabling the design of proactive traffic
safety policy interventions. However, existing methods struggle to interpret
the complex interplay among various sources of traffic crash data, including
numeric characteristics, textual reports, crash imagery, environmental
conditions, and driver behavior records. As a result, they often fail to
capture the rich semantic information and intricate interrelationships embedded
in these diverse data sources, limiting their ability to identify critical
crash risk factors. In this research, we propose TrafficSafe, a framework that
adapts LLMs to reframe crash prediction and feature attribution as text-based
reasoning. A multi-modal crash dataset including 58,903 real-world reports
together with belonged infrastructure, environmental, driver, and vehicle
information is collected and textualized into TrafficSafe Event Dataset. By
customizing and fine-tuning LLMs on this dataset, the TrafficSafe LLM achieves
a 42% average improvement in F1-score over baselines. To interpret these
predictions and uncover contributing factors, we introduce TrafficSafe
Attribution, a sentence-level feature attribution framework enabling
conditional risk analysis. Findings show that alcohol-impaired driving is the
leading factor in severe crashes, with aggressive and impairment-related
behaviors having nearly twice the contribution for severe crashes compared to
other driver behaviors. Furthermore, TrafficSafe Attribution highlights pivotal
features during model training, guiding strategic crash data collection for
iterative performance improvements. The proposed TrafficSafe offers a
transformative leap in traffic safety research, providing a blueprint for
translating advanced AI technologies into responsible, actionable, and
life-saving outcomes.


## AI 摘要

本文提出TrafficSafe框架，利用大语言模型（LLMs）将交通事故预测重构为基于文本的推理任务。该框架整合多模态数据（58,903份真实事故报告及环境、驾驶员等信息），通过微调LLMs使预测F1分数提升42%。创新性地开发了TrafficSafe Attribution特征归因方法，揭示酒驾是严重事故首要因素，其风险贡献是其他驾驶行为的两倍。该框架不仅能识别关键风险因素，还能指导数据收集策略优化，为将AI技术转化为可操作的交通安全解决方案提供了新范式，具有显著的社会价值。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T04:03:56Z
- **目录日期**: 2025-05-22
