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

生成式人工智能(GAI)的出现改变了知识获取方式。虽然GAI不能完全信任，但其输出的可变性可被转化为反映预测确信度的先验分布。本研究提出一种非参数贝叶斯框架，将GAI生成模型作为Dirichlet过程先验的基线，通过调整超参数评估AI先验的信息量。后验模拟通过优化方法(而非条件采样)快速生成独立同分布样本，实现了包含观察数据和AI标记假数据的增强数据集分析。该方法以概率方式整合AI预测，支持推理和不确定性量化，为决策提供了一种新范式。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T01:28:27Z
- **目录日期**: 2025-04-03
