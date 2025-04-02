# FlowSeries: Anomaly Detection in Financial Transaction Flows

**URL**: http://arxiv.org/abs/2503.15896v2

## 原始摘要

In recent years, the digitization and automation of anti-financial crime
(AFC) investigative processes have faced significant challenges, particularly
the need for interpretability of AI model results and the lack of labeled data
for training. Network analysis has emerged as a valuable approach in this
context.
  In this paper, we present WeirdFlows, a top-down search pipeline for
detecting potentially fraudulent transactions and non-compliant agents. In a
transaction network, fraud attempts are often based on complex transaction
patterns that change over time to avoid detection. The WeirdFlows pipeline
requires neither an a priori set of patterns nor a training set. In addition,
by providing elements to explain the anomalies found, it facilitates and
supports the work of an AFC analyst.
  We evaluate WeirdFlows on a dataset from Intesa Sanpaolo (ISP) bank,
comprising 80 million cross-country transactions over 15 months, benchmarking
our implementation of the algorithm. The results, corroborated by ISP AFC
experts, highlight its effectiveness in identifying suspicious transactions and
actors, particularly in the context of the economic sanctions imposed in the EU
after February 2022. This demonstrates \textit{WeirdFlows}' capability to
handle large datasets, detect complex transaction patterns, and provide the
necessary interpretability for formal AFC investigations.


## AI 摘要

WeirdFlows是一种新型的反金融犯罪(AFC)检测系统，采用自上而下的搜索方法识别可疑交易和违规主体。该系统无需预先设定模式或训练数据，能有效检测随时间变化的复杂交易模式，并提供可解释的分析结果。在Intesa Sanpaolo银行8000万笔跨国交易(15个月)的测试中，该系统成功识别出可疑交易和主体，特别是在2022年2月后欧盟经济制裁背景下表现突出。WeirdFlows能处理海量数据、检测复杂模式，并为正式调查提供可解释性分析，显著提升了AFC调查的效率和准确性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T08:01:07Z
- **目录日期**: 2025-04-02
