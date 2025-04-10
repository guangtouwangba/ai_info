# Beyond the Hype: Embeddings vs. Prompting for Multiclass Classification Tasks

**URL**: http://arxiv.org/abs/2504.04277v2

## 原始摘要

Are traditional classification approaches irrelevant in this era of AI hype?
We show that there are multiclass classification problems where predictive
models holistically outperform LLM prompt-based frameworks. Given text and
images from home-service project descriptions provided by Thumbtack customers,
we build embeddings-based softmax models that predict the professional category
(e.g., handyman, bathroom remodeling) associated with each problem description.
We then compare against prompts that ask state-of-the-art LLM models to solve
the same problem. We find that the embeddings approach outperforms the best LLM
prompts in terms of accuracy, calibration, latency, and financial cost. In
particular, the embeddings approach has 49.5% higher accuracy than the
prompting approach, and its superiority is consistent across text-only,
image-only, and text-image problem descriptions. Furthermore, it yields
well-calibrated probabilities, which we later use as confidence signals to
provide contextualized user experience during deployment. On the contrary,
prompting scores are overly uninformative. Finally, the embeddings approach is
14 and 81 times faster than prompting in processing images and text
respectively, while under realistic deployment assumptions, it can be up to 10
times cheaper. Based on these results, we deployed a variation of the
embeddings approach, and through A/B testing we observed performance consistent
with our offline analysis. Our study shows that for multiclass classification
problems that can leverage proprietary datasets, an embeddings-based approach
may yield unequivocally better results. Hence, scientists, practitioners,
engineers, and business leaders can use our study to go beyond the hype and
consider appropriate predictive models for their classification use cases.


## AI 摘要

该研究表明，在家庭服务项目分类任务中，基于嵌入的softmax模型全面优于当前最先进的LLM提示方法。在Thumbtack客户提供的文本和图像数据上，嵌入方法比最佳LLM提示的准确率高49.5%，且在文本、图像及混合输入中表现一致更优。此外，嵌入方法具有更好的概率校准性、更快的处理速度（图像快14倍，文本快81倍）和更低成本（最多便宜10倍）。通过A/B测试验证后，研究团队已部署该方案。这说明对于可利用专有数据的多分类问题，传统嵌入方法仍明显优于LLM提示框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T13:09:05Z
- **目录日期**: 2025-04-10
