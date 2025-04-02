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

WeirdFlows是一种新型自上而下的搜索方法，用于检测金融欺诈交易和违规行为。该方法无需预先设定模式或训练数据，能有效识别随时间变化的复杂交易模式，并提供可解释的分析结果。在Intesa Sanpaolo银行8000万笔跨国交易的真实测试中，该方法成功识别出可疑交易和账户，特别是在2022年2月欧盟经济制裁后的场景下表现突出。WeirdFlows解决了传统反金融犯罪(AFC)调查中AI模型可解释性差和标记数据缺乏的难题，为分析师提供了高效可靠的支持工具，展现了处理大规模数据和复杂模式的强大能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T17:01:02Z
- **目录日期**: 2025-04-02
