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

研究表明，在专业服务分类任务中，基于嵌入的softmax模型全面优于当前最先进的LLM提示方法。该研究使用Thumbtack客户的家政服务项目描述（文本+图像），比较了两种方法预测专业类别（如水电工、浴室改造）的效果。嵌入模型在准确率（高出49.5%）、校准性、延迟（图像处理快14倍，文本快81倍）和成本（最高便宜10倍）方面均表现更好，且适用于纯文本、纯图像及图文混合场景。通过A/B测试验证后，该方案已部署应用。这说明针对可利用专有数据的多分类问题，传统嵌入方法仍具显著优势，需根据实际需求选择合适模型而非盲目追随LLM热潮。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T15:02:07Z
- **目录日期**: 2025-04-10
