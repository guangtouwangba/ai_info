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

本文提出WeirdFlows反金融犯罪检测系统，通过自上而下的网络分析方法识别可疑交易和违规主体。该系统无需预定义模式或训练数据，能自动检测随时间演变的复杂交易模式，并提供可解释的分析结果。在Intesa Sanpaolo银行8000万笔跨国交易数据集上的测试表明，该系统能有效识别可疑交易（特别是2022年2月后欧盟经济制裁相关案例），其分析结果获得银行反金融犯罪专家的认可。WeirdFlows展现了处理大规模数据、检测复杂模式的能力，同时满足反洗钱调查所需的可解释性要求。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T06:00:58Z
- **目录日期**: 2025-04-02
