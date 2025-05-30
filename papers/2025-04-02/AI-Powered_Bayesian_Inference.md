# AI-Powered Bayesian Inference

**URL**: http://arxiv.org/abs/2502.19231v2

## 原始摘要

The advent of Generative Artificial Intelligence (GAI) has heralded an
inflection point that changed how society thinks about knowledge acquisition.
While GAI cannot be fully trusted for decision-making, it may still provide
valuable information that can be integrated into a decision pipeline. Rather
than seeing the lack of certitude and inherent randomness of GAI as a problem,
we view it as an opportunity. Indeed, variable answers to given prompts can be
leveraged to construct a prior distribution which reflects assuredness of AI
predictions. This prior distribution may be combined with tailored datasets for
a fully Bayesian analysis with an AI-driven prior. In this paper, we explore
such a possibility within a non-parametric Bayesian framework. The basic idea
consists of assigning a Dirichlet process prior distribution on the
data-generating distribution with AI generative model as its baseline.
Hyper-parameters of the prior can be tuned out-of-sample to assess the
informativeness of the AI prior. Posterior simulation is achieved by computing
a suitably randomized functional on an augmented data that consists of observed
(labeled) data as well as fake data whose labels have been imputed using AI.
This strategy can be parallelized and rapidly produces iid samples from the
posterior by optimization as opposed to sampling from conditionals. Our method
enables (predictive) inference and uncertainty quantification leveraging AI
predictions in a coherent probabilistic manner.


## AI 摘要

生成式人工智能（GAI）的兴起改变了知识获取方式。尽管GAI的决策不完全可信，但其生成的信息仍可整合到决策流程中。研究者将GAI输出的不确定性视为优势，利用其随机性构建反映预测可信度的先验分布，并与定制数据集结合进行贝叶斯分析。本文提出非参数贝叶斯框架，以GAI生成模型为基线的狄利克雷过程作为先验分布，通过超参数调优评估先验信息量。后验模拟通过优化标记数据和AI生成的模拟数据实现，支持并行计算并快速生成独立同分布样本，从而以概率方式实现预测推断和不确定性量化。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T23:01:11Z
- **目录日期**: 2025-04-02
