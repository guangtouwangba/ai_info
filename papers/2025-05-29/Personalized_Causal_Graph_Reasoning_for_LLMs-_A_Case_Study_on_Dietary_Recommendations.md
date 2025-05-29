# Personalized Causal Graph Reasoning for LLMs: A Case Study on Dietary Recommendations

**URL**: http://arxiv.org/abs/2503.00134v2

## 原始摘要

Large Language Models (LLMs) effectively leverage common-sense knowledge for
general reasoning, yet they struggle with personalized reasoning when tasked
with interpreting multifactor personal data. This limitation restricts their
applicability in domains that require context-aware decision-making tailored to
individuals. This paper introduces Personalized Causal Graph Reasoning as an
agentic framework that enhances LLM reasoning by incorporating personal causal
graphs derived from data of individuals. These graphs provide a foundation that
guides the LLM's reasoning process. We evaluate it on a case study on
nutrient-oriented dietary recommendations, which requires personal reasoning
due to the implicit unique dietary effects. We propose a counterfactual
evaluation to estimate the efficiency of LLM-recommended foods for glucose
management. Results demonstrate that the proposed method efficiently provides
personalized dietary recommendations to reduce average glucose iAUC across
three time windows, which outperforms the previous approach. LLM-as-a-judge
evaluation results indicate that our proposed method enhances personalization
in the reasoning process.


## AI 摘要

该研究提出了一种个性化因果图推理框架，通过整合个人数据构建的因果图来增强大语言模型(LLMs)的个性化推理能力。针对传统LLMs在解释多因素个人数据时的局限性，该框架在营养导向的饮食推荐案例中进行了验证。采用反事实评估方法分析LLM推荐食物对血糖管理的效果，结果显示该方法能有效降低三个时间窗口的平均葡萄糖iAUC值，表现优于现有方法。LLM作为评判者的评估也证实，该方法显著提升了推理过程的个性化水平，为需要个性化决策的领域提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T19:02:16Z
- **目录日期**: 2025-05-29
