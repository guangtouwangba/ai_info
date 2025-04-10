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

该研究表明，在家庭服务项目分类任务中，基于嵌入的softmax模型在准确性、校准性、延迟和成本方面全面优于当前最先进的LLM提示方法。嵌入模型的准确率比提示方法高49.5%，且在处理纯文本、纯图像及图文混合输入时表现一致更优。此外，嵌入方法能生成校准良好的概率估计，处理速度比LLM快14-81倍，成本可降低达10倍。通过A/B测试验证后，该方案已投入实际应用。研究证明，对于可利用专有数据集的多分类问题，传统嵌入方法可能比LLM提示框架更具优势，为实际业务场景提供了更优选择。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T19:01:42Z
- **目录日期**: 2025-04-10
