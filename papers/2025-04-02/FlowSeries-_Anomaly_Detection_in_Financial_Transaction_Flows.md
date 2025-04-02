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

本文介绍了WeirdFlows，一种用于检测金融犯罪的可解释AI方法。该方法采用自上而下的网络分析策略，无需预先定义欺诈模式或训练数据，即可识别可疑交易和违规主体。在Intesa Sanpaolo银行8000万笔跨国交易的实际测试中，该方法成功检测出复杂交易模式（特别是2022年2月后欧盟经济制裁相关的异常），其分析结果获得了反金融犯罪专家的验证。WeirdFlows的优势在于处理大规模数据、发现动态变化的欺诈模式，并提供可解释的分析结果，有效支持反金融犯罪调查工作。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T07:01:10Z
- **目录日期**: 2025-04-02
