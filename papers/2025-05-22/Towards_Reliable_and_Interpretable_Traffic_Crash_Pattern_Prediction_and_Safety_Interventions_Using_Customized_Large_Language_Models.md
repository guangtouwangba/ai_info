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

本文提出TrafficSafe框架，利用大语言模型(LLM)处理多模态交通事故数据（包括数值特征、文本报告、图像、环境条件和驾驶行为记录），将事故预测重构为基于文本的推理任务。该框架收集了58,903份真实事故报告及相关数据，构建TrafficSafe事件数据集。经微调的TrafficSafe LLM在F1分数上比基线方法平均提升42%。通过提出的句子级特征归因框架TrafficSafe Attribution，研究发现酒驾是严重事故的主因，其风险贡献是其他驾驶行为的两倍。该框架还能指导数据收集策略，为交通安全研究提供了AI驱动的创新方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T08:02:30Z
- **目录日期**: 2025-05-22
