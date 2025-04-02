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

本文介绍了WeirdFlows，一种用于检测金融欺诈交易和非合规代理的自上而下搜索方法。该方法无需预定义模式或训练数据，能有效识别随时间变化的复杂交易模式，并提供可解释的分析结果。在Intesa Sanpaolo银行8000万笔跨国交易数据上的测试表明，WeirdFlows能高效处理大规模数据，尤其在2022年2月后欧盟经济制裁背景下表现出色。该系统既能发现可疑交易和行为者，又能为反金融犯罪调查提供必要的可解释性，解决了AI模型结果解释性和标注数据缺乏的行业痛点。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T20:01:04Z
- **目录日期**: 2025-04-02
