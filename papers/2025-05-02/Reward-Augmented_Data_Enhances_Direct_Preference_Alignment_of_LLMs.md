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

现有的大语言模型（LLM）偏好对齐方法主要关注相对偏好，却忽略了响应质量，导致模型可能过度拟合或错误舍弃高质量被拒响应。本研究提出基于奖励条件的LLM策略，通过分析数据集中的完整响应质量谱系，帮助模型泛化至更优区域。方法包括对偏好数据重新标注，构建奖励增强数据集。实验表明，该方法显著提升了DPO性能，不仅最大化偏好数据效用，还缓解了"遗忘"问题，其效果超越单纯的数据扩展。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-02T12:02:14Z
- **目录日期**: 2025-05-02
