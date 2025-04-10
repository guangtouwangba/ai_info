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

该研究表明，在利用专有多模态数据（文本+图像）的多分类任务中，传统嵌入方法（embeddings-based softmax）全面优于当前最先进的LLM提示工程。在家庭服务项目分类任务中，嵌入模型比LLM提示准确率高49.5%，且在概率校准、响应速度（图像快14倍/文本快81倍）和成本（最高节省10倍）方面均具优势。通过A/B测试验证后，该方案已投入实际应用。研究证实，对于能利用专有数据的分类场景，传统方法仍能提供LLM无法比拟的实用价值，建议根据具体需求而非技术热度选择解决方案。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T16:01:59Z
- **目录日期**: 2025-04-10
