# Can Performant LLMs Be Ethical? Quantifying the Impact of Web Crawling Opt-Outs

**URL**: http://arxiv.org/abs/2504.06219v1

## 原始摘要

The increasing adoption of web crawling opt-outs by copyright holders of
online content raises critical questions about the impact of data compliance on
large language model (LLM) performance. However, little is known about how
these restrictions (and the resultant filtering of pretraining datasets) affect
the capabilities of models trained using these corpora. In this work, we
conceptualize this effect as the $\textit{data compliance gap}$ (DCG), which
quantifies the performance difference between models trained on datasets that
comply with web crawling opt-outs, and those that do not. We measure the data
compliance gap in two settings: pretraining models from scratch and continual
pretraining from existing compliant models (simulating a setting where
copyrighted data could be integrated later in pretraining). Our experiments
with 1.5B models show that, as of January 2025, compliance with web data
opt-outs does not degrade general knowledge acquisition (close to 0\% DCG).
However, in specialized domains such as biomedical research, excluding major
publishers leads to performance declines. These findings suggest that while
general-purpose LLMs can be trained to perform equally well using fully open
data, performance in specialized domains may benefit from access to
high-quality copyrighted sources later in training. Our study provides
empirical insights into the long-debated trade-off between data compliance and
downstream model performance, informing future discussions on AI training
practices and policy decisions.


## AI 摘要

这项研究探讨了网络爬虫退出机制对大型语言模型(LLM)性能的影响，提出了"数据合规差距"(DCG)概念来衡量合规与非合规训练数据间的性能差异。实验使用15亿参数模型发现，截至2025年1月，遵守网络数据退出协议不会显著影响通用知识获取(DCG接近0%)，但在生物医学等专业领域，排除主要出版商会导致性能下降。研究表明，通用LLM可以使用完全开放数据训练，而专业领域模型可能受益于后期训练中访问高质量版权内容。这些发现为数据合规与模型性能的权衡提供了实证依据。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T01:28:32Z
- **目录日期**: 2025-04-10
