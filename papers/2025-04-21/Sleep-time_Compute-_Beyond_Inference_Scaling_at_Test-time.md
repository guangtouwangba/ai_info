# Sleep-time Compute: Beyond Inference Scaling at Test-time

**URL**: http://arxiv.org/abs/2504.13171v1

## 原始摘要

Scaling test-time compute has emerged as a key ingredient for enabling large
language models (LLMs) to solve difficult problems, but comes with high latency
and inference cost. We introduce sleep-time compute, which allows models to
"think" offline about contexts before queries are presented: by anticipating
what queries users might ask and pre-computing useful quantities, we can
significantly reduce the compute requirements at test-time. To demonstrate the
efficacy of our method, we create modified versions of two reasoning tasks -
Stateful GSM-Symbolic and Stateful AIME. We find that sleep-time compute can
reduce the amount of test-time compute needed to achieve the same accuracy by ~
5x on Stateful GSM-Symbolic and Stateful AIME and that by scaling sleep-time
compute we can further increase accuracy by up to 13% on Stateful GSM-Symbolic
and 18% on Stateful AIME. Furthermore, we introduce Multi-Query GSM-Symbolic,
which extends GSM-Symbolic by including multiple related queries per context.
By amortizing sleep-time compute across related queries about the same context
using Multi-Query GSM-Symbolic, we can decrease the average cost per query by
2.5x. We then conduct additional analysis to understand when sleep-time compute
is most effective, finding the predictability of the user query to be well
correlated with the efficacy of sleep-time compute. Finally, we conduct a
case-study of applying sleep-time compute to a realistic agentic SWE task.


## AI 摘要

本文提出"睡眠时间计算"方法，通过让大语言模型在查询前预先处理上下文并预测可能的问题，显著降低推理时的计算需求。实验表明，该方法在Stateful GSM-Symbolic和Stateful AIME任务上能减少5倍计算量，同时提升13%-18%的准确率。通过Multi-Query GSM-Symbolic任务分摊计算成本，每个查询的平均成本降低了2.5倍。研究发现，用户查询的可预测性与该方法的有效性高度相关。最后，该方法在一个实际软件工程代理任务中进行了验证，展示了其实际应用潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T02:29:51Z
- **目录日期**: 2025-04-21
