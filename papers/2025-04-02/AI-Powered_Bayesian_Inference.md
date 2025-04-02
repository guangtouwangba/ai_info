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

生成式人工智能（GAI）的出现改变了知识获取方式，其不确定性可被转化为构建先验分布的机会。本研究提出一种非参数贝叶斯方法，将GAI生成结果作为狄利克雷过程先验的基线，通过调整超参数评估AI先验的信息量。后验模拟结合观测数据和AI生成的模拟标记数据，采用优化而非条件采样快速获得独立同分布样本。该方法实现了基于AI预测的（预测性）推断和不确定性量化，以概率化方式将GAI信息整合到决策流程中，为AI辅助决策提供新范式。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T13:08:10Z
- **目录日期**: 2025-04-02
