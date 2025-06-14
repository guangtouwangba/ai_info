# JuStRank: Benchmarking LLM Judges for System Ranking

**URL**: http://arxiv.org/abs/2412.09569v2

## 原始摘要

Given the rapid progress of generative AI, there is a pressing need to
systematically compare and choose between the numerous models and
configurations available. The scale and versatility of such evaluations make
the use of LLM-based judges a compelling solution for this challenge.
Crucially, this approach requires first to validate the quality of the LLM
judge itself. Previous work has focused on instance-based assessment of LLM
judges, where a judge is evaluated over a set of responses, or response pairs,
while being agnostic to their source systems. We argue that this setting
overlooks critical factors affecting system-level ranking, such as a judge's
positive or negative bias towards certain systems. To address this gap, we
conduct the first large-scale study of LLM judges as system rankers. System
scores are generated by aggregating judgment scores over multiple system
outputs, and the judge's quality is assessed by comparing the resulting system
ranking to a human-based ranking. Beyond overall judge assessment, our analysis
provides a fine-grained characterization of judge behavior, including their
decisiveness and bias.


## AI 摘要

随着生成式AI的快速发展，如何系统评估和选择不同模型及配置成为关键问题。基于大语言模型(LLM)的评估方法因其规模优势备受关注，但需先验证LLM评估者本身的质量。现有研究多关注实例级评估，而忽视了系统级排序中的关键因素（如评估者对特定系统的偏好偏差）。为此，本研究首次对LLM评估者作为系统排序工具进行了大规模分析，通过聚合多个系统输出的评分生成系统得分，并与人工排序对比来评估LLM评估质量。研究不仅考察整体表现，还深入分析了评估者的决策倾向性和偏差特征。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T20:02:25Z
- **目录日期**: 2025-06-11
