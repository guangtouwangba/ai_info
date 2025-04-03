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

WeirdFlows是一种自上而下的搜索管道，用于检测潜在欺诈交易和不合规代理。该方法无需预先设定模式或训练数据，通过分析交易网络中的异常模式来识别欺诈行为，并提供可解释的结果辅助反金融犯罪调查。在Intesa Sanpaolo银行8000万笔跨国交易数据集上的测试表明，WeirdFlows能有效处理大规模数据，识别复杂交易模式，尤其在2022年2月后欧盟经济制裁背景下表现突出。该方法兼具可扩展性和可解释性，为反金融犯罪调查提供了实用工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T01:28:17Z
- **目录日期**: 2025-04-03
