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

该研究提出了一种基于个性化因果图推理的智能框架，以增强大语言模型（LLMs）在个性化推理上的能力。针对LLMs在处理多因素个人数据时的局限性，该框架通过构建个人因果图来指导推理过程。在营养导向的饮食推荐案例中，采用反事实评估方法验证了该框架能有效降低血糖波动（iAUC），效果优于传统方法。LLM-as-a-judge评估进一步证实该方法显著提升了推理过程的个性化水平。这一成果为需要个性化决策的领域提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T11:02:22Z
- **目录日期**: 2025-05-29
