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

本文介绍了WeirdFlows，一种用于检测金融欺诈交易和违规代理的自动化分析流程。该方法基于交易网络分析，无需预先定义欺诈模式或训练数据，即可识别随时间变化的复杂交易模式，并提供可解释的异常结果，辅助反金融犯罪调查。研究在Intesa Sanpaolo银行8000万笔跨境交易数据上进行验证，结果显示该方法能有效识别可疑交易（特别是2022年2月后欧盟经济制裁相关交易），具备处理大规模数据、检测复杂模式的能力，同时满足调查所需的可解释性要求。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T10:01:10Z
- **目录日期**: 2025-04-02
