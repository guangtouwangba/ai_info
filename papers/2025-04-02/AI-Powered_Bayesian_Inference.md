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

生成式人工智能（GAI）的兴起改变了知识获取方式。尽管GAI的决策不完全可靠，但其输出的可变性可被转化为反映预测确信度的先验分布，与定制数据集结合进行贝叶斯分析。本研究提出非参数贝叶斯框架：以GAI生成模型为基线的狄利克雷过程先验，通过超参数调整评估AI先验信息量，并利用真实数据与AI标注的合成数据生成后验独立同分布样本。该方法通过优化（而非条件采样）实现并行化计算，支持基于AI预测的概率推断与不确定性量化，将GAI的随机性转化为统计优势。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T06:01:12Z
- **目录日期**: 2025-04-02
