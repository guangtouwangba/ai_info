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

WeirdFlows是一种自上而下的搜索流程，用于检测金融交易网络中的可疑交易和不合规主体。该方法无需预先设定模式或训练数据，能识别随时间变化的复杂欺诈模式，并提供可解释的异常分析结果。在Intesa Sanpaolo银行8000万笔跨国交易数据测试中，该算法有效识别了可疑交易（特别是2022年2月后欧盟经济制裁相关案例），其检测能力和结果可解释性得到了反金融犯罪专家的验证。WeirdFlows解决了AI模型可解释性和标注数据缺乏的挑战，支持大规模数据集分析和正式调查。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T12:01:11Z
- **目录日期**: 2025-04-02
