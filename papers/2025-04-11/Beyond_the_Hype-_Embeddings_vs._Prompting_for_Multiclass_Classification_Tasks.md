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

研究表明，在专业服务分类任务（如维修、浴室改造等）中，基于嵌入的Softmax模型在准确率、校准性、延迟和成本上全面优于当前最先进的LLM提示方法。该研究使用Thumbtack客户提供的文本和图像数据，发现嵌入模型比提示方法的准确率高49.5%，且在纯文本、纯图像和图文混合任务中表现一致更优。此外，嵌入模型生成的概率校准良好，适合作为部署中的置信信号，而提示方法的得分信息量不足。嵌入方法处理速度更快（图像快14倍，文本快81倍），成本可降低至1/10。实际A/B测试验证了离线分析结果，证明针对特定分类任务，传统嵌入方法仍具优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-11T01:29:10Z
- **目录日期**: 2025-04-11
