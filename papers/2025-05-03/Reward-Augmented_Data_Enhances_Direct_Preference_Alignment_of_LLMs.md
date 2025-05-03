# Reward-Augmented Data Enhances Direct Preference Alignment of LLMs

**URL**: http://arxiv.org/abs/2410.08067v4

## 原始摘要

Preference alignment in Large Language Models (LLMs) has significantly
improved their ability to adhere to human instructions and intentions. However,
existing direct alignment algorithms primarily focus on relative preferences
and often overlook the qualitative aspects of responses, despite having access
to preference data that includes reward scores from judge models during AI
feedback. Striving to maximize the implicit reward gap between the chosen and
the slightly inferior rejected responses can cause overfitting and unnecessary
unlearning of the high-quality rejected responses. The unawareness of the
reward scores also drives the LLM to indiscriminately favor the low-quality
chosen responses and fail to generalize to optimal responses that are sparse in
data. To overcome these shortcomings, our study introduces reward-conditioned
LLM policies that discern and learn from the entire spectrum of response
quality within the dataset, helping extrapolate to more optimal regions. We
propose an effective yet simple data relabeling method that conditions the
preference pairs on quality scores to construct a reward-augmented dataset. The
experiments across various benchmarks and diverse models demonstrate that our
approach consistently boosts DPO by a considerable margin. Through
comprehensive ablation studies, we demonstrate that our method not only
maximizes the utility of preference data but also mitigates the issue of
unlearning, demonstrating its broad effectiveness beyond mere data expansion.
Our code is available at
https://github.com/shenao-zhang/reward-augmented-preference.


## AI 摘要

该研究针对大型语言模型（LLM）偏好对齐中忽视响应质量的问题，提出了一种基于奖励条件化的改进方法。现有算法仅关注相对偏好，容易过度拟合或错误遗忘高质量响应。新方法通过数据重标注技术，将偏好对与质量评分结合构建奖励增强数据集，使模型能学习响应质量的整体分布，从而泛化到更优响应。实验表明，该方法显著提升了DPO（Direct Preference Optimization）性能，有效缓解了错误遗忘问题，并最大化偏好数据的效用。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-03T17:01:59Z
- **目录日期**: 2025-05-03
