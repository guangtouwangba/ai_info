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

该研究指出，现有大型语言模型（LLM）偏好对齐方法过度关注相对偏好而忽视响应质量，导致对低质量选择的盲目偏好和高质量拒绝响应的遗忘。为此，作者提出一种奖励条件化的LLM策略，通过数据重标方法将偏好对与质量评分结合，构建奖励增强数据集。实验表明，该方法显著提升了DPO（Direct Preference Optimization）性能，有效利用偏好数据并缓解遗忘问题，在多种基准和模型上表现优异。该方法不仅优化了数据效用，还展现出超越单纯数据扩展的广泛有效性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-04T01:29:46Z
- **目录日期**: 2025-05-04
